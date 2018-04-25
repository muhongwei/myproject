package controllers

import (
	"encoding/json"
	//"myproject/models/log"
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

	// var vlog log.Log
	// sess := c.StartSession()
	// vlog.UserName = interface{}(sess.Get("username")).(string)
	// vlog.API = "/user/pvc/list"
	// vlog.Method = "get"
	// log.InsertLog(vlog)

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
	c.ServeJSON()
}

//GetPV GetPV
func (c *PVController) GetPV() {
	name := c.GetString(":name")
	a := pv.GetPV(name)
	c.Data["json"] = a
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
	c.ServeJSON()
}
