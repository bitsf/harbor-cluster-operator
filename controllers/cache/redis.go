package cache

import (
	goharborv1 "github.com/goharbor/harbor-cluster-operator/api/v1"
	"github.com/goharbor/harbor-cluster-operator/lcm"
)

// RedisReconciler implement the Reconciler interface and lcm.Controller interface.
type RedisReconciler struct {
	HarborCluster *goharborv1.HarborCluster
	Client        k8s.Client
	Recorder      record.EventRecorder
	Log           logr.Logger
	DClient       k8s.DClient
	Scheme        *runtime.Scheme
	ExpectCR      *unstructured.Unstructured
	ActualCR      *unstructured.Unstructured
	Labels        map[string]string
	Name          string
	Namespace     string
	CRStatus      *lcm.CRStatus
}

// Reconciler implements the reconcile logic of redis service
func (redis *RedisReconciler) Reconcile() (*lcm.CRStatus, error) {
	return nil, nil
}

func (redis *RedisReconciler) Provision(spec *goharborv1.HarborCluster) (*lcm.CRStatus, error) {
	panic("implement me")
}

func (redis *RedisReconciler) Delete() (*lcm.CRStatus, error) {
	panic("implement me")
}

func (redis *RedisReconciler) ScaleUp(newReplicas uint64) (*lcm.CRStatus, error) {
	panic("implement me")
}

func (redis *RedisReconciler) ScaleDown(newReplicas uint64) (*lcm.CRStatus, error) {
	panic("implement me")
}

func (redis *RedisReconciler) Update(spec *goharborv1.HarborCluster) (*lcm.CRStatus, error) {
	panic("implement me")
}
