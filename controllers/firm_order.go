package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"strconv"
	"time"
	"ytym1/models"
)

type FormOrderController struct {
	beego.Controller
}

func (this *FormOrderController) Get() {
	if this.GetSession("Uname") == nil {
		this.Redirect("/login", 302)
	}
	this.TplName = "firm_order.html"
	uid, _ := strconv.Atoi(this.Ctx.Input.Param(":id"))
	//name := this.Input().Get("uname")
	fmt.Println(uid)
	this.Data["Post"] = models.GetFirm_order(uid)

	op := this.Input().Get("op")
	switch op {
	case "add":
		name := this.Input().Get("sure")
		fmt.Println(name)
		if len(name) == 0 {
			break
		}
		err := models.AddCategory(name)
		if err != nil {
			beego.Error(err)
		}
		this.Redirect("/pay_success?uname="+name, 302)

		return
	case "del":
		id := this.Input().Get("id")
		if len(id) == 0 {
			break
		}

	}
}
func (this *FormOrderController) Post() {
	fmt.Println("保存订单")
	var custorder models.Cust_order
	id := this.Input().Get("castid")
	intid, err := strconv.Atoi(id)
	if err != nil {

	}

	custorder.Cast_id = intid
	custorder.U_name = this.Input().Get("username")
	custorder.U_tel = this.Input().Get("usertel")
	custorder.U_address = this.Input().Get("useraddress")
	custorder.U_price = this.Input().Get("userprice")
	custorder.U_info = this.Input().Get("sure")
	custorder.CreateTime = time.Now()
	models.OrderAdd(custorder)
	Udanhoa := "32323"
	this.Redirect("/pay_success?uid="+Udanhoa, 302)
	//this.Ctx.Redirect(302, "/")

}
