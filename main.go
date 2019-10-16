package main

import (
	"github.com/astaxie/beego/orm"
	"room/models"
	_ "room/routers"
	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
)

func init()  {
	orm.Debug = true
	models.Init()
	beego.SetStaticPath("/static","static/")
}


func main() {
	beego.Run()
}

