package model

import (
	"github.com/astaxie/beego/orm"
)

type StockDaily struct {
	Id                 int64
	Name               string
	Code               string
	CurPrice           int64  //当前价格
	HighPrice          int64  //最高价格
	LowPrice           int64  //最低价格
	IncreaseAmount     int64  //涨跌额
	IncreaseRate       int64  //涨跌幅
	Turnover           int64  //成交额
	Volumn             int64  //成交量
	MainBuy            int64  //主力买入
	MainBuyRate        int64  //主力买入占比
	MainSell           int64  //主力卖出
	MainSellRate       int64  //主力卖出占比
	IndividualBuy      int64  //散户买入
	IndividualBuyRate  int64  //散户买入占比
	IndividualSell     int64  //散户卖出
	IndividualSellRate int64  //散户卖出占比
	MarketDay          string //日期
}

func (m *Model) InsertOrUpdateStockDaily(sd *StockDaily) (int64, error) {
	var t StockDaily
	err := m.o.QueryTable("stock_daily").Filter("code", sd.Code).Filter("market_day", sd.MarketDay).One(&t)
	if err == nil {
		sd.Id = t.Id
		return m.o.Update(sd)
	}
	if err == orm.ErrNoRows {
		return m.o.Insert(sd)
	}
	return 0, nil
}
