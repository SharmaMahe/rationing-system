package main

import (
	_ "app/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
    _ "github.com/go-sql-driver/mysql" // import your required driver
)

func init() {
    // create connection and set default database
    orm.RegisterDataBase("default", "mysql", beego.AppConfig.String("mysqluser")+":"+beego.AppConfig.String("mysqlpass")+"@/"+beego.AppConfig.String("mysqldb")+"?charset=utf8", 30)
}

func main() {
	beego.Run()
}

