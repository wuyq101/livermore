package model

import (
	"github.com/astaxie/beego/orm"
)

type MoneyFlow struct {
	Id           int64
	Name         string
	Code         string
	ClosingPrice int64 //收盘价
	IncreaseRate int64 //涨跌幅
	TurnoverRate int64 //换手率
	NetAmount    int64 //净流入
	RatioAmount  int64 //净流入 比率 1587 --> 15.87%
	R0           int64 // 超大单成交额
	R0Net        int64 // 超大单净流入
	R1           int64 //大单成交额
	R1Net        int64 //大单净流入
	R2           int64 //小单成交额
	R2Net        int64 //小单净流入
	R3           int64 //散单成交额
	R3Net        int64 //散单净流入
	MarketDay    string
}

func (m *Model) GetLastMoneyFlow(code string) (*MoneyFlow, error) {
	var t MoneyFlow
	err := m.o.QueryTable("money_flow").Filter("code", code).OrderBy("-market_day").Limit(1).One(&t)
	return &t, err
}

func (m *Model) InsertOrUpdateMoneyFlow(mf *MoneyFlow) (int64, error) {
	var t MoneyFlow
	err := m.o.QueryTable("money_flow").Filter("code", mf.Code).Filter("market_day", mf.MarketDay).One(&t)
	if err == nil {
		mf.Id = t.Id
		return m.o.Update(mf)
	}
	if err == orm.ErrNoRows {
		return m.o.Insert(mf)
	}
	return 0, nil
}
