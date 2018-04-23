package main

import (
	_ "myproject/routers"

	"github.com/astaxie/beego"
)

func main() {
	beego.BConfig.WebConfig.Session.SessionOn = true
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.SetStaticPath("/user/assets", "static/assets")
	beego.SetStaticPath("/user/pod/assets", "static/assets")
	beego.SetStaticPath("/user/replicationcontroller/assets", "static/assets")
	beego.SetStaticPath("/user/service/assets", "static/assets")
	beego.SetStaticPath("/user/node/assets", "static/assets")
	beego.SetStaticPath("/user/log/assets", "static/assets")
	beego.SetStaticPath("/user/image/assets", "static/assets")
	beego.SetStaticPath("/user/service/get/assets", "static/assets")
	beego.Run()
}