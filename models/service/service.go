package service

import (
	"myproject/models/myk8sclient"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
)

//ServiceInfo ServiceInfo
type ServiceInfo struct {
	Name     string             `json:"name,omitempty"`
	Label    map[string]string  `json:"labels,omitempty"`
	Type     corev1.ServiceType `json:"type,omitempty"`
	Ports    []Port             `json:"ports,omitempty"`
	Selector map[string]string  `json:"selector,omitempty"`
}

//Port Port
type Port struct {
	Port       int32 `json:"port"`       //服务的端口
	TargetPort int   `json:"targetport"` //容器应用监听的端口
	NodePort   int32 `json:"nodeport"`   //物理机端口对外提供服务
}

//ListServices ListServices
func ListServices() *corev1.ServiceList {
	client := myk8sclient.MyK8sClient()
	services, err := client.CoreV1().Services("").List(metav1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}
	return services

}

//GetService GetService
func GetService(name string) *corev1.Service {
	client := myk8sclient.MyK8sClient()
	service, err := client.CoreV1().Services("default").Get(name, metav1.GetOptions{})
	if err != nil {
		panic(err.Error())
	}
	return service
}

//DeleteService DeleteService
func DeleteService(name string) error {
	client := myk8sclient.MyK8sClient()
	err := client.CoreV1().Services("default").Delete(name, &metav1.DeleteOptions{})
	return err
}

//CreateService CreateService
func CreateService(req *ServiceInfo) error {
	client := myk8sclient.MyK8sClient()
	service := &corev1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name: req.Name,
			Labels: map[string]string{
				"name": req.Name,
			},
		},
		Spec: corev1.ServiceSpec{
			Selector: req.Selector,
			Type:     req.Type,
		},
	}
	var ports []corev1.ServicePort
	for _, port := range req.Ports {
		p := corev1.ServicePort{
			Port:       port.Port,
			TargetPort: intstr.FromInt(port.TargetPort),
			NodePort:   port.NodePort,
		}
		ports = append(ports, p)
	}
	service.Spec.Ports = ports

	_, err := client.CoreV1().Services("default").Create(service)
	return err
}
