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
	beego.Router("/user/login", &controllers.UserController{})
	beego.Router("/user/info", &controllers.UserController{}, "get:UserInfo")
	beego.Router("/user/register", &controllers.UserController{}, "post:RegisterUser")
	beego.Router("/user/delete/:id", &controllers.UserController{}, "delete:DeleteUser")
	beego.Router("/user/registerpage", &controllers.UserController{}, "get:UserRegisterPage")
	beego.Router("/user/getalluser", &controllers.UserController{}, "get:GetAllUser")

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

	beego.Router("/user/pv/list", &controllers.PVController{}, "get:UserListPV")
	beego.Router("/user/pv/create", &controllers.PVController{}, "get:UserCreatePV")
	beego.Router("/user/pv/get/:name", &controllers.PVController{}, "get:UserGetPV")

	beego.Router("/user/pvc/list", &controllers.PVCController{}, "get:UserListPVC")
	beego.Router("/user/pvc/create", &controllers.PVCController{}, "get:UserCreatePVC")
	beego.Router("/user/pvc/get/:name", &controllers.PVCController{}, "get:UserGetPVC")

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
	beego.Router("/pvc/get/:name", &controllers.PVCController{}, "get:GetPVC")
	beego.Router("/pvc/delete/:name", &controllers.PVCController{}, "delete:DeletePVC")
	beego.Router("/pvc/create", &controllers.PVCController{}, "post:CreatePVC")
	//PV PV
	beego.Router("/pv/list", &controllers.PVController{}, "get:ListPV")
	beego.Router("/pv/get/:name", &controllers.PVController{}, "get:GetPV")
	beego.Router("/pv/delete/:name", &controllers.PVController{}, "delete:DeletePV")
	beego.Router("/pv/create", &controllers.PVController{}, "post:CreatePV")
}
