package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	//"ytym1/models"
)

type PersonalController struct {
	beego.Controller
}

func (this *PersonalController) Get() {
	this.TplName = "personal.html"
	//Personal, err := models.GetAllOrder()
	//if err != nil {
	//	beego.Error(err)
	//}
	//this.Data["Personal"] = Personal
	o := orm.NewOrm()
	var maps []orm.Params
	o.Raw("select ord.id,ord.Cast_id,ord.U_info,ord.create_time,ord.U_dhao,ord.U_name,cust.U_infor,cust.U_pic from cust_order as ord left join custom as cust on ord.Cast_id=cust.id where ord.U_name=?", "张先生").Values(&maps)
	this.Data["Personal"] = maps
	fmt.Println(maps)
}
