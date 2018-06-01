package controllers

import (
	"github.com/astaxie/beego"
)

type AdminController struct {
	beego.Controller
}

func (this *AdminController) Get() {
	this.Data["IsLogin"] = checkAccount(this.Ctx)
	this.TplName = ""
}
