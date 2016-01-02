package model

import (
	"github.com/astaxie/beego/orm"
)

type Stock struct {
	Id           int64
	Name         string //股票名称
	Code         string //股票代码
	TotalCapital int64  //总市值  单位股
	CurrCapital  int64  //流通市值 单位股
}

func (m *Model) InsertOrUpdateStock(stock *Stock) (int64, error) {
	var t Stock
	err := m.o.QueryTable("stock").Filter("code", stock.Code).One(&t)
	if err == nil {
		stock.Id = t.Id
		return m.o.Update(stock)
	}
	if err == orm.ErrNoRows {
		return m.o.Insert(stock)
	}
	return 0, nil
}

func (m *Model) GetStockByCode(code string) (*Stock, error) {
	var s Stock
	err := m.o.QueryTable("stock").Filter("code", code).One(&s)
	return &s, err
}

func (m *Model) GetAllStocks() ([]*Stock, error) {
	var stocks []*Stock
	_, err := m.o.QueryTable("stock").All(&stocks)
	return stocks, err
}
