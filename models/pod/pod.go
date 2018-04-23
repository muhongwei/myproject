package pod

import (
	"myproject/models/myk8sclient"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

//ListPods ListPods
func ListPods() *corev1.PodList {
	client := myk8sclient.MyK8sClient()
	pods, err := client.CoreV1().Pods("").List(metav1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}
	return pods

}

//GetPod GetPod
func GetPod(podname string, namespace string) *corev1.Pod {
	client := myk8sclient.MyK8sClient()
	pod, err := client.CoreV1().Pods(namespace).Get(podname, metav1.GetOptions{})
	if err != nil {
		panic(err.Error())
	}
	return pod
}

//DeletePod DeletePod
func DeletePod(podname string, namespace string) error {
	client := myk8sclient.MyK8sClient()
	err := client.CoreV1().Pods(namespace).Delete(podname, &metav1.DeleteOptions{})
	return err

}
