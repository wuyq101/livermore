package model

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"github.com/wuyq101/livermore/config"
)

type Model struct {
	o orm.Ormer
}

func NewModel() *Model {
	ret := &Model{
		o: orm.NewOrm(),
	}
	ret.o.Using("default")
	return ret
}

func init() {
	conf := config.Instance()
	orm.Debug = true
	orm.RegisterDataBase("default", "mysql", conf.DbConnStr)
	orm.SetMaxIdleConns("default", 30)
	orm.SetMaxOpenConns("default", 30)
	orm.RegisterModel(new(Stock))
}
