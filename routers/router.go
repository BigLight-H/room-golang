package routers

import (
	"room/controllers"
	"github.com/astaxie/beego"
	"room/controllers/admin"
)

func init() {
	//登陆首页
    beego.Router("/", &controllers.LoginController{}, "get:Index")

    //后台首页
	beego.Router("/admin/index", &admin.AdminController{}, "get:Index")
	beego.Router("/admin/types", &admin.AdminController{}, "get:TypesTpl")
	beego.Router("/admin/rebate", &admin.AdminController{}, "get:RebateTpl")
	beego.Router("/admin/rebate/list", &admin.AdminController{}, "get:RebateList")
	beego.Router("/admin/types/list", &admin.AdminController{}, "get:TypesList")
	beego.Router("/admin/types/add", &admin.AdminController{}, "get:TypesTpl")
	beego.Router("/admin/rebate/add", &admin.AdminController{}, "get:TypesList")
	beego.Router("/admin/rebate/add", &admin.AdminController{}, "post:AddRebate")
}
