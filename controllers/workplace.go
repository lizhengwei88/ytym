package controllers

import (
	"github.com/astaxie/beego"
	//"ytym/models"
)

type WorkController struct {
	beego.Controller
}

func (this *WorkController) Get() {
	this.TplName = "workplace.html"
}
