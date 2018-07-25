package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	//"strconv"
	"ytym1/models"
)

type PayController struct {
	beego.Controller
}

func (this *PayController) Get() {
	//uid, _ := strconv.Atoi(this.Ctx.Input.Param(":uname"))
	name := this.Input().Get("uid")
	fmt.Println(name)
	this.Data["Post"] = models.GetPaySuccess(name)
	this.TplName = "pay_success.html"
}
