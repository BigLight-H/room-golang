package admin

import (
	"room/models"
	"room/util"
	"time"
)

type AdminController struct {
	baseController
}

func (a *AdminController) Index() {
	str := ""
	types := []*models.Types{}
	a.o.QueryTable(new(models.Types).TableName()).All(&types)
	a.Data["types"] = types
	for _, val := range types {
		if val.Pid == 0 {
			str += "<optgroup label="+ val.Type +">"
			for _, value := range types {
				if val.Id == value.Pid {
					str += "<option value='"+ value.Money +"'>"+ value.Type +"</option>"
				}
			}
			str += "</optgroup>"
		}
	}
	a.Data["select"] = str
	a.TplName = "admin/index.html"
}

//类型页面
func (a *AdminController) TypesTpl() {
	types := []*models.Types{}
	a.o.QueryTable(new(models.Types).TableName()).All(&types)
	a.Data["types"] = types
	a.TplName = "admin/types.html"
}

//类型列表
func (a *AdminController) TypesList() {
	a.TplName = "admin/types-list.html"
}

//添加类型
func (a *AdminController) AddTypes() {

}

//修改类型
func (a *AdminController) UpdateTypes() {

}

//删除类型
func (a *AdminController) DelTypes() {

}

//活动页面
func (a *AdminController) RebateTpl() {
	a.TplName = "admin/form_plugins.html"
}

//添加活动
func (a *AdminController) AddRebate() {
	title := a.GetString("title")
	money := a.GetString("money")
	date_star := a.GetString("date_star")
	date_end := a.GetString("date_end")
	rebate := models.Rebate{}
	rebate.Title = title
	rebate.TallyMoney = money
	rebate.StartDate = util.DateChange(date_star+" 00:00:00")
	rebate.EndDate = util.DateChange(date_end+" 23:59:59")
	rebate.Created = util.DateChange(time.Now().Format("2006年01月02日 15:04:05"))
	rebate.Updated = util.DateChange(time.Now().Format("2006年01月02日 15:04:05"))
	_, err := a.o.Insert(&rebate)
	if err != nil {
		a.Data["json"] = 0
		a.ServeJSON()
	} else {
		a.Data["json"] = 1
		a.ServeJSON()
	}
}

//添加列表
func (a *AdminController) RebateList() {
	a.TplName = "admin/types-list.html"
}

//删除活动
func (a *AdminController) DelRebate() {

}

//修改活动
func (a *AdminController) UpdateRebate() {

}
