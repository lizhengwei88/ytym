package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"strconv"
	//"ytym1/models"
)

type MeinfoController struct {
	beego.Controller
}

func (this *MeinfoController) Get() {
	this.TplName = "meinfo.html"
	id, _ := strconv.Atoi(this.Ctx.Input.Param(":id"))

	//var articles []*models.Article
	//ords, err := orm.NewOrm().QueryTable("article").Filter("User", 1).RelatedSel().All(&articles)
	//if err != nil {
	//	beego.Error(err)
	//}

	var posts []orm.Params
	orm.NewOrm().Raw("SELECT cus.U_name,cus.U_infor,cus.U_pic,ord.U_name as ordname,ord.U_tel,ord.U_address,ord.U_info,ord.U_price FROM cust_order as ord,custom as cus WHERE ord.cast_id=cus.id and ord.id = ?", id).Values(&posts)
	fmt.Println(posts)
	this.Data["Ordss"] = posts

}
