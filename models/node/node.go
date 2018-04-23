package node

import (
	"myproject/models/myk8sclient"
	// "strconv"
	// "time"

	//"github.com/golang/glog"
	//resource "k8s.io/apimachinery/pkg/api/resource"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

//Node defination
type Node struct {
	IP            string            `json:"ip"`
	Name          string            `json:"name"`
	Lables        map[string]string `json:"lables"`
	Age           string            `json:"age"`
	Unschedulable bool              `json:"unschedulable"`
	SystemInfo    SystemInfo
	NodeCapacity  NodeCapacity
	Conditions    []Conditions
}

//SystemInfo defination
type SystemInfo struct {
	MachineID               string `json:"machineID"`
	SystemUUID              string `json:"systemUUID"`
	BootID                  string `json:"bootID"`
	KernelVersion           string `json:"kernelVersion"`
	OSImage                 string `json:"osImage"`
	ContainerRuntimeVersion string `json:"containerRuntimeVersion"`
	KubeletVersion          string `json:"kubeletVersion"`
	KubeProxyVersion        string `json:"kubeletProxyVersion"`
	OperatingSystem         string `json:"operationSystem"`
	Architecture            string `json:"architecture"`
}

//NodeCapacity NodeCapacity
type NodeCapacity struct {
	CPU    *resource.Quantity `json:"cpu"`
	Memory *resource.Quantity `json:"memory"`
	Pods   *resource.Quantity `json:"pods"`
}

//Conditions Conditions
type Conditions struct {
	Type               string `json:"type"`
	Status             string `json:"status"`
	LastHeartbeatTime  string `json:"lastHeartbeatTime"`
	LastTransitionTime string `json:"lastTransitionTime"`
	Reason             string `json:"reason"`
	Message            string `json:"message"`
}

//ListNode ListNode
func ListNode() *corev1.NodeList {
	client := myk8sclient.MyK8sClient()
	nodes, err := client.CoreV1().Nodes().List(metav1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}
	return nodes

}

//GetNode GetNode
func GetNode(nodeName string) *corev1.Node {
	client := myk8sclient.MyK8sClient()
	node, err := client.CoreV1().Nodes().Get(nodeName, metav1.GetOptions{})
	if err != nil {
		panic(err.Error())
	}
	return node
}
