package k8sutil

import (
	"context"
	"fmt"
	"net"
	"os"
	"sync"
	"time"

	rainbondv1alpha1 "github.com/goodrain/rainbond-operator/pkg/apis/rainbond/v1alpha1"
	rainbondversiond "github.com/goodrain/rainbond-operator/pkg/generated/clientset/versioned"
	"github.com/goodrain/rainbond-operator/pkg/util/commonutil"
	"github.com/goodrain/rainbond-operator/pkg/util/constants"
	"github.com/sirupsen/logrus"
	corev1 "k8s.io/api/core/v1"
	k8sErrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/tools/reference"
	"k8s.io/kubectl/pkg/describe"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var once sync.Once
var clientset kubernetes.Interface

func GetClientSet() kubernetes.Interface {
	if clientset == nil {
		once.Do(func() {
			config := MustNewKubeConfig("")
			clientset = kubernetes.NewForConfigOrDie(config)
		})
	}
	return clientset
}

func MustNewKubeConfig(kubeconfigPath string) *rest.Config {
	if kubeconfigPath != "" {
		cfg, err := clientcmd.BuildConfigFromFlags("", kubeconfigPath)
		if err != nil {
			panic(err)
		}
		return cfg
	}

	cfg, err := InClusterConfig()
	if err != nil {
		panic(err)
	}
	return cfg
}

func NewKubeConfig() (*rest.Config, error) {
	cfg, err := InClusterConfig()
	if err != nil {
		return nil, err
	}
	return cfg, nil
}

func InClusterConfig() (*rest.Config, error) {
	// Work around https://github.com/kubernetes/kubernetes/issues/40973
	// See https://github.com/coreos/etcd-operator/issues/731#issuecomment-283804819
	if len(os.Getenv("KUBERNETES_SERVICE_HOST")) == 0 {
		addrs, err := net.LookupHost("kubernetes.default.svc")
		if err != nil {
			panic(err)
		}
		os.Setenv("KUBERNETES_SERVICE_HOST", addrs[0])
	}
	if len(os.Getenv("KUBERNETES_SERVICE_PORT")) == 0 {
		os.Setenv("KUBERNETES_SERVICE_PORT", "443")
	}
	cfg, err := rest.InClusterConfig()
	if err != nil {
		return nil, err
	}
	return cfg, nil
}

// HostPath returns a pointer to the HostPathType value passed in.
func HostPath(hostpath corev1.HostPathType) *corev1.HostPathType {
	return &hostpath
}

// MountPropagationMode returns a pointer to the MountPropagationMode value passed in.
func MountPropagationMode(moundPropagationMode corev1.MountPropagationMode) *corev1.MountPropagationMode {
	return &moundPropagationMode
}

// PersistentVolumeReclaimPolicy returns a pointer to the PersistentVolumeReclaimPolicy value passed in.
func PersistentVolumeReclaimPolicy(persistentVolumeReclaimPolicy corev1.PersistentVolumeReclaimPolicy) *corev1.PersistentVolumeReclaimPolicy {
	return &persistentVolumeReclaimPolicy
}

func UpdateCRStatus(client client.Client, obj runtime.Object) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	err := client.Status().Update(ctx, obj)
	if err != nil {
		err = client.Status().Update(ctx, obj)
		if err != nil {
			return fmt.Errorf("update custom resource status: %v", err)
		}
	}
	return nil
}

func MaterRoleLabel(key string) map[string]string {
	var labels map[string]string
	switch key {
	case describe.LabelNodeRolePrefix + "master":
		labels = map[string]string{
			key: "",
		}
	case describe.NodeLabelRole:
		labels = map[string]string{
			key: "master",
		}
	}
	return labels
}

