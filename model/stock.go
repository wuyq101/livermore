package model

type Stock struct {
	Id             int64
	Name           string //股票名称
	Code           string //股票代码
	Type           int    //类型 深圳 or 上海
	TotalValue     int64  //总市值  单位分
	CirculateValue int64  //流通市值 单位分
}

func (m *Model) InsertStock(stock *Stock) (int64, error) {
	return m.o.Insert(stock)
}
