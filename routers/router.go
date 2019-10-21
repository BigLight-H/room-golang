package routers

import (
	"room/controllers"
	"github.com/astaxie/beego"
	"room/controllers/admin"
)

func init() {
	//登陆首页
    beego.Router("/", &controllers.LoginController{}, "get:Index")
    beego.Router("/register", &controllers.LoginController{}, "get:Register")
    beego.Router("/register", &controllers.LoginController{}, "post:Register")

    //后台首页
	beego.Router("/admin/index", &admin.AdminController{}, "get:Index")
	beego.Router("/admin/types", &admin.AdminController{}, "get:TypesTpl")
	beego.Router("/admin/rebate", &admin.AdminController{}, "get:RebateTpl")
	beego.Router("/admin/rebate/list", &admin.AdminController{}, "get:RebateList")
	beego.Router("/admin/rebate/update/:id", &admin.AdminController{}, "get:UpdateRebate")
	beego.Router("/admin/rebate/del", &admin.AdminController{}, "post:DelRebate")
	beego.Router("/admin/types/list", &admin.AdminController{}, "get:TypesList")
	beego.Router("/admin/types/add", &admin.AdminController{}, "get:TypesTpl")
	beego.Router("/admin/types/update/?:id", &admin.AdminController{}, "get:UpdateTypes")
	beego.Router("/admin/types/add", &admin.AdminController{}, "post:AddTypes")
	beego.Router("/admin/types/del", &admin.AdminController{}, "post:DelTypes")
	beego.Router("/admin/rebate/add", &admin.AdminController{}, "get:TypesList")
	beego.Router("/admin/rebate/add", &admin.AdminController{}, "post:AddRebate")
}
