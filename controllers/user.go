package controllers

import (
	"myproject/models/log"
	"myproject/models/user"
	"strconv"

	"github.com/astaxie/beego"
	"github.com/golang/glog"

	"crypto/md5"
	"encoding/hex"
)

//UserController UserController
type UserController struct {
	beego.Controller
}

//Get Get
func (c *UserController) Get() {

	c.TplName = "login.html"
}

//Post Post
func (c *UserController) Post() {
	//将页面和接口调用对应的客户端ip和UserAgent信息作为日志打印到标准控制台
	glog.Infoln("login Post():")
	glog.Infoln("userip:", c.Ctx.Input.IP())
	glog.Infoln("useragent:", c.Ctx.Request.UserAgent())

	//新建models.User类型变量来储存用户登录信息
	var u user.User
	inputs := c.Input()
	u.UserName = inputs.Get("username")
	var password string = inputs.Get("pwd")

	sess := c.StartSession()
	sess.Set("username", u.UserName)

	//将输入密码进行MD5加密与数据库进行对比
	h := md5.New()
	h.Write([]byte(password))
	cipherStr := h.Sum(nil)
	u.UserPassword = hex.EncodeToString(cipherStr)

	err := user.ValidateUser(u)
	if err == nil {
		//记录用户登录信息
		glog.Infoln("username:", u.UserName)
		glog.Infoln("userpassword:", u.UserPassword)
		glog.Infoln("login success")

		var vlog log.Log
		sess := c.StartSession()
		vlog.UserName = interface{}(sess.Get("username")).(string)
		vlog.API = "/user/login"
		vlog.Method = "post"
		log.InsertLog(vlog)

		c.Redirect("/user/info", 302)
		// c.TplName = "pods.html"
	} else {
		//记录用户登录信息
		glog.Infoln("username:", u.UserName)
		glog.Infoln("userpassword:", u.UserPassword)
		glog.Infoln("login failed")
		c.TplName = "error.html"
	}
}
func (this *UserController) RegisterUser() {
	//将页面和接口调用对应的客户端ip和UserAgent信息作为日志打印到标准控制台
	glog.Infoln("regist Post():")
	glog.Infoln("userip:", this.Ctx.Input.IP())
	glog.Infoln("useragent:", this.Ctx.Request.UserAgent())

	var u user.User
	inputs := this.Input()
	u.UserName = inputs.Get("username")
	//将密码进行MD5加密
	var password string = inputs.Get("pwd")
	h := md5.New()
	h.Write([]byte(password))
	cipherStr := h.Sum(nil)
	u.UserPassword = hex.EncodeToString(cipherStr)

	u.UserIntroduction = inputs.Get("introduction")
	//将用户信息保存到数据库，跳回登录页面。保存失败进入err页面
	err := user.SaveUser(u)
	if err != nil {
		glog.Errorln(err)
		this.CustomAbort(400, err.Error())
	}
	this.Redirect("/user/info", 302)

}

//DeleteUser DeleteUser
func (c *UserController) DeleteUser() {
	var u user.User
	var id = c.GetString(":id")
	glog.Infoln("id:" + id)
	var err1 error
	u.UserId, err1 = strconv.Atoi(id)
	if err1 != nil {
		glog.Errorln(err1)
	}
	err := user.DeleteUser(u)
	if err != nil {
		glog.Errorln(err)
	}
	c.Redirect("/user/info", 302)
}

//GetAllUser GetAllUser
func (c *UserController) GetAllUser() {
	users := user.GetAllUser()
	c.Data["json"] = users
	c.ServeJSON()

}

//UserInfo UserInfo
func (c *UserController) UserInfo() {

	c.TplName = "userinfo.html"
}

//UserRegisterPage UserRegisterPage
func (c *UserController) UserRegisterPage() {
	c.TplName = "userregister.html"
}
