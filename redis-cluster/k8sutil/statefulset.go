package k8sutil

import (
	"context"
	v1 "k8s.io/api/apps/v1"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type IStatefulsetControl interface {
	CreateStatefulset(*v1.StatefulSet) error
	UpdateStatefulset(set *v1.StatefulSet) error
	GetStatefulset(namespace, name string) (*v1.StatefulSet, error)
	ListStatefulsetByLables(namespace string, labels map[string]string) (*v1.StatefulSet, error)
	DeleteStatefulsetByName(namespace, name string) error
	DeleteStatefuslet(set *v1.StatefulSet) error
}

type statefulsetController struct {
	client client.Client
}

func NewStatefulsetController(client client.Client) IStatefulsetControl {
	return &statefulsetController{client: client}
}

func (s *statefulsetController) CreateStatefulset(ss *v1.StatefulSet) error {
	return s.client.Create(context.TODO(), ss)
}

/**
更新statefulset
*/
func (s *statefulsetController) UpdateStatefulset(set *v1.StatefulSet) error {
	return s.client.Update(context.TODO(), set)
}

/**
get statefulset by namespace and statefulset name
*/
func (s *statefulsetController) GetStatefulset(namespace, name string) (*v1.StatefulSet, error) {
	ss := &v1.StatefulSet{}
	err := s.client.Get(context.TODO(), types.NamespacedName{
		Namespace: namespace,
		Name:      name,
	}, ss)

	return ss, err
}

/**
list statefulset by labels
*/
func (s *statefulsetController) ListStatefulsetByLables(namespace string, labels map[string]string) (*v1.StatefulSet, error) {
	sts := &v1.StatefulSet{}
	err := s.client.List(context.TODO(), sts, client.InNamespace(namespace), client.MatchingLabels(labels))
	return sts, err
}

/**
delete statefulset
*/
func (s *statefulsetController) DeleteStatefuslet(set *v1.StatefulSet) error {
	return s.client.Delete(context.TODO(), set)
}

/**
delete sts by statefulset name
*/
func (s *statefulsetController) DeleteStatefulsetByName(namespace, name string) error {
	set, e := s.GetStatefulset(namespace, name)
	if e != nil {
		return e
	}
	return s.client.Delete(context.TODO(), set)
}
