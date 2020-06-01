package k8sutil

import (
	"context"
	v1 "k8s.io/api/core/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type IServiceControl interface {
	CreateService(svc *v1.Service) error
	UpdateService(svc *v1.Service) error
}

type serviceController struct {
	client client.Client
}

func NewServiceController(client client.Client) IServiceControl {
	return &serviceController{client: client}
}

func (s *serviceController) CreateService(svc *v1.Service) error {
	return s.client.Create(context.TODO(), svc)
}

func (s *serviceController) UpdateService(svc *v1.Service) error {
	return s.client.Update(context.TODO(), svc)
}
