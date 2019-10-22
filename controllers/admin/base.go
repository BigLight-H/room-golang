package admin

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"room/models"
	"strings"
)

type baseController struct {
	beego.Controller
	o orm.Ormer
	controllerName string
	actionName     string
}

//Json结构体
type Json struct {
	Msg string
	Status int
}

//侧边栏结构体
type Sidebar struct {
	Name string
	Url string
	Font string
	Active string
}

//构造函数
func (p *baseController) Prepare()  {
	controllerName, actionName := p.GetControllerAndAction()
	p.controllerName = strings.ToLower(controllerName[0 : len(controllerName)-10])
	p.actionName = strings.ToLower(actionName)
	p.o = orm.NewOrm()
	if p.controllerName == "admin"{
		if p.GetSession("user") == nil {
			p.History("", "/")
		} else {
			p.GetSidebar()
			p.Data["user_data"] = p.GetSession("user")
		}
	}

}

//跳转重定向
func (p *baseController) History(msg string, url string)  {
	if url == "" {
		p.Ctx.WriteString("<script>alert('"+msg+"');window.history.go(-1);</script>")
		p.StopRun()
	} else {
		p.Redirect(url,302)
	}
}

//返回json信息
func (p *baseController) MsgBack(msg string, status int)  {
	data := &Json{
		msg,
		status,
	}
	p.Data["json"] = data
	p.ServeJSON()
}

//侧边栏展示
func (p *baseController) GetSidebar() {
	url := p.Ctx.Request.RequestURI//获取当前页路由
	num, _ :=p.o.QueryTable(new(models.Sidebar).TableName()).Filter("url_path", url).Count()
	if num > 0 {
		sidebar := []*models.Sidebar{}
		side := p.GetSession("sidebar_list")
		if side == nil {
			p.o.QueryTable(new(models.Sidebar).TableName()).All(&sidebar)
			p.SetSession("sidebar_list", sidebar)
			//遍历出侧边栏
			tags := make(map[int]interface{})
			for k, v := range sidebar {
				s := &Sidebar{
					Name:v.TabName,
					Url:v.UrlPath,
					Font:v.FontClass,
					Active:"",
				}
				if v.UrlPath == url {
					s.Active = "active"
				}
				tags[k] = s
			}
			p.Data["sidebar"] = tags
			p.SetSession("sidebar",tags)
		} else {
			//遍历出侧边栏
			tags := make(map[int]interface{})
			for k, v := range side.([]*models.Sidebar) {
				s := &Sidebar{
					Name:v.TabName,
					Url:v.UrlPath,
					Font:v.FontClass,
					Active:"",
				}
				if v.UrlPath == url {
					s.Active = "active"
				}
				tags[k] = s
			}
			p.Data["sidebar"] = tags
			p.SetSession("sidebar",tags)
		}
	} else {
		p.Data["sidebar"] = p.GetSession("sidebar")
	}
}