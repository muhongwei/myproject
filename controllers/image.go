package controllers

import (
	"myproject/models/image"
	"myproject/models/log"
	// "myproject/models/log"

	"github.com/astaxie/beego"
)

//ImageController ImageController
type ImageController struct {
	beego.Controller
}

//ListImages ListImages
func (c *ImageController) ListImages() {
	a := image.ListImages()
	c.Data["json"] = a

	var vlog log.Log
	sess := c.StartSession()
	username := sess.Get("username")
	if username == nil {
		c.CustomAbort(400, "login first")
		return
	} else {
		vlog.UserName = interface{}(username).(string)
		vlog.API = "/user/image/list"
		vlog.Method = "get"
		log.InsertLog(vlog)

	}

	c.ServeJSON()
}

//DelImage DelImage
func (c *ImageController) DelImage() {
	imagename := c.GetString("imagename")
	tag := c.GetString("tag")
	err := image.DelImage(imagename, tag)
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
		vlog.API = "/user/image/delete/" + imagename
		vlog.Method = "delete"
		log.InsertLog(vlog)

	}

	c.ServeJSON()
}

//UserListImage UserListImage
func (c *ImageController) UserListImage() {
	c.TplName = "images.html"
}

//UserPushImage UserPushImage
func (c *ImageController) UserPushImage() {
	c.TplName = "pushimage.html"
}
