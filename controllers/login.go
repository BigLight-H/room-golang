package controllers

import "github.com/astaxie/beego"

type LoginController struct {
	beego.Controller
}

func (p *LoginController) Index() {
	p.TplName = "login.html"
}

