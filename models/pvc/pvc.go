package pvc

import (
	"fmt"
	"myproject/models/myk8sclient"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

//ListPVC ListPVC
func ListPVC() *corev1.PersistentVolumeClaimList {
	client := myk8sclient.MyK8sClient()
	pvcs, err := client.CoreV1().PersistentVolumeClaims("").List(metav1.ListOptions{})
	// pvcs, err := client.CoreV1().PersistentVolumeClaims("").List(metav1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}
	// pv := &corev1.PersistentVolumeClaim{}
	// client.CoreV1().PersistentVolumeClaims("").Create(pv)
	//client.CoreV1().PersistentVolumeClaims("").Update()
	//&corev1.BetaStorageClassAnnotation
	for _, pvc := range pvcs.Items {
		fmt.Println(pvc.ObjectMeta.Name)
	}
	return pvcs

}
