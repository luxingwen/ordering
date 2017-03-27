package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/config"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

var (
	AppConfig config.Configer
	Token     string
)

func init() {
	AppConfig = beego.AppConfig
	Token = AppConfig.DefaultString("token::token", "")
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", AppConfig.DefaultString("mysql::server", ""))
}
