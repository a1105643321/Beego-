package main

import (
	"github.com/astaxie/beego/orm"
	_ "guess/routers"
	"github.com/astaxie/beego"
	_"github.com/go-sql-driver/mysql"
)

func init(){
	orm.RegisterDriver("mysql",orm.DRMySQL)
	orm.RegisterDataBase("default","mysql","root:960815@/subject?charset=utf8")
}

func main() {
	beego.Run()
}

