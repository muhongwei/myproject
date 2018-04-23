package controllers

import (
	"myproject/models/log"
	"myproject/models/pvc"

	"github.com/astaxie/beego"
)

//PVCController PVCController
type PVCController struct {
	beego.Controller
}

//ListPVC ListPVC
func (c *PVCController) ListPVC() {
	a := pvc.ListPVC()
	c.Data["json"] = a

	var vlog log.Log
	sess := c.StartSession()
	vlog.UserName = interface{}(sess.Get("username")).(string)
	vlog.API = "/user/pvc/list"
	vlog.Method = "get"
	log.InsertLog(vlog)

	c.ServeJSON()
}
