package controllers

import (
	"myproject/models/log"
	"myproject/models/node"

	"github.com/astaxie/beego"
)

//NodeController NodeController
type NodeController struct {
	beego.Controller
}

//ListNodes ListNodes
func (c *NodeController) ListNodes() {
	a := node.ListNode()
	c.Data["json"] = a

	var vlog log.Log
	sess := c.StartSession()
	username := sess.Get("username")
	if username == nil {
		c.CustomAbort(400, "login first")
		return
	} else {
		vlog.UserName = interface{}(username).(string)
		vlog.API = "/user/node/list"
		vlog.Method = "get"
		log.InsertLog(vlog)

	}

	c.ServeJSON()
}

//GetNode GetNode
func (c *NodeController) GetNode() {
	nodeName := c.GetString(":nodename")
	node := node.GetNode(nodeName)
	c.Data["json"] = node

	var vlog log.Log
	sess := c.StartSession()
	username := sess.Get("username")
	if username == nil {
		c.CustomAbort(400, "login first")
		return
	} else {
		vlog.UserName = interface{}(username).(string)
		vlog.API = "/user/node/get/" + nodeName
		vlog.Method = "get"
		log.InsertLog(vlog)

	}

	c.ServeJSON()

}

//UserListNode UserListNode
func (c *NodeController) UserListNode() {
	c.TplName = "nodes.html"
}

//UserGetNode UserGetNode
func (c *NodeController) UserGetNode() {
	nodename := c.GetString(":name")
	c.Data["Name"] = nodename
	c.TplName = "node.html"
}
