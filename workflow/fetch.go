package workflow

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/guotie/gogb2312"
	"github.com/wuyq101/livermore/logger"
	"github.com/wuyq101/livermore/model"
	"github.com/wuyq101/livermore/util"
	"regexp"
	"strconv"
	"strings"
	"time"
	"unicode/utf8"
)

func (w *WorkFlow) FetchStockInfo(codes []string) error {
	for _, code := range codes {
		w.FetchStockInfoByCode(code)
		w.FetchStockDailyByCode(code)
		w.FetchMoneyFlow(code)
	}
	return nil
}

func (w *WorkFlow) FetchMoneyFlow(code string) error {
	mf, err := w.m.GetLastMoneyFlow(code)
	cnt := int64(10000) //抓取条数
	if err == nil && mf != nil {
		t, _ := time.ParseInLocation("20060102", mf.MarketDay, time.Local)
		cnt = int64(time.Now().Sub(t).Hours()/24.0) + 2
	}
	logger.Info("cnt %d, code %s", cnt, code)
	stock, err := w.m.GetStockByCode(code)
	if err != nil {
		logger.Error("Failed to get stock by code %s, err %+v", code, err)
		return err
	}
	url := fmt.Sprintf("http://vip.stock.finance.sina.com.cn/quotes_service/api/json_v2.php/MoneyFlow.ssl_qsfx_lscjfb?page=1&num=%d&sort=opendate&asc=0&daima=%s", cnt, code)
	content, err := util.HttpGet(url)
	html := string(content)
	html = strings.Replace(html, "opendate", `"opendate"`, -1)
	html = strings.Replace(html, "trade", `"trade"`, -1)
	html = strings.Replace(html, "turnover", `"turnover"`, -1)
	html = strings.Replace(html, "netamount", `"netamount"`, -1)
	html = strings.Replace(html, "ratioamount", `"ratioamount"`, -1)
	html = strings.Replace(html, "changeratio", `"changeratio"`, -1)
	html = strings.Replace(html, "r0", `"r0"`, -1)
	html = strings.Replace(html, `"r0"_net`, `"r0_net"`, -1)
	html = strings.Replace(html, "r1", `"r1"`, -1)
	html = strings.Replace(html, `"r1"_net`, `"r1_net"`, -1)
	html = strings.Replace(html, "r2", `"r2"`, -1)
	html = strings.Replace(html, `"r2"_net`, `"r2_net"`, -1)
	html = strings.Replace(html, "r3", `"r3"`, -1)
	html = strings.Replace(html, `"r3"_net`, `"r3_net"`, -1)
	reg := regexp.MustCompile("{(.*?)}")
	strs := reg.FindAllString(html, -1)
	out := make(map[string]interface{})
	for i := 0; i < len(strs); i++ {
		err = json.Unmarshal([]byte(strs[i]), &out)
		if err == nil {
			mf, err := w.parseMoneyFlow(out, code, stock.Name)
			if err == nil {
				w.m.InsertOrUpdateMoneyFlow(mf)
			} else {
				logger.Error("err: %+v", err)
			}
		} else {
			logger.Error("err %+v", err)
		}
	}
	return nil
}

func (w *WorkFlow) parseMoneyFlow(mp map[string]interface{}, code string, name string) (*model.MoneyFlow, error) {
	mf := &model.MoneyFlow{
		Name: name,
		Code: code,
	}
	str := mp["opendate"].(string)
	str = strings.Replace(str, "-", "", -1)
	mf.MarketDay = str
	//收盘价
	mf.ClosingPrice = w.getFloat2Int64(mp, "trade", 100)
	//涨跌幅
	mf.IncreaseRate = w.getFloat2Int64(mp, "changeratio", 10000)
	//换手率
	mf.TurnoverRate = w.getFloat2Int64(mp, "turnover", 100)
	//净流入
	mf.NetAmount = w.getFloat2Int64(mp, "netamount", 10000)
	//净流入比率
	mf.RatioAmount = w.getFloat2Int64(mp, "ratioamount", 10000)
	mf.R0 = w.getFloat2Int64(mp, "r0", 10000)
	mf.R0Net = w.getFloat2Int64(mp, "r0_net", 10000)
	mf.R1 = w.getFloat2Int64(mp, "r1", 10000)
	mf.R1Net = w.getFloat2Int64(mp, "r1_net", 10000)
	mf.R2 = w.getFloat2Int64(mp, "r2", 10000)
	mf.R2Net = w.getFloat2Int64(mp, "r2_net", 10000)
	mf.R3 = w.getFloat2Int64(mp, "r3", 10000)
	mf.R3Net = w.getFloat2Int64(mp, "r3_net", 10000)
	return mf, nil
}

func (w *WorkFlow) FetchAllStockInfo() error {
	stocks, err := w.m.GetAllStocks()
	if err != nil {
		logger.Error("Failed to get all stocks, err %+v", err)
		return err
	}
	codes := make([]string, 0)
	for _, s := range stocks {
		codes = append(codes, s.Code)
	}
	return w.FetchStockInfo(codes)
}

