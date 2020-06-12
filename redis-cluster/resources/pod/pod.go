package pod

import (
	"fmt"
	v12 "k8s.io/api/core/v1"
	v1 "redis-cluster/api/v1"
)

const redisPodName = "redis"

func NewPodForCr(cluster *v1.RedisCluster, password *v12.EnvVar) v12.Container {
	container := v12.Container{
		Name:    redisPodName,
		Image:   cluster.Spec.Image,
		Command: getRedisCommand(cluster, password),
		Ports: []v12.ContainerPort{
			{
				Name:          "client",
				ContainerPort: 6379,
				Protocol:      v12.ProtocolTCP,
			},
			{
				Name:          "gossip",
				ContainerPort: 16379,
				Protocol:      v12.ProtocolTCP,
			},
		},
		EnvFrom: nil,
		Env: []v12.EnvVar{
			{
				Name: "POD_IP",
				ValueFrom: &v12.EnvVarSource{
					FieldRef: &v12.ObjectFieldSelector{
						FieldPath: "status.podIP",
					},
				},
			},
		},
		Resources:                v12.ResourceRequirements{},
		VolumeMounts:             nil,
		VolumeDevices:            nil,
		LivenessProbe:            nil,
		ReadinessProbe:           nil,
		StartupProbe:             nil,
		Lifecycle:                nil,
		TerminationMessagePath:   "",
		TerminationMessagePolicy: "",
		ImagePullPolicy:          "",
		SecurityContext:          nil,
		Stdin:                    false,
		StdinOnce:                false,
		TTY:                      false,
	}

	//判度password是否为空
	if password != nil {
		container.Env = append(container.Env, *password)

	}

	return container
}

/**
获取redis启动时
*/
func getRedisCommand(cluster *v1.RedisCluster, password *v12.EnvVar) []string {

	cmd := []string{
		"redis-server",
		"/conf/redis.conf",
		"--cluster-enabled yes",
	}
	if password != nil {
		cmd = append(cmd, fmt.Sprintf("--requirepass '$(%s)'", "redis_password"),
			fmt.Sprintf("--masterauth '$(%s)'", "redis_password"))
	}
	return cmd
}
