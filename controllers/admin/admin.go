package admin

import (
	"room/models"
	"room/util"
	"strconv"
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
	rebate := []*models.Rebate{}
	a.o.QueryTable(new(models.Rebate).TableName()).All(&rebate)
	a.Data["rebate"] = rebate
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
	list := []*models.Types{}
	a.o.QueryTable(new(models.Types).TableName()).All(&list)
	a.Data["list"] = list
	a.TplName = "admin/types-list.html"
}

//添加类型
func (a *AdminController) AddTypes() {
	ids := a.GetString("ids")
	title := a.GetString("title")
	money := a.GetString("money")
	sel := a.GetString("sel")
	types := models.Types{}
	types.Money = money
	types.Type = title
	if sel != "" {
		types.Pid = util.StrToInt(sel)
	} else {
		types.Pid = 0
	}
	if ids != "" {
		old_t := models.Types{Id:util.StrToInt(ids)}
		old_t.Path = types.Path
		old_t.Money = types.Money
		old_t.Type = types.Type
		old_t.Pid = types.Pid
		_,err := a.o.Update(&old_t)
		if err == nil {
			t := models.Types{Id:util.StrToInt(ids)}
			t.Path = a.getTypePath(types.Pid)+","+ids
			_,err := a.o.Update(&t,"Path")
			if err == nil {
				a.Data["json"] = 1
				a.ServeJSON()
			}
		}
	} else {
		id,err := a.o.Insert(&types)
		if err == nil {
			t := models.Types{Id:util.Int64ToInt(id)}
			t.Path = a.getTypePath(types.Pid)+","+strconv.FormatInt(id,10)
			_,err := a.o.Update(&t,"Path")
			if err == nil {
				a.Data["json"] = 1
				a.ServeJSON()
			}
		}
	}
	a.Data["json"] = 0
	a.ServeJSON()
}

func (a *AdminController) getTypePath(pid int) string {
	if pid > 0 {
		t := models.Types{Id:pid}
		err := a.o.Read(&t)
		if err == nil {
			return t.Path
		}
	}
	return "0"
}

//修改类型
func (a *AdminController) UpdateTypes() {
	str := ""
	id := util.StrToInt(a.GetString(":id"))
	t := models.Types{Id:id}
	a.o.Read(&t)
	a.Data["data"] = t
	types := []*models.Types{}
	a.o.QueryTable(new(models.Types).TableName()).All(&types)
	a.Data["types"] = types
	for _, val := range types {
		if val.Pid == 0 {
			if val.Id == t.Pid {
				str += "<option value='"+ util.IntToStr(val.Id) +"' selected>"+ val.Type +"</option>"
			} else {
				str += "<option value='"+ util.IntToStr(val.Id) +"'>"+ val.Type +"</option>"
			}
			for _, value := range types {
				if val.Id == value.Pid {
					if value.Id == t.Pid {
						str += "<option value='"+ util.IntToStr(value.Id) +"' selected>&nbsp;&nbsp;&nbsp;&nbsp;"+ value.Type +"</option>"
					} else {
						str += "<option value='"+ util.IntToStr(value.Id) +"'>&nbsp;&nbsp;&nbsp;&nbsp;"+ value.Type +"</option>"
					}
				}
			}
		}
	}
	a.Data["str"] = str
	a.TplName = "admin/update-types.html"
}

//删除类型
func (a *AdminController) DelTypes() {
	id := a.GetString("id")
	_,err := a.o.Raw("delete from md_types where id = "+id+" or pid = "+id).Exec()
	if err != nil {
		a.Data["json"] = 0
	} else {
		a.Data["json"] = 1
	}
	a.ServeJSON()
}

//活动页面
func (a *AdminController) RebateTpl() {
	a.TplName = "admin/rebate.html"
}

//添加活动
func (a *AdminController) AddRebate() {
	ids := a.GetString("ids")
	title := a.GetString("title")
	money := a.GetString("money")
	date_star := a.GetString("date_star")
	date_end := a.GetString("date_end")
	moneys := a.GetString("push_moneys")
	rebate := models.Rebate{}
	rebate.Title = title
	rebate.TallyMoney = money
	rebate.Rebate = moneys
	rebate.StartDate = util.DateChange(date_star+" 00:00:00")
	rebate.EndDate = util.DateChange(date_end+" 23:59:59")
	rebate.Created = util.DateChange(time.Now().Format("2006年01月02日 15:04:05"))
	rebate.Updated = util.DateChange(time.Now().Format("2006年01月02日 15:04:05"))
	if ids !="" {
		r := models.Rebate{Id:util.StrToInt(ids)}
		r.Title = title
		r.TallyMoney = money
		r.Rebate = moneys
		r.StartDate = util.DateChange(date_star+" 00:00:00")
		r.EndDate = util.DateChange(date_end+" 23:59:59")
		_, err := a.o.Update(&r)
		if err != nil {
			a.Data["json"] = 0
		} else {
			a.Data["json"] = 1
		}
	} else {
		_, err := a.o.Insert(&rebate)
		if err != nil {
			a.Data["json"] = 0
		} else {
			a.Data["json"] = 1
		}
	}
	a.ServeJSON()
}

//活动列表
func (a *AdminController) RebateList() {
	list := []*models.Rebate{}
	a.o.QueryTable(new(models.Rebate).TableName()).All(&list)
	a.Data["list"] = list
	a.TplName = "admin/rebate-list.html"
}

//删除活动
func (a *AdminController) DelRebate() {
	id := a.GetString("id")
	r := models.Rebate{Id:util.StrToInt(id)}
	_,err := a.o.Delete(&r)
	if err != nil {
		a.Data["json"] = 0
	} else {
		a.Data["json"] = 1
	}
	a.ServeJSON()
}

//修改活动
func (a *AdminController) UpdateRebate() {
	id := a.GetString(":id")
	r := models.Rebate{Id:util.StrToInt(id)}
	a.o.Read(&r)
	a.Data["data"] = r
	a.TplName = "admin/update-rebate.html"
}

//退出
func (p *AdminController) Logout() {
	p.DestroySession()
	p.History("退出登录", "/")
}