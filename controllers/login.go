package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type LoginController struct {
	beego.Controller
}

func (this *LoginController) Get() {
	this.TplName = "login.html"
	var aa = this.GetSession("Uname")
	if aa != nil {
		this.Redirect("/", 302)
	}
}
func (this *LoginController) Post() {
	var maps []orm.Params
	this.TplName = "login.html"
	var name = this.Input().Get("Name")
	var pwd = this.Input().Get("pwd")
	num, err := orm.NewOrm().Raw("SELECT * FROM user WHERE Mobile = ? and password = ?", name, pwd).Values(&maps)
	if err == nil && num > 0 {
		this.Data["json"] = num
		this.SetSession("Uname", maps[0]["Name"])
	} else {
		this.Data["json"] = 0
	}
	this.ServeJSON()
}
