package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"math/rand"
	"strconv"
	"time"
	"ytym1/models"
)

type RegController struct {
	beego.Controller
}

func (this *RegController) Get() {
	this.TplName = "reg.html"
	abc := this.Ctx.Input.Param("mobTel")
	fmt.Println(abc)
	const mobile = "18911786075" // 这里填写手机号码
	//var result = models.SendMsgToMobile(mobile, "验证码")
	//fmt.Println(result)

}
func (this *RegController) AjaxReg() {
	mobTel := this.Input().Get("mobTel")
	fmt.Println(mobTel)
	cnt, _ := orm.NewOrm().QueryTable("user").Filter("Mobile", mobTel).Count() // SELECT COUNT(*) FROM USER
	if cnt != 0 {
		this.Data["json"] = 1
	} else {
		this.Data["json"] = 0
		rand.Seed(time.Now().Unix())
		rnd := rand.Intn(9999)
		this.SetSession("ma", rnd)
		var c = "注册验证码:" + strconv.Itoa(rnd) + ";10分钟内有效。"
		var result = models.SendMsgToMobile(mobTel, c)
		fmt.Println(result)
	}
	this.ServeJSON()

	//this.Data["json"] = `[{"cpname":"cxh1","sex":"man"},{"cpname":"cxh1","sex":"man1"}]`
	//this.Data["json"] = 1

}
func (this *RegController) Post() {
	this.TplName = "reg.html"
	mobtel := this.Input().Get("mobTel")
	pwd := this.Input().Get("password2")

	models.AddUserReg(mobtel, pwd)

	this.Redirect("/login", 302)

}
