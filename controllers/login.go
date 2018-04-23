package controllers

import (
	"myproject/models/log"
	"myproject/models/user"

	"github.com/astaxie/beego"
	"github.com/golang/glog"

	"crypto/md5"
	"encoding/hex"
)

type LoginController struct {
	beego.Controller
}

func (c *LoginController) Get() {

	c.TplName = "login.html"
}

//Post Post
func (c *LoginController) Post() {
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

		c.Redirect("/user/pod/list", 302)
		// c.TplName = "pods.html"
	} else {
		//记录用户登录信息
		glog.Infoln("username:", u.UserName)
		glog.Infoln("userpassword:", u.UserPassword)
		glog.Infoln("login failed")
		c.TplName = "error.html"
	}
}
func (c *LoginController) UserInfo() {

	c.TplName = "userinfo.html"
}
