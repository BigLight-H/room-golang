package admin

type AdminController struct {
	baseController
}

func (a *AdminController) Index() {
	a.TplName = "admin/index.html"
}
