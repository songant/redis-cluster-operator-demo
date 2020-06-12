package manager

import (
	"fmt"
	"k8s.io/apimachinery/pkg/api/errors"
	v1 "redis-cluster/api/v1"
	"redis-cluster/k8sutil"
	"redis-cluster/resources/statefulsets"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"strconv"
)

type IEnsureResource interface {
	EnsureRedisStatefulsets(cluster *v1.RedisCluster, labels map[string]string) (bool, error)
	EnsureRedisSvc(cluster *v1.RedisCluster) (bool, error)
}

type realEnsureResource struct {
	statefulsetClient k8sutil.IStatefulsetControl
	svcClient         k8sutil.IServiceControl
}

func NewRealEnsureResource(client client.Client) IEnsureResource {
	return &realEnsureResource{
		statefulsetClient: k8sutil.NewStatefulsetController(client),
		svcClient:         k8sutil.NewServiceController(client),
	}
}

/**

 */
func (r *realEnsureResource) EnsureRedisStatefulsets(cluster *v1.RedisCluster, labels map[string]string) (bool, error) {
	masterSize, _ := strconv.Atoi(cluster.Spec.MasterSize) //主节点的数量即statefulset的数量
	for i := 0; i < masterSize; i++ {
		name := fmt.Sprintf("redis-cluster-%s-%d", cluster.Spec.Name, i)
		svcName := fmt.Sprintf("%s-%d", cluster.Spec.ServiceName, i)
		labels["statefulset"] = cluster.Name
		if stsupdate, e := r.ensureRedisStatefulset(cluster, name, svcName, labels); e != nil {
			return false, e
		} else if stsupdate {
			return stsupdate, nil
		}
	}
	return false, nil
}

func (r *realEnsureResource) ensureRedisStatefulset(cluster *v1.RedisCluster, ssName, svcName string, labels map[string]string) (bool, error) {
	_, err := r.statefulsetClient.GetStatefulset(cluster.Namespace, ssName)
	if err == nil { //证明redis statefulset 存在
		//通过相关属性判断statefulset是否需要更新
		//TODO:1)判断从节点的数量；2）判断镜像是否更改；3）判断创建pod时的资源限制是否更改；4）判断redis的password是否更改

	} else if err != nil && errors.IsNotFound(err) {
		//创建cr for statefulset
		statefulsetReplicas := cluster.Spec.Replicas + 1 //每一个statefulset 创建pod的数量
		statefulsetCr, err := statefulsets.NewStatefulsetForCr(cluster, ssName, svcName, labels, statefulsetReplicas)
		if err != nil {
			return false, err
		}
		return false, r.statefulsetClient.CreateStatefulset(&statefulsetCr)
	}
	//sts = sts
	return false, nil

}

/**

 */
func (r *realEnsureResource) EnsureRedisSvc(cluster *v1.RedisCluster) (bool, error) {
	return true, nil
}
