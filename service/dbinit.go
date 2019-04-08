package service

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"log"
	"github.com/astaxie/beego"
)

var engine *xorm.Engine

func init()  {
	mysqlurl := beego.AppConfig.String("mysql.url")
	var err error
	engine, err = xorm.NewEngine("mysql", mysqlurl)
	if err != nil {
		log.Fatal("初始化数据库失败")
	}
	err = engine.Ping()
	if err != nil {
		log.Fatal("初始化数据库失败")
	}
}