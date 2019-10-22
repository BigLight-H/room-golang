package admin

import (
	"room/models"
	"room/util"
	"time"
)

type LoginController struct {
	baseController
}
//登陆页面
func (p *LoginController) Index() {
	p.TplName = "login.html"
}

//登陆
func (p *LoginController) Login() {
	name := p.GetString("name")
	pwd := util.Md5(p.GetString("pwd"))
	if name == "" {
		p.MsgBack("用户名不能为空!", 0)
		panic("用户名不能为空!")
	}
	user := []*models.Users{}
	err := p.o.QueryTable(new(models.Users).TableName()).Filter("name", name).Filter("pwd", pwd).One(&user)
	if err != nil {
		p.MsgBack("登录失败!", 0)
	} else {
		p.SetSession("user", user[0])
		p.MsgBack("登陆成功!", 1)
	}

}

//注册
func (p *LoginController) Register() {
	if p.Ctx.Request.Method == "POST" {
		name := p.GetString("name")
		email := p.GetString("email")
		pwd := p.GetString("pwd")
		pwd1 := p.GetString("pwd1")
		num := p.GetString("num")
		if pwd != pwd1 {
			p.MsgBack("密码不一致!", 0)
			panic("go back!")
		}
		if util.VerifyEmailFormat(email)  != true {
			p.MsgBack("邮箱格式不正确!", 0)
			panic("go back!")
		}
		if p.checkUserName(name) != true {
			p.MsgBack("用户名已经存在!", 0)
			panic("go back!")
		}
		if p.checkEmail(email) != true {
			p.MsgBack("邮箱已存在!", 0)
			panic("go back!")
		}
		user := models.Users{}
		user.Name = name
		user.Email = email
		user.Pwd = util.Md5(pwd)
		user.Logo = "/static/img/"+num+".png"
		user.Created = util.DateChange(time.Now().Format("2006年01月02日 15:04:05"))
		user.Updated = util.DateChange(time.Now().Format("2006年01月02日 15:04:05"))
		_, err := p.o.Insert(&user)
		if err != nil {
			p.MsgBack("注册失败!", 0)
			panic("go back!")
		}
		p.MsgBack("注册成功!",1)
	} else {
		p.TplName = "register.html"
	}
}

//验证用户名
func (p *LoginController) checkUserName(name string) bool {
	num,err := p.o.QueryTable(new(models.Users).TableName()).Filter("name", name).Count()
	if err ==nil {
		if num == 0 {
			return true
		}
	}
	return false
}

//验证邮箱
func (p *LoginController) checkEmail(email string) bool {
	num,err := p.o.QueryTable(new(models.Users).TableName()).Filter("email", email).Count()
	if err ==nil {
		if num == 0 {
			return true
		}
	}
	return false
}