func (w *WorkFlow) FetchStockDailyByCode(code string) error {
	fmt.Println("Start to fetch stock daily info for ", code)
	content, err := util.HttpGet("http://stocks.sina.cn/aj/sh/info?code=" + code)
	if err != nil {
		logger.Error("Fethc stock daily err %+v", err)
		return err
	}
	html := w.convertUnicodeString(string(content))
	out := make(map[string]interface{})
	idx := strings.Index(html, `"data":`)
	html = html[idx+7 : len(html)-1]
	err = json.Unmarshal([]byte(html), &out)
	if err != nil {
		logger.Error("Failed to parse json for stock daily info, err %+v", err)
		return err
	}
	stock, err := w.m.GetStockByCode(code)
	if err != nil {
		logger.Error("Failed to get stock by code %s, err %+v", code, err)
		return err
	}
	sd := &model.StockDaily{
		Code: stock.Code,
		Name: stock.Name,
	}
	hq := out["hq"].(map[string]interface{})
	//涨跌额 单位分
	str := hq["zhangdiee"].(string)
	v, _ := strconv.ParseFloat(str, 64)
	sd.IncreaseAmount = int64(v*100 + 0.5)
	//涨跌幅   百分比
	str = hq["zhangdiefu"].(string)
	v, _ = strconv.ParseFloat(str[0:len(str)-1], 64)
	sd.IncreaseRate = int64(v*100 + 0.5)
	//当前价格  单位分
	str = hq["zuixin"].(string)
	v, _ = strconv.ParseFloat(str, 64)
	sd.CurPrice = int64(v*100 + 0.5)
	//最高价格 单位分
	str = hq["zuigao"].(string)
	v, _ = strconv.ParseFloat(str, 64)
	sd.HighPrice = int64(v*100 + 0.5)
	//最低价格 单位分
	str = hq["zuidi"].(string)
	v, _ = strconv.ParseFloat(str, 64)
	sd.LowPrice = int64(v*100 + 0.5)
	//成交量  手
	str = hq["chengjiaoliang"].(string)
	r, _ := utf8.DecodeLastRune([]byte(str))
	idx = strings.IndexRune(str, r)
	str = str[0:idx]
	str = strings.Replace(str, ",", "", -1)
	intv, _ := strconv.ParseInt(str, 10, 64)
	sd.Volumn = intv
	//成交额  分
	str = hq["chengjiaoe"].(string)
	r, _ = utf8.DecodeLastRune([]byte(str))
	idx = strings.IndexRune(str, r)
	v, _ = strconv.ParseFloat(str[0:idx], 64)
	sd.Turnover = int64(v*10000000000 + 0.5)

	ms := out["market_status"].(map[string]interface{})
	//交易日期
	t := ms["time"].(string)
	tm, _ := time.ParseInLocation("2006-01-02 15:04:05", t, time.Local)
	sd.MarketDay = tm.Format("20060102")

	ff := out["fund_flow"].(map[string]interface{})
	//主力买入   分
	sd.MainBuy = w.getFloat2Int64(ff, "zhulimairu", 1000000)
	sd.MainBuyRate = w.getFloat2Int64(ff, "zhulimairubi", 100)
	//主力卖出
	sd.MainSell = w.getFloat2Int64(ff, "zhulimaichu", 1000000)
	sd.MainSellRate = w.getFloat2Int64(ff, "zhulimaichubi", 100)
	//散户买入
	sd.IndividualBuy = w.getFloat2Int64(ff, "sanhumairu", 1000000)
	sd.IndividualBuyRate = w.getFloat2Int64(ff, "sanhumairubi", 100)
	//散户卖出
	sd.IndividualSell = w.getFloat2Int64(ff, "sanhumaichu", 1000000)
	sd.IndividualSellRate = w.getFloat2Int64(ff, "sanhumaichubi", 100)
	w.m.InsertOrUpdateStockDaily(sd)
	return nil
}

func (w *WorkFlow) getFloat2Int64(mp map[string]interface{}, key string, rate int64) int64 {
	str := mp[key].(string)
	v, _ := strconv.ParseFloat(str, 64)
	return int64(v*float64(rate) + 0.5)
}

func (w *WorkFlow) convertUnicodeString(str string) string {
	buf := bytes.NewBufferString("")
	length := len(str)
	for i := 0; i < length; i++ {
		if i+1 < length && str[i] == '\\' && str[i+1] == 'u' {
			//read next four byte
			tmp := str[i+2 : i+6]
			v, _ := strconv.ParseInt(tmp, 16, 64)
			buf.WriteRune(rune(v))
			i = i + 5
			continue
		}
		buf.WriteByte(str[i])
	}
	return buf.String()
}

func (w *WorkFlow) FetchStockInfoByCode(code string) error {
	fmt.Println("Start to fetch data for ", code)
	content, err := util.HttpGet("http://finance.sina.com.cn/realstock/company/" + code + "/nc.shtml")
	if err != nil {
		logger.Error("Fetch stock data err %+v", err)
		return err
	}
	content, _, _, _ = gogb2312.ConvertGB2312(content)
	//	fmt.Println(string(content))
	//start to parse content, extract stock info
	html := string(content)
	//总股本 万股
	reg := regexp.MustCompile(`var totalcapital = (.*)+; //`)
	totalCapitalStr := reg.FindStringSubmatch(html)[1]
	v, _ := strconv.ParseFloat(totalCapitalStr, 64)
	totalCapital := int64(v * 10000)
	//流通股本 万股
	reg = regexp.MustCompile("var currcapital = (.*)+; //")
	currCapitalStr := reg.FindStringSubmatch(html)[1]
	v, _ = strconv.ParseFloat(currCapitalStr, 64)
	currCapital := int64(v * 10000)
	//股票名称
	reg = regexp.MustCompile(`var stockname = '(.*)+'; //`)
	name := reg.FindStringSubmatch(html)[1]
	fmt.Printf("total capital %d, curr capital %d name %s \n", totalCapital, currCapital, name)
	stock := &model.Stock{
		Name:         name,
		Code:         code,
		TotalCapital: totalCapital, //
		CurrCapital:  currCapital,
	}
	w.m.InsertOrUpdateStock(stock)
	return nil
}
