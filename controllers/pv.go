package controllers

import (
	"encoding/json"
	//"myproject/models/log"
	"myproject/models/log"
	"myproject/models/pv"

	"github.com/astaxie/beego"
	"github.com/golang/glog"
)

//PVController PVController
type PVController struct {
	beego.Controller
}

//ListPV ListPV
func (c *PVController) ListPV() {
	a := pv.ListPV()
	c.Data["json"] = a

	var vlog log.Log
	sess := c.StartSession()
	username := sess.Get("username")
	if username == nil {
		c.CustomAbort(400, "login first")
		return
	} else {
		vlog.UserName = interface{}(username).(string)
		vlog.API = "/user/pv/list"
		vlog.Method = "get"
		log.InsertLog(vlog)

	}

	c.ServeJSON()
}

//CreatePV CreatePV
func (c *PVController) CreatePV() {
	body := c.Ctx.Input.CopyBody(beego.BConfig.MaxMemory)
	request := &pv.PVMessage{}
	marshalerr := json.Unmarshal([]byte(body), request)
	if marshalerr != nil {
		glog.Errorln(marshalerr)
		c.Abort("request body err")

	}
	glog.Errorln(request)
	err := pv.CreatePV(request)
	if err == nil {
		c.Data["json"] = map[string]string{"result": "ok"}
	} else {
		c.CustomAbort(400, err.Error())
		return
		//c.Data["json"] = map[string]string{"result": err.Error()}
	}

	var vlog log.Log
	sess := c.StartSession()
	username := sess.Get("username")
	if username == nil {
		c.CustomAbort(400, "login first")
		return
	} else {
		vlog.UserName = interface{}(username).(string)
		vlog.API = "/user/pv/create"
		vlog.Method = "post"
		log.InsertLog(vlog)

	}
	c.ServeJSON()
}

//GetPV GetPV
func (c *PVController) GetPV() {
	name := c.GetString(":name")
	a := pv.GetPV(name)
	c.Data["json"] = a

	var vlog log.Log
	sess := c.StartSession()
	username := sess.Get("username")
	if username == nil {
		c.CustomAbort(400, "login first")
		return
	} else {
		vlog.UserName = interface{}(username).(string)
		vlog.API = "/user/pv/get/" + name
		vlog.Method = "get"
		log.InsertLog(vlog)

	}

	c.ServeJSON()
}

//DeletePV DeletePV
func (c *PVController) DeletePV() {
	name := c.GetString(":name")
	err := pv.DeletePV(name)
	if err == nil {
		c.Data["json"] = map[string]string{"result": "ok"}
	} else {
		c.CustomAbort(400, err.Error())
		return
		//c.Data["json"] = map[string]string{"result": err.Error()}
	}
	var vlog log.Log
	sess := c.StartSession()
	username := sess.Get("username")
	if username == nil {
		c.CustomAbort(400, "login first")
		return
	} else {
		vlog.UserName = interface{}(username).(string)
		vlog.API = "/user/pv/delete/" + name
		vlog.Method = "delete"
		log.InsertLog(vlog)

	}
	c.ServeJSON()
}

//UserListPV UserListPV
func (c *PVController) UserListPV() {
	c.TplName = "pvlist.html"
}

//UserCreatePV UserCreatePV
func (c *PVController) UserCreatePV() {
	c.TplName = "pvcreate.html"
}

//UserGetPV UserGetPV
func (c *PVController) UserGetPV() {
	name := c.GetString(":name")
	c.Data["PVName"] = name
	c.TplName = "pv.html"
}
