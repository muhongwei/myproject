package controllers

import (
	"fmt"
	"myproject/models/log"

	"github.com/astaxie/beego"
)

type LogController struct {
	beego.Controller
}

func (c *LogController) ListLog() {
	c.TplName = "loglist.html"
}
func (c *LogController) ListLogJson() {
	logs := log.ListLog()
	c.Data["json"] = logs
	var vlog log.Log
	sess := c.StartSession()
	username := sess.Get("username")
	if username == nil {
		c.CustomAbort(400, "login first")
		return
	} else {
		vlog.UserName = interface{}(username).(string)
		vlog.API = "/user/log/list"
		vlog.Method = "get"
		log.InsertLog(vlog)

	}

	c.ServeJSON()

}
func (c *LogController) Loglog() {
	fmt.Println("logcontroler")
}
