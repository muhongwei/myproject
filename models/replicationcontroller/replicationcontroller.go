package replicationcontroller

import (
	"errors"
	"myproject/models/myk8sclient"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// RcMessage RcMessage
type RcMessage struct {
	Name       string      `json:"name"`
	Replicas   int32       `json:"replicas"`
	Containers []Container `json:"containers"`
}

//Container Container
type Container struct {
	ContainerName  string   `json:"containername"`
	ContainerImage string   `json:"containerimage"`
	Envs           []Env    `json:"env,omitempty"`
	Ports          []int32  `json:"ports"`
	Commands       []string `json:"command,omitempty"`
}

//Env Env
type Env struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

//ListReplicationControllers ListReplicationControllers
func ListReplicationControllers() *corev1.ReplicationControllerList {
	client := myk8sclient.MyK8sClient()
	rcs, err := client.CoreV1().ReplicationControllers("").List(metav1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}
	return rcs

}

// GetReplicationController GetReplicationController
func GetReplicationController(name string) *corev1.ReplicationController {
	client := myk8sclient.MyK8sClient()
	rc, err := client.CoreV1().ReplicationControllers("default").Get(name, metav1.GetOptions{})
	if err != nil {
		panic(err.Error())
	}
	return rc
}

// DeleteReplicationController DeleteReplicationController
func DeleteReplicationController(name string, namespace string) error {
	client := myk8sclient.MyK8sClient()
	rc := GetReplicationController(name)
	var num int32
	num = 0
	rc.Spec.Replicas = &num
	_, err2 := client.CoreV1().ReplicationControllers(namespace).Update(rc)
	if err2 != nil {
		panic("删除失败")
	}
	err := client.CoreV1().ReplicationControllers(namespace).Delete(name, &metav1.DeleteOptions{})
	return err
}

//CreateReplicationControllers CreateReplicationControllers
func CreateReplicationControllers(req *RcMessage) error {
	//判断请求数据内容
	if req.Name == "" {
		return errors.New("name can't be null")
	}
	if req.Replicas == 0 {
		var a int32
		a = 1
		req.Replicas = a
	}
	if req.Containers == nil {
		return errors.New("container can't be null")
	}
	for i := 0; i < len(req.Containers); i++ {
		if req.Containers[i].ContainerName == "" {
			return errors.New("containername can't be null")
		}
		if req.Containers[i].ContainerImage == "" {
			return errors.New("containerimage can't be null")
		}
	}
	client := myk8sclient.MyK8sClient()
	//副本控制器
	rc := &corev1.ReplicationController{
		ObjectMeta: metav1.ObjectMeta{
			Name: req.Name,
			Labels: map[string]string{
				"name": req.Name,
			},
		},
	}
	//副本控制器容器信息
	rc.Spec = corev1.ReplicationControllerSpec{
		Replicas: &req.Replicas,
		Selector: map[string]string{
			"name": req.Name,
		},
		Template: &corev1.PodTemplateSpec{
			ObjectMeta: metav1.ObjectMeta{
				Labels: map[string]string{
					"name": req.Name,
				},
			},
		},
	}
	//副本控制器容器模板
	var containers []corev1.Container
	for _, container := range req.Containers {
		var ports []corev1.ContainerPort
		var envs []corev1.EnvVar
		for _, port := range container.Ports {
			p := corev1.ContainerPort{
				ContainerPort: port,
			}
			for _, e := range container.Envs {
				env := corev1.EnvVar{
					Name:  e.Name,
					Value: e.Value,
				}
				envs = append(envs, env)
			}
			ports = append(ports, p)
		}
		//副本控制器容器模板
		c := corev1.Container{
			Name:    container.ContainerName,
			Image:   container.ContainerImage,
			Env:     envs,
			Ports:   ports,
			Command: container.Commands,
		}
		containers = append(containers, c)
	}
	rc.Spec.Template.Spec.Containers = containers

	// exists, geterr := client.CoreV1().ReplicationControllers("default").Get(req.Name, metav1.GetOptions{})
	// if exists != nil {
	// 	return errors.New("The name have existed in the cluster")
	// }
	_, err := client.CoreV1().ReplicationControllers("default").Create(rc)
	return err
}
