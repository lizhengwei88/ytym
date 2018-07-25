package controllers

import (
	"github.com/astaxie/beego"
	//"ytym/models"
	//"fmt"
	"strconv"
)

type ProdController struct {
	beego.Controller
}

func (this *ProdController) Get() {
	id, _ := strconv.Atoi(this.Ctx.Input.Param(":id"))

	//sthis.Redirect("/prodetail?id="+i, 302)
	this.Data["uid"] = id
	this.TplName = "prodetail.html"
}
