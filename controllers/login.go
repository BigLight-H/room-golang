package controllers

import "github.com/astaxie/beego"

type LoginController struct {
	beego.Controller
}
//登陆页面
func (p *LoginController) Index() {
	p.TplName = "login.html"
}

//登陆
func (p *LoginController) Login() {
	p.TplName = "login.html"
}

//注册
func (p *LoginController) Register() {

	p.TplName = "register.html"
}

//退出
func (p *LoginController) Logout() {
	p.TplName = "login.html"
}
