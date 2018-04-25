package pvc

import (
	"fmt"
	"myproject/models/myk8sclient"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type PVCMessage struct {
	Name        string `json:"name"`
	AccessModes string `json:"accessmodes"`
	Storage     string `json:"storage"`
}

//ListPVC ListPVC
func ListPVC() *corev1.PersistentVolumeClaimList {
	client := myk8sclient.MyK8sClient()
	pvcs, err := client.CoreV1().PersistentVolumeClaims("").List(metav1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}
	for _, pvc := range pvcs.Items {
		fmt.Println(pvc.ObjectMeta.Name)
	}
	return pvcs

}

//CreatePVC CreatePVC
func CreatePVC(reqpvc *PVCMessage) error {
	client := myk8sclient.MyK8sClient()
	quantity, err := resource.ParseQuantity(reqpvc.Storage)
	if err != nil {
		return err
	}
	pvc := &corev1.PersistentVolumeClaim{
		ObjectMeta: metav1.ObjectMeta{
			Name: reqpvc.Name,
		},
		Spec: corev1.PersistentVolumeClaimSpec{
			AccessModes: []corev1.PersistentVolumeAccessMode{
				corev1.PersistentVolumeAccessMode(reqpvc.AccessModes),
			},
			Resources: corev1.ResourceRequirements{
				Requests: corev1.ResourceList{
					"storage": quantity,
				},
			},
		},
	}
	_, err2 := client.CoreV1().PersistentVolumeClaims("default").Create(pvc)

	return err2
}

//DeletePVC DeletePVC
func DeletePVC(name string) error {
	client := myk8sclient.MyK8sClient()
	err := client.CoreV1().PersistentVolumeClaims("default").Delete(name, &metav1.DeleteOptions{})
	return err
}

//GetPVC GetPVC
func GetPVC(name string) *corev1.PersistentVolumeClaim {
	client := myk8sclient.MyK8sClient()
	pvc, err := client.CoreV1().PersistentVolumeClaims("default").Get(name, metav1.GetOptions{})
	if err != nil {
		panic(err.Error())
	}
	return pvc
}
