package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	//"strings"
	//"net/http"
	"ytym1/models"
)

type LoveController struct {
	beego.Controller
}

func (c *LoveController) Ajax() {
	var users []models.Custom
	orm.Debug = true //打开查询日志
	o := orm.NewOrm()
	o.Using("default") // 默认使用 default
	num, err := o.Raw("select *from custom").QueryRows(&users)
	if err != nil {
		fmt.Println("user nums: ", err)
	}
	fmt.Println(num)
	orm.NewOrm().QueryTable("custom").All(&users)
	for _, v := range users {
		fmt.Println(v)
		//return v
	}
	//	c.Data["json"] = &users
	c.Data["json"] = `[{"cpname":"cxh","sex":"man"},{"cpname":"cxh1","sex":"man1"}]`
	c.ServeJSON()
	fmt.Println("开始")
	fmt.Println(users)
	fmt.Println("结束")
}
func (this *LoveController) Get() {
	this.TplName = "love.html"
}
