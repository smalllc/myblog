package controllers

import (
	"beeblog/models"

	"github.com/astaxie/beego"
)

type HomeController struct {
	beego.Controller
}

func (this *HomeController) Get() {
	this.Data["IsHome"] = true
	this.Data["IsLogin"] = checkAccount(this.Ctx)
	this.TplName = "home.html"

	topics, err := models.GetAllTopics(
		this.Input().Get("cate"), this.Input().Get("lable"), true)
	if err != nil {
		beego.Error(err)
	}
	this.Data["Topics"] = topics

	categories, err := models.GetAllCategories()
	if err != nil {
		beego.Error(err)
	}
	this.Data["Categories"] = categories
}
