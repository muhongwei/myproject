package controllers

import (
	"encoding/json"
	"myproject/models/log"
	"myproject/models/pvc"

	"github.com/astaxie/beego"
	"github.com/golang/glog"
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
	username := sess.Get("username")
	if username == nil {
		c.CustomAbort(400, "login first")
		return
	} else {
		vlog.UserName = interface{}(username).(string)
		vlog.API = "/user/pvc/list"
		vlog.Method = "get"
		log.InsertLog(vlog)

	}

	c.ServeJSON()
}

//CreatePVC CreatePVC
func (c *PVCController) CreatePVC() {
	body := c.Ctx.Input.CopyBody(beego.BConfig.MaxMemory)
	request := &pvc.PVCMessage{}
	marshalerr := json.Unmarshal([]byte(body), request)
	if marshalerr != nil {
		glog.Errorln(marshalerr)
		c.Abort("request body err")

	}
	glog.Errorln(request)
	err := pvc.CreatePVC(request)
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
		vlog.API = "/user/pvc/create"
		vlog.Method = "post"
		log.InsertLog(vlog)

	}

	c.ServeJSON()
}

//GetPVC GetPVC
func (c *PVCController) GetPVC() {
	name := c.GetString(":name")
	a := pvc.GetPVC(name)
	c.Data["json"] = a

	var vlog log.Log
	sess := c.StartSession()
	username := sess.Get("username")
	if username == nil {
		c.CustomAbort(400, "login first")
		return
	} else {
		vlog.UserName = interface{}(username).(string)
		vlog.API = "/user/pvc/get/" + name
		vlog.Method = "get"
		log.InsertLog(vlog)

	}

	c.ServeJSON()
}

//DeletePVC DeletePVC
func (c *PVCController) DeletePVC() {
	name := c.GetString(":name")
	err := pvc.DeletePVC(name)
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
		vlog.API = "/user/pvc/delete/" + name
		vlog.Method = "delete"
		log.InsertLog(vlog)

	}

	c.ServeJSON()
}

//UserListPVC UserListPVC
func (c *PVCController) UserListPVC() {
	c.TplName = "pvclist.html"
}

//UserCreatePVC UserCreatePVC
func (c *PVCController) UserCreatePVC() {
	c.TplName = "pvccreate.html"
}

//UserGetPVC UserGetPVC
func (c *PVCController) UserGetPVC() {
	c.TplName = "pvc.html"
}
