package operator

import (
	acrt "github.com/appscode/go/runtime"
	"github.com/appscode/kubed/pkg/util"
	"github.com/appscode/log"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/apimachinery/pkg/watch"
	apiv1 "k8s.io/client-go/pkg/api/v1"
	"k8s.io/client-go/tools/cache"
)

// Blocks caller. Intended to be called as a Go routine.
func (op *Operator) WatchPersistentVolumeClaims() {
	if !util.IsPreferredAPIResource(op.KubeClient, apiv1.SchemeGroupVersion.String(), "PersistentVolumeClaim") {
		log.Warningf("Skipping watching non-preferred GroupVersion:%s Kind:%s", apiv1.SchemeGroupVersion.String(), "PersistentVolumeClaim")
		return
	}

	defer acrt.HandleCrash()

	lw := &cache.ListWatch{
		ListFunc: func(opts metav1.ListOptions) (runtime.Object, error) {
			return op.KubeClient.CoreV1().PersistentVolumeClaims(apiv1.NamespaceAll).List(metav1.ListOptions{})
		},
		WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
			return op.KubeClient.CoreV1().PersistentVolumeClaims(apiv1.NamespaceAll).Watch(metav1.ListOptions{})
		},
	}
	_, ctrl := cache.NewInformer(lw,
		&apiv1.PersistentVolumeClaim{},
		op.syncPeriod,
		cache.ResourceEventHandlerFuncs{
			DeleteFunc: func(obj interface{}) {
				if pvc, ok := obj.(*apiv1.PersistentVolumeClaim); ok {
					log.Infof("PersistentVolumeClaim %s@%s deleted", pvc.Name, pvc.Namespace)
					op.TrashCan.Save(pvc.ObjectMeta, obj)
				}
			},
		},
	)
	ctrl.Run(wait.NeverStop)
}
