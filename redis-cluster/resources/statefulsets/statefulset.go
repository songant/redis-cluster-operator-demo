package statefulsets

import (
	v1 "k8s.io/api/apps/v1"
	v14 "k8s.io/api/core/v1"
	v13 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	v12 "redis-cluster/api/v1"
	"redis-cluster/constant"
)

func NewStatefulsetForCr(cluster *v12.RedisCluster, ssName, svcName string, labels map[string]string, statefulsetReplicas int32) (v1.StatefulSet, error) {

	namespace := cluster.Namespace
	ss := v1.StatefulSet{
		ObjectMeta: v13.ObjectMeta{
			Name:      ssName,
			Namespace: namespace,
			Labels:    labels,
			OwnerReferences: []v13.OwnerReference{
				*v13.NewControllerRef(cluster, schema.GroupVersionKind{
					Group:   constant.SchemeGroupVersion.Group,
					Version: constant.SchemeGroupVersion.Version,
					Kind:    "RedisCluster",
				}),
			},
		},

		Spec: v1.StatefulSetSpec{
			Replicas:    &statefulsetReplicas,
			ServiceName: svcName,
			Selector: &v13.LabelSelector{
				MatchLabels: labels,
			},
			Template: v14.PodTemplateSpec{
				ObjectMeta: v13.ObjectMeta{
					Labels: labels,
				},
				Spec: v14.PodSpec{},
			},
		},
	}
	return ss, nil
}
