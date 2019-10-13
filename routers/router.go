package routers

import (
	"room/controllers"
	"github.com/astaxie/beego"
	"room/controllers/admin"
)

func init() {
    beego.Router("/", &controllers.LoginController{}, "get:Index")

    //后台首页
	beego.Router("/admin/index", &admin.AdminController{}, "get:Index")
}
