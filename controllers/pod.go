package controllers

import (
	//"myproject/models/log"
	"myproject/models/log"
	"myproject/models/pod"

	"github.com/astaxie/beego"
	"github.com/golang/glog"
	//corev1 "k8s.io/api/core/v1"
)

//PodController PodController
type PodController struct {
	beego.Controller
}

//ListPods ListPods
// @Title Get Pod list
// @Description Get Pod list
// @Success 200 {object} corev1.PodList
// @Failure 400 no enough input
// @Failure 500 get pods common error
// @router /list [get]
func (c *PodController) ListPods() {
	a := pod.ListPods()
	c.Data["json"] = a

	// var vlog log.Log
	// sess := c.StartSession()
	// vlog.UserName = interface{}(sess.Get("username")).(string)
	// vlog.API = "/user/pod/list"
	// vlog.Method = "get"
	// log.InsertLog(vlog)

	c.ServeJSON()
}

//GetPod GetPod
// @Title Get Pod
// @Description Get Pod
// @Success 200 {object} corev1.Pod
// @Param   name query   string  true       "name"
// @Param   namespace query   string  true       "namespace"
// @Failure 400 no enough input
// @Failure 500 get pod common error
// @router /get [get]
func (c *PodController) GetPod() {
	podname := c.GetString(":name")
	namespace := c.GetString("namespace")
	pod := pod.GetPod(podname, namespace)
	c.Data["json"] = pod

	var vlog log.Log
	sess := c.StartSession()
	username := sess.Get("username")
	if username == nil {
		c.CustomAbort(400, "login first")
		return
	} else {
		vlog.UserName = interface{}(username).(string)
		vlog.API = "/user/pod/get/" + podname
		vlog.Method = "get"
		log.InsertLog(vlog)

	}
	c.ServeJSON()
}

//DeletePod DeletePod
// @Title Delete Pod
// @Description Delete Pod
// @Success 200 {object} corev1.Pod
// @Param   name query   string  true       "name"
// @Param   namespace query   string  true       "namespace"
// @Failure 400 no enough input
// @Failure 500 delete pod common error
// @router /delete [delete]
func (c *PodController) DeletePod() {
	podname := c.GetString(":name")
	namespace := c.GetString("namespace")
	err := pod.DeletePod(podname, namespace)
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
		vlog.API = "/user/pod/delete/" + podname
		vlog.Method = "get"
		log.InsertLog(vlog)

	}
	c.ServeJSON()
}

//UserGetPod UserGetPod
func (c *PodController) UserGetPod() {
	c.Data["Name"] = c.GetString(":name")
	c.Data["Namespace"] = c.GetString("namespace")
	glog.Infoln(c.GetString(":name"))
	glog.Infoln(c.GetString("namespace"))
	c.TplName = "pod.html"
}

//UserListPod UserListPod
func (c *PodController) UserListPod() {
	c.TplName = "pods.html"
}
