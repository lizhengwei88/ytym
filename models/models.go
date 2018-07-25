package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

func RegisterDB() {
	dbtype := beego.AppConfig.String("db_type")
	dbuser := beego.AppConfig.String("db_user")
	dbPassword := beego.AppConfig.String("db_password")
	dbase := beego.AppConfig.String("db_base")

	//注册默认数据库
	maxIdle := 30
	maxConn := 30
	orm.RegisterDataBase("default", "mysql", dbuser+":"+dbPassword+"@tcp(127.0.0.1:3306)/"+dbase+"?charset=utf8", maxIdle, maxConn)
	//注册数据库，第一个参数必须为"default"
}
func init() {
	// 需要在init中注册定义的model
	orm.RegisterModel(new(Userinfo))
	orm.RegisterModel(new(Custom))
	orm.RegisterModel(new(Cust_order))
	orm.RegisterModel(new(User))
	orm.RegisterModel(new(Article))
	orm.RegisterModel(new(Post))

}

/**
 * account: 短信平台账号
 * pswd: 短信平台密码
 */
const url = beego.AppConfig.String("dxurl")
const account = beego.AppConfig.String("account") 这里填写短信平台账号
const pswd = beego.AppConfig.String("pswd")    // 这里填写短信平台密码

const smUrl = url + "?account=" + account + "&pswd=" + pswd + "&mobile=%s&msg=%s"

/**
 * 发送验证码
 */
func SendMsgToMobile(mobile string, content string) bool {
	strUrl := fmt.Sprintf(smUrl, mobile, content)
	return RemoteCall(strUrl) != nil
}

/**
 * HTTP通信
 */
func RemoteCall(strUrl string) []byte {
	r, err := http.NewRequest("GET", strUrl, nil)
	if err != nil {
		fmt.Println("http.NewRequest: ", err.Error())
		return nil
	}

	// r.Proto = "HTTP/1.0"
	// r.ProtoMajor = 1
	// r.ProtoMinor = 0
	fmt.Println(r.Proto)

	resp, err := http.DefaultClient.Do(r)
	if err != nil {
		fmt.Println("http.DefaultClient.Do: ", err.Error())
		return nil
	}

	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		fmt.Println("resp.StatusCode!=http.StatusOK: ", resp.StatusCode)
		return nil
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil && err != io.EOF {
		fmt.Println("ioutil.ReadAll: ", err.Error())
		return nil
	}

	fmt.Println(string(data))
	return data
}

type User struct {
	Id       int        `json:"id"`
	Name     string     `json:"name"`
	Password string     `json:"password"`
	Nickname string     `json:"nickname"`
	Mobile   string     `json:"mobile"`
	Age      int        `json:"age"`
	Articles []*Article `orm:"reverse(many)"`
}
type Post struct {
	Id    int    `orm:"auto"`
	Title string `orm:"size(100)"`
	User  *User  `orm:"rel(fk)"`
}
type Article struct {
	Id      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
	User    *User  `json:"user" orm:"rel(fk)"`
}

// 分类
type Userinfo struct {
	Id         int64 `orm:"index"`
	U_name     string
	CreateTime time.Time
}

// custom产品表
type Custom struct {
	Id         int64  `orm:"index"`
	U_name     string `json:"custname"`
	U_infor    string
	U_pic      string
	U_price    string
	CreateTime time.Time
}

// cust_order订单表
type Cust_order struct {
	Id          int64 `orm:"index"`
	U_dhao      string
	Cast_id     int
	U_name      string `json:"ordname"`
	U_tel       string
	U_price     string
	U_address   string
	U_info      string
	CreateTime  time.Time
	Cust_orders *Cust_order `orm:"rel(fk)"`
}

