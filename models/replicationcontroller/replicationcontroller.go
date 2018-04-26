package replicationcontroller

import (
	"errors"
	"myproject/models/myk8sclient"

	"github.com/golang/glog"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// RcMessage RcMessage
type RcMessage struct {
	Name       string       `json:"name"`
	Replicas   int32        `json:"replicas"`
	Containers []Container  `json:"containers"`
	Volumes    []VolumeInfo `json:"volume"`
}

//Container Container
type Container struct {
	ContainerName  string             `json:"containername"`
	ContainerImage string             `json:"containerimage"`
	Envs           []Env              `json:"env,omitempty"`
	Ports          []int32            `json:"ports"`
	Commands       []string           `json:"command,omitempty"`
	VolumeMounts   []VolumeMountsInfo `json:"volumemountinfo"`
}

//VolumeInfo VolumeInfo
type VolumeInfo struct {
	Name      string `json:"name"`
	ClaimName string `json:"claimname"`
}

//VolumeMountsInfo VolumeMountsInfo
type VolumeMountsInfo struct {
	Name      string `json:"name"`
	MountPath string `json:"mountpath"`
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
		var volumeMounts []corev1.VolumeMount

		for _, port := range container.Ports {
			p := corev1.ContainerPort{
				ContainerPort: port,
			}
			ports = append(ports, p)

		}

		for _, e := range container.Envs {
			env := corev1.EnvVar{
				Name:  e.Name,
				Value: e.Value,
			}
			envs = append(envs, env)
		}

		for _, v := range container.VolumeMounts {
			vm := corev1.VolumeMount{
				Name:      v.Name,
				MountPath: v.MountPath,
			}
			volumeMounts = append(volumeMounts, vm)
		}

		//副本控制器容器模板
		c := corev1.Container{
			Name:         container.ContainerName,
			Image:        container.ContainerImage,
			Env:          envs,
			Ports:        ports,
			Command:      container.Commands,
			VolumeMounts: volumeMounts,
		}
		containers = append(containers, c)
	}
	rc.Spec.Template.Spec.Containers = containers

	//volume绑定pvc信息
	var volumes []corev1.Volume
	for _, volume := range req.Volumes {
		v := corev1.Volume{
			Name: volume.Name,
			VolumeSource: corev1.VolumeSource{
				PersistentVolumeClaim: &corev1.PersistentVolumeClaimVolumeSource{
					ClaimName: volume.ClaimName,
				},
			},
		}
		volumes = append(volumes, v)
	}
	rc.Spec.Template.Spec.Volumes = volumes
	// exists, geterr := client.CoreV1().ReplicationControllers("default").Get(req.Name, metav1.GetOptions{})
	// if exists != nil {
	// 	return errors.New("The name have existed in the cluster")
	// }
	//fmt.Println(rc)
	_, err := client.CoreV1().ReplicationControllers("default").Create(rc)
	if err != nil {
		glog.Errorln(err)
	}
	return err
}
