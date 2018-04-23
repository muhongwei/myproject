package controllers

import (
	"fmt"
	"myproject/models/log"
	"myproject/models/service"

	"encoding/json"

	"github.com/astaxie/beego"
	"github.com/golang/glog"
)

//ServiceController ServiceController
type ServiceController struct {
	beego.Controller
}

//ListServices ListServices
func (c *ServiceController) ListServices() {
	a := service.ListServices()
	glog.Infoln(a)
	c.Data["json"] = a

	var vlog log.Log
	sess := c.StartSession()
	username := sess.Get("username")
	if username == nil {
		c.CustomAbort(400, "login first")
		return
	} else {
		vlog.UserName = interface{}(username).(string)
		vlog.API = "/user/service/list"
		vlog.Method = "get"
		log.InsertLog(vlog)

	}

	c.ServeJSON()
}

//GetService GetService
func (c *ServiceController) GetService() {
	name := c.GetString(":name")
	s := service.GetService(name)
	c.Data["json"] = s

	var vlog log.Log
	sess := c.StartSession()
	username := sess.Get("username")
	if username == nil {
		c.CustomAbort(400, "login first")
		return
	} else {
		vlog.UserName = interface{}(username).(string)
		vlog.API = "/user/service/get/" + name
		vlog.Method = "get"
		log.InsertLog(vlog)

	}

	c.ServeJSON()
}

//DeleteService DeleteService
func (c *ServiceController) DeleteService() {
	name := c.GetString(":name")
	err := service.DeleteService(name)
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
		vlog.API = "/user/service/delete"
		vlog.Method = "get"
		log.InsertLog(vlog)

	}

	c.ServeJSON()

}

//CreateService CreateService
func (c *ServiceController) CreateService() {
	body := c.Ctx.Input.CopyBody(beego.BConfig.MaxMemory)
	fmt.Println(string(body))
	request := &service.ServiceInfo{}
	marshalerr := json.Unmarshal([]byte(body), request)
	if marshalerr != nil {
		c.Abort("request body err")
	}
	err := service.CreateService(request)
	if err == nil {
		c.Data["json"] = map[string]string{"result": "ok"}
	} else {
		glog.Errorln(err)
		c.CustomAbort(400, err.Error())
		c.Data["json"] = map[string]string{"result": err.Error()}
	}

	var vlog log.Log
	sess := c.StartSession()
	username := sess.Get("username")
	if username == nil {
		c.CustomAbort(400, "login first")
		return
	} else {
		vlog.UserName = interface{}(username).(string)
		vlog.API = "/user/service/create"
		vlog.Method = "get"
		log.InsertLog(vlog)

	}

	c.ServeJSON()

}

//UserListService UserListService
func (c *ServiceController) UserListService() {
	c.TplName = "services.html"
}

//UserGetService UserGetService
func (c *ServiceController) UserGetService() {
	c.Data["Servicename"] = c.GetString(":name")
	c.TplName = "service.html"
}

//UserCreateService UserCreateService
func (c *ServiceController) UserCreateService() {
	c.TplName = "servicecreate.html"
}