func AddUserReg(mobtel, pwd string) {
	fmt.Println(mobtel)
	fmt.Println("结束")
	user := &User{
		Mobile:   mobtel,
		Password: pwd,
	}
	orm.NewOrm().Insert(user)

}
func OrderAdd(cust_order Cust_order) (cst Cust_order) {
	o := orm.NewOrm()
	//插入分类数据
	_, err := o.Insert(&cust_order)
	if err != nil {
		fmt.Println(err)
	}
	return cst
}
func AddCategory(name string) error {
	o := orm.NewOrm()
	cate := &Userinfo{
		U_name:     name,
		CreateTime: time.Now(),
	}
	//先查询表中是否有相同分类
	//****************?******************
	//r := o.Raw("select count(*) from userinfo WHERE name = ?", name)
	//先查询表中是否有相同分类
	qs := o.QueryTable("userinfo")
	fmt.Println(cate)
	err := qs.Filter("U_name", name).One(cate)
	fmt.Println("(")
	fmt.Println(err)
	fmt.Println(")")
	if err == nil {
		return err
	}
	//插入分类数据
	_, err = o.Insert(cate)
	if err != nil {
		return err
	}
	return nil
}
func GetPaySuccess(name string) (cust_order Cust_order) {
	//	userinfo := new(Userinfo)
	fmt.Println(name)
	fmt.Println("id")
	err := orm.NewOrm().QueryTable("cust_order").Filter("id", name).Limit(1).One(&cust_order)
	if err == nil {
		fmt.Println(cust_order)
	}
	return
}
func GetFirm_order(id int) (custom Custom) {
	//	userinfo := new(Userinfo)
	fmt.Println("id")
	err := orm.NewOrm().QueryTable("custom").Filter("id", id).Limit(1).One(&custom)
	if err == nil {
		fmt.Println(custom)
	}
	return
}
func GetAllCustom(isDesc bool) (custom []*Custom, err error) {
	o := orm.NewOrm()
	custom = make([]*Custom, 0)
	qs := o.QueryTable("custom")
	if isDesc {
		_, err = qs.OrderBy("id").All(&custom)
	} else {
		_, err = qs.All(&custom)
	}
	return custom, err
}
func GetAllOrder() (cust_order []*Cust_order, err error) {
	o := orm.NewOrm()

	num, err := o.Raw("select * from cust_order").QueryRows(&cust_order)
	fmt.Println(cust_order)
	if err == nil {
		fmt.Println("user nums: ", num)
	}

	return cust_order, err
}

type Userord struct {
	Id       int
	UserName string
}

func OrderIdShow(id int) (cust_order Cust_order) {
	//	err := orm.NewOrm().QueryTable("cust_order").Filter("id", id).Limit(1).One(&cust_order)
	//orm.NewOrm().QueryTable("cust_order").Filter("id", id).RelatedSel().All(&cust_order)
	//orm.NewOrm().QueryTable("cust_order").Filter("Custom", 1).RelatedSel().All(&cust_order)
	o := orm.NewOrm()
	//o.Raw("SELECT cust_order.U_name,custom.U_pic FROM cust_order,custom WHERE cust_order.cast_id=custom.id and custom.id = ?", 1).QueryRow(&cust_order)
	fmt.Println("id")
	//	err := orm.NewOrm().QueryTable("cust_order").Filter("id", id).Limit(1).One(&cust_order)
	//err := orm.NewOrm().Raw("SELECT cust_order.U_name,custom.U_pic FROM cust_order,custom WHERE cust_order.cast_id=custom.id and custom.id = ?", 1).QueryRow(&cust_order)
	var maps []orm.Params
	num, err := o.Raw("SELECT user.Name,custom.U_pic FROM user,custom WHERE user.id=custom.id and user.id = ?", 1).Values(&maps)
	if err == nil && num > 0 {
		fmt.Println(maps[0]["Name"])  // slene
		fmt.Println(maps[0]["U_pic"]) // slene
	}
	fmt.Println(num)
	return
}
func OrderShow(id int) (cust_order Cust_order) {
	o := orm.NewOrm()
	var maps []orm.Params
	num, err := o.Raw("SELECT user.Name,custom.U_pic FROM user,custom WHERE user.id=custom.id and user.id = ?", id).Values(&maps)
	if err == nil && num > 0 {
		fmt.Println(maps[0]["Name"])  // slene
		fmt.Println(maps[0]["U_pic"]) // slene
	}
	fmt.Println(maps)
	return
}

func OrderSh(id int) (users User) {

	//var users []User
	// 获取 QueryBuilder 对象. 需要指定数据库驱动参数。
	// 第二个返回值是错误对象，在这里略过
	qb, _ := orm.NewQueryBuilder("mysql")
	// 构建查询对象
	//qb.Select("user.Name",
	//	"post.Title").
	//	From("user").
	//	InnerJoin("post").On("user.id=post.user_id").
	//	Where("post.id=1")
	//OrderBy("id").Desc().
	//Limit(10).Offset(0)
	qb.Select("user.name",
		"post.Title").
		From("user", "post").
		Where("user.id=post.user_id and post.id=1")
	// 导出 SQL 语句
	sql := qb.String()
	fmt.Println(sql)
	// 执行 SQL 语句
	o := orm.NewOrm()
	o.Raw(sql).QueryRow(&users)
	//fmt.Println(users)
	err := o.Raw("SELECT user.name, post.Title FROM user,post WHERE user.id=post.user_id and post.id = ?", 1).QueryRow(&users)
	if err == nil {
		fmt.Println(users)
	}
	return users
}
