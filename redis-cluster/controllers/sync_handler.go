package controllers

import (
	v1 "redis-cluster/api/v1"
	"redis-cluster/controllers/manager"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type syncContext struct {
	cluster *v1.RedisCluster
	ensure  manager.IEnsureResource
	client  client.Client
}

func ensureRedisCluster(ctx *syncContext) error {

	ctx.ensure = manager.NewRealEnsureResource(ctx.client)
	labels := map[string]string{}
	ctx.ensure.EnsureRedisStatefulsets(ctx.cluster, labels)
	return nil
}
