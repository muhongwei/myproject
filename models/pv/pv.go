package pv

import (
	"fmt"
	"myproject/models/myk8sclient"

	"github.com/golang/glog"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type PVMessage struct {
	Name        string `json:"name"`
	AccessModes string `json:"accessmodes"`
	Storage     string `json:"storage"`
	Path        string `json:"path"`
	Server      string `json:"server"`
}

//ListPV ListPV
func ListPV() *corev1.PersistentVolumeList {
	client := myk8sclient.MyK8sClient()
	pvs, err := client.CoreV1().PersistentVolumes().List(metav1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}
	for _, pv := range pvs.Items {
		fmt.Println(pv.ObjectMeta.Name)
	}
	return pvs

}

//CreatePV CreatePV
func CreatePV(reqpv *PVMessage) error {
	client := myk8sclient.MyK8sClient()
	quantity, err := resource.ParseQuantity(reqpv.Storage)
	if err != nil {
		return err
	}

	pv := &corev1.PersistentVolume{
		ObjectMeta: metav1.ObjectMeta{
			Name: reqpv.Name,
		},
		Spec: corev1.PersistentVolumeSpec{
			AccessModes: []corev1.PersistentVolumeAccessMode{
				corev1.PersistentVolumeAccessMode(reqpv.AccessModes),
			},
			Capacity: corev1.ResourceList{
				"storage": quantity,
			},
			PersistentVolumeSource: corev1.PersistentVolumeSource{
				NFS: &corev1.NFSVolumeSource{
					Path:   reqpv.Path,
					Server: reqpv.Server,
				},
			},
		},
	}
	fmt.Println(pv)
	_, err2 := client.CoreV1().PersistentVolumes().Create(pv)
	if err2 != nil {
		glog.Errorln(err2)
	}
	return err2
}

//DeletePV DeletePV
func DeletePV(name string) error {
	client := myk8sclient.MyK8sClient()
	err := client.CoreV1().PersistentVolumes().Delete(name, &metav1.DeleteOptions{})
	return err
}

//GetPV GetPV
func GetPV(name string) *corev1.PersistentVolume {
	client := myk8sclient.MyK8sClient()
	pv, err := client.CoreV1().PersistentVolumes().Get(name, metav1.GetOptions{})
	if err != nil {
		panic(err.Error())
	}
	return pv
}
