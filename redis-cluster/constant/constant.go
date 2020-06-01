package constant

import "k8s.io/apimachinery/pkg/runtime/schema"

var (
	SchemeGroupVersion = schema.GroupVersion{
		Group:   "redis.song.com",
		Version: "v1",
	}
)
