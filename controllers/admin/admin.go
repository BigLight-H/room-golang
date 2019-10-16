package admin

import (
	"room/models"
	"strconv"
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
					str += "<option value='"+ strconv.FormatFloat(value.Money, 'E', -1, 64) +"'>"+ value.Type +"</option>"
				}
			}
			str += "</optgroup>"
		}
	}
	a.Data["select"] = str
	a.TplName = "admin/index.html"
}
