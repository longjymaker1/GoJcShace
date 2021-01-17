package main

import (
	_ "GoJcShare/models"
	_ "GoJcShare/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	maxIdle := 500
	maxConn := 500
	orm.RegisterDataBase("default", "mysql",
		"root:123qwe@tcp(127.0.0.1:3306)/go_jc_share?charset=utf8&loc=Asia%2FShanghai",
		maxIdle, maxConn)
}

func main() {
	orm.Debug = true
	orm.RunSyncdb("default", false, true)

	beego.Run()
}