func PersistentVolumeClaimForGrdata(ns, claimName string, accessModes []corev1.PersistentVolumeAccessMode, labels map[string]string, storageClassName string, storageRequest int64) *corev1.PersistentVolumeClaim {
	size := resource.NewQuantity(storageRequest*1024*1024*1024, resource.BinarySI)
	pvc := &corev1.PersistentVolumeClaim{
		ObjectMeta: metav1.ObjectMeta{
			Name:      claimName,
			Namespace: ns,
			Labels:    labels,
		},
		Spec: corev1.PersistentVolumeClaimSpec{
			AccessModes: accessModes,
			Resources: corev1.ResourceRequirements{
				Requests: map[corev1.ResourceName]resource.Quantity{
					corev1.ResourceStorage: *size,
				},
			},
			StorageClassName: commonutil.String(storageClassName),
		},
	}

	return pvc
}

func GetFoobarPVC(ctx context.Context, client client.Client, ns string) (*corev1.PersistentVolumeClaim, error) {
	pvc := corev1.PersistentVolumeClaim{}
	err := client.Get(ctx, types.NamespacedName{Namespace: ns, Name: constants.FoobarPVC}, &pvc)
	if err != nil {
		return nil, err
	}
	return &pvc, nil
}

// EventsForPersistentVolumeClaim -
func EventsForPersistentVolumeClaim(pvc *corev1.PersistentVolumeClaim) (*corev1.EventList, error) {
	clientset := GetClientSet()
	ref, err := reference.GetReference(scheme.Scheme, pvc)
	if err != nil {
		return nil, err
	}
	ref.Kind = ""
	if _, isMirrorPod := pvc.Annotations[corev1.MirrorPodAnnotationKey]; isMirrorPod {
		ref.UID = types.UID(pvc.Annotations[corev1.MirrorPodAnnotationKey])
	}
	events, err := clientset.CoreV1().Events(pvc.GetNamespace()).Search(scheme.Scheme, ref)
	return events, err
}

// IsPodReady checks if the given pod is ready or not.
func IsPodReady(pod *corev1.Pod) bool {
	for _, condition := range pod.Status.Conditions {
		if condition.Type == corev1.PodReady && condition.Status == corev1.ConditionTrue {
			return true
		}
	}
	return false
}

// IsPodCompleted checks if the given pod is ready or not.
func IsPodCompleted(pod *corev1.Pod) bool {
	for _, condition := range pod.Status.Conditions {
		if condition.Type == corev1.PodReady && condition.Status == corev1.ConditionFalse && condition.Reason == "PodCompleted" {
			return true
		}
	}
	return false
}

// CreateIfNotExists -
func CreateIfNotExists(ctx context.Context, client client.Client, obj runtime.Object, meta metav1.Object) error {
	err := client.Get(ctx, types.NamespacedName{Name: meta.GetName(), Namespace: meta.GetNamespace()}, obj)
	if err != nil {
		if !k8sErrors.IsNotFound(err) {
			return err
		}
		return client.Create(ctx, obj)
	}
	return nil
}

// CreateOrUpdateRbdComponent -
func CreateOrUpdateRbdComponent(clientset rainbondversiond.Interface, cpt *rainbondv1alpha1.RbdComponent) error {
	old, err := clientset.RainbondV1alpha1().RbdComponents(cpt.Namespace).Get(cpt.GetName(), metav1.GetOptions{})
	if err != nil {
		if k8sErrors.IsNotFound(err) {
			logrus.Infof("[CreateOrUpdateRbdComponent] create rbdcomponent %s/%s", cpt.Namespace, cpt.Name)
			_, err := clientset.RainbondV1alpha1().RbdComponents(cpt.Namespace).Create(cpt)
			if err != nil {
				return err
			}
			return nil
		}
		return err
	}

	logrus.Infof("[CreateOrUpdateRbdComponent] update rbdcomponent %s/%s", cpt.Namespace, cpt.Name)
	cpt.ResourceVersion = old.ResourceVersion
	cpt.Spec.Replicas = old.Spec.Replicas
	_, err = clientset.RainbondV1alpha1().RbdComponents(cpt.Namespace).Update(cpt)
	if err != nil {
		return err
	}
	return nil
}

// ListNodes returns all nodes.
func ListNodes(ctx context.Context, c client.Client) ([]corev1.Node, error) {
	nodeList := &corev1.NodeList{}
	if err := c.List(ctx, nodeList); err != nil {
		return nil, err
	}
	return nodeList.Items, nil
}
