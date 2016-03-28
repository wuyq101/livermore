package model

import (
	"github.com/astaxie/beego/orm"
)

type StockDaily struct {
	Id                 int64
	Name               string
	Code               string
	CurPrice           int64  //当前价格  单位分
	HighPrice          int64  //最高价格  单位分
	LowPrice           int64  //最低价格  单位分
	IncreaseAmount     int64  //涨跌额    单位分
	IncreaseRate       int64  //涨跌幅  百分比
	Turnover           int64  //成交额   单位分
	Volumn             int64  //成交量  手
	MainBuy            int64  //主力买入   单位分
	MainBuyRate        int64  //主力买入占比  单位万分之一
	MainSell           int64  //主力卖出 单位分
	MainSellRate       int64  //主力卖出占比 单位万分之一
	IndividualBuy      int64  //散户买入 单位分
	IndividualBuyRate  int64  //散户买入占比 单位万分之一
	IndividualSell     int64  //散户卖出 单位分
	IndividualSellRate int64  //散户卖出占比 单位万分之一
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

func (m *Model) GetLastStockDaily(code string) (*StockDaily, error) {
	var sd StockDaily
	err := m.o.QueryTable("stock_daily").Filter("code", code).OrderBy("-market_day").Limit(1).One(&sd)
	if err != nil {
		return nil, err
	}
	return &sd, nil
}
