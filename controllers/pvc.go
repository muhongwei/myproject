package controllers

import (
	"encoding/json"
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

	// var vlog log.Log
	// sess := c.StartSession()
	// vlog.UserName = interface{}(sess.Get("username")).(string)
	// vlog.API = "/user/pvc/list"
	// vlog.Method = "get"
	// log.InsertLog(vlog)

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
	c.ServeJSON()
}

//GetPVC GetPVC
func (c *PVCController) GetPVC() {
	name := c.GetString(":name")
	a := pvc.GetPVC(name)
	c.Data["json"] = a
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
	c.ServeJSON()
}
