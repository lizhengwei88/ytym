package main

import (
	"github.com/astaxie/beego"
	"ytym1/controllers"
	"ytym1/models"
	_ "ytym1/routers"
)

func init() {
	models.RegisterDB()
}
func main() {
	// 注册 beego 路由
	beego.Router("/", &controllers.MainController{})
	beego.Router("/workplace", &controllers.WorkController{})
	beego.AutoRouter(&controllers.WorkController{})
	beego.Router("/login", &controllers.LoginController{})
	beego.AutoRouter(&controllers.LoginController{})
	beego.Router("/reg", &controllers.RegController{})
	beego.Router("/ajaxreg", &controllers.RegController{}, "*:AjaxReg")
	beego.AutoRouter(&controllers.RegController{})
	beego.Router("/meinfo", &controllers.MeinfoController{})
	beego.Router("/meinfo/:id([0-9+])", &controllers.MeinfoController{})
	beego.AutoRouter(&controllers.MeinfoController{})
	beego.Router("/prodetail/:id([0-9]+)", &controllers.ProdController{})
	beego.AutoRouter(&controllers.ProdController{})
	beego.Router("/firm_order/:id([0-9]+)", &controllers.FormOrderController{})
	beego.AutoRouter(&controllers.FormOrderController{})
	beego.Router("/pay_success", &controllers.PayController{})
	beego.AutoRouter(&controllers.PayController{})
	beego.Router("/love", &controllers.LoveController{})
	beego.Router("/ajax", &controllers.LoveController{}, "*:Ajax")
	beego.AutoRouter(&controllers.LoveController{})
	beego.Router("/personal", &controllers.PersonalController{})
	beego.AutoRouter(&controllers.PersonalController{})
	beego.Run()
}
