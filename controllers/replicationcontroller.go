package controllers

import (
	"fmt"
	"myproject/models/log"
	"myproject/models/replicationcontroller"

	"encoding/json"

	"github.com/astaxie/beego"
	"github.com/golang/glog"
)

//RcController RcController
type RcController struct {
	beego.Controller
}

//ListRcs ListRcs
func (c *RcController) ListRcs() {
	a := replicationcontroller.ListReplicationControllers()
	c.Data["json"] = a

	var vlog log.Log
	sess := c.StartSession()
	username := sess.Get("username")
	if username == nil {
		c.CustomAbort(400, "login first")
		return
	} else {
		vlog.UserName = interface{}(username).(string)
		vlog.API = "/user/replicationcontroller/list"
		vlog.Method = "get"
		log.InsertLog(vlog)

	}

	c.ServeJSON()
}

//GetRc GetRc
func (c *RcController) GetRc() {
	name := c.GetString(":name")
	a := replicationcontroller.GetReplicationController(name)
	c.Data["json"] = a

	var vlog log.Log
	sess := c.StartSession()
	username := sess.Get("username")
	if username == nil {
		c.CustomAbort(400, "login first")
		return
	} else {
		vlog.UserName = interface{}(username).(string)
		vlog.API = "/user/replicationcontroller/get"
		vlog.Method = "get"
		log.InsertLog(vlog)

	}

	c.ServeJSON()
}

//DeleteRc DeleteRc
func (c *RcController) DeleteRc() {
	name := c.GetString(":name")
	namespace := c.GetString("namespace")
	fmt.Println(namespace)
	err := replicationcontroller.DeleteReplicationController(name, namespace)
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
		vlog.API = "/user/replicationcontroller/delete"
		vlog.Method = "get"
		log.InsertLog(vlog)

	}

	c.ServeJSON()
}

//CreateRc CreateRc
func (c RcController) CreateRc() {
	body := c.Ctx.Input.CopyBody(beego.BConfig.MaxMemory)
	request := &replicationcontroller.RcMessage{}
	glog.Errorln(string(body))
	marshalerr := json.Unmarshal([]byte(body), request)
	if marshalerr != nil {
		glog.Errorln(marshalerr)
		c.Abort("request body err")

	}
	glog.Errorln(request)

	err := replicationcontroller.CreateReplicationControllers(request)
	if err == nil {
		c.Data["json"] = map[string]string{"result": "ok"}
	} else {
		c.CustomAbort(400, err.Error())
		return
		//c.Data["json"] = map[string]string{"result": err.Error()}
	}

	// var vlog log.Log
	// sess := c.StartSession()
	// username := sess.Get("username")
	// if username == nil {
	// 	c.CustomAbort(400, "login first")
	// 	return
	// } else {
	// 	vlog.UserName = interface{}(username).(string)
	// 	vlog.API = "/user/replicationcontroller/create"
	// 	vlog.Method = "get"
	// 	log.InsertLog(vlog)

	// }

	c.ServeJSON()
}

//UserListRC UserListRC
func (c *RcController) UserListRC() {
	c.TplName = "replicationcontrollerlist.html"
}

//UserCreateRC UserCreateRC
func (c *RcController) UserCreateRC() {
	c.TplName = "replicationcontrollercreate.html"
}

//UserGetRC UserGetRC
func (c *RcController) UserGetRC() {
	name := c.GetString(":name")
	c.Data["RCName"] = name
	c.TplName = "replicationcontroller.html"
}
