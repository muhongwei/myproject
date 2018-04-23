// @APIVersion 1.0.0
// @Title mobile API
// @Description mobile has every tool to get any job done, so codename for the new mobile APIs.
// @Contact astaxie@gmail.com
package routers

import (
	"myproject/controllers"

	"github.com/astaxie/beego"
)

func init() {
	// ns :=
	//     beego.NewNamespace("/v1",
	//         beego.NSNamespace("/pod",
	//             beego.NSInclude(
	//                 &controllers.PodController{},
	//             ),
	//         ),
	//         // beego.NSNamespace("/catalog",
	//         //     beego.NSInclude(
	//         //         &controllers.CatalogController{},
	//         //     ),
	//         // ),
	//         // beego.NSNamespace("/newsletter",
	//         //     beego.NSInclude(
	//         //         &controllers.NewsLetterController{},
	//         //     ),
	//         // ),
	//         // beego.NSNamespace("/cms",
	//         //     beego.NSInclude(
	//         //         &controllers.CMSController{},
	//         //     ),
	//         // ),
	//         // beego.NSNamespace("/suggest",
	//         //     beego.NSInclude(
	//         //         &controllers.SearchController{},
	//         //     ),
	//         // ),
	//     )
	// beego.AddNamespace(ns)
	beego.Router("/", &controllers.MainController{})

	//User User
	beego.Router("/user/login", &controllers.LoginController{})
	beego.Router("/user/info", &controllers.LoginController{}, "get:UserInfo")

	beego.Router("/user/node/list", &controllers.NodeController{}, "get:UserListNode")
	beego.Router("/user/node/:name", &controllers.NodeController{}, "get:UserGetNode")

	beego.Router("/user/pod/list", &controllers.PodController{}, "get:UserListPod")
	beego.Router("/user/pod/:name", &controllers.PodController{}, "get:UserGetPod")

	beego.Router("/user/replicationcontroller/list", &controllers.RcController{}, "get:UserListRC")
	beego.Router("/user/replicationcontroller/create", &controllers.RcController{}, "get:UserCreateRC")
	beego.Router("/user/replicationcontroller/:name", &controllers.RcController{}, "get:UserGetRC")

	beego.Router("/user/service/list", &controllers.ServiceController{}, "get:UserListService")
	beego.Router("/user/service/create", &controllers.ServiceController{}, "get:UserCreateService")
	beego.Router("/user/service/get/:name", &controllers.ServiceController{}, "get:UserGetService")

	beego.Router("/user/log/list", &controllers.LogController{}, "get:ListLog")

	beego.Router("/user/image/list", &controllers.ImageController{}, "get:UserListImage")
	beego.Router("/user/image/push", &controllers.ImageController{}, "get:UserPushImage")

	//Log Log
	beego.Router("/log/list", &controllers.LogController{}, "get:ListLogJson")

	//Node Node
	beego.Router("/node/list", &controllers.NodeController{}, "get:ListNodes")
	beego.Router("/node/:nodename", &controllers.NodeController{}, "get:GetNode")

	//Pod Pod
	beego.Router("/pod/list", &controllers.PodController{}, "get:ListPods")
	beego.Router("/pod/:name", &controllers.PodController{}, "get:GetPod")
	beego.Router("/pod/delete/:name", &controllers.PodController{}, "delete:DeletePod")

	//Service Service
	beego.Router("/service/list", &controllers.ServiceController{}, "get:ListServices")
	beego.Router("/service/:name", &controllers.ServiceController{}, "get:GetService")
	beego.Router("/service/delete/:name", &controllers.ServiceController{}, "delete:DeleteService")
	beego.Router("/service/create", &controllers.ServiceController{}, "post:CreateService")

	//ReplicationController ReplicationController
	beego.Router("/replicationcontroller/list", &controllers.RcController{}, "get:ListRcs")
	beego.Router("/replicationcontroller/:name", &controllers.RcController{}, "get:GetRc")
	beego.Router("/replicationcontroller/delete/:name", &controllers.RcController{}, "delete:DeleteRc")
	beego.Router("/replicationcontroller/create", &controllers.RcController{}, "post:CreateRc")

	//Image Image
	beego.Router("/image/list", &controllers.ImageController{}, "get:ListImages")
	beego.Router("/image/delete", &controllers.ImageController{}, "delete:DelImage")
	//PVC PVC
	beego.Router("/pvc/list", &controllers.PVCController{}, "get:ListPVC")
}