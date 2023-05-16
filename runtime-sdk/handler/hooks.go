package handler

import (
	"context"
	"fmt"
	"time"

	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/cluster-api/api/v1beta1"
	runtimehooksv1 "sigs.k8s.io/cluster-api/exp/runtime/hooks/api/v1alpha1"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type ExtensionHandler struct {
	client client.Client
}

func NewExtensionHandlers(client client.Client) *ExtensionHandler {
	return &ExtensionHandler{
		client: client,
	}
}

func (e *ExtensionHandler) DoBeforeClusterDelete(ctx context.Context, request *runtimehooksv1.BeforeClusterDeleteRequest, response *runtimehooksv1.BeforeClusterDeleteResponse) {
	log := ctrl.LoggerFrom(ctx)
	log.Info("DoBeforeClusterDelete is called")
	log.Info("Kind: ", request.GetObjectKind().GroupVersionKind().Kind, "ClusterName: ", request.Cluster.GetName())

	// Your implementation
	obj := request.DeepCopy()
	if obj.Cluster.Status.Phase != "Running" || !obj.Cluster.Status.ControlPlaneReady || !obj.Cluster.Status.InfrastructureReady {
		response.CommonResponse = runtimehooksv1.CommonResponse{
			Status:  runtimehooksv1.ResponseStatusFailure,
			Message: fmt.Sprintf("Cluster is: %s, control plane ready: %v, infrastructure ready: %v", obj.Cluster.Status.Phase, obj.Cluster.Status.ControlPlaneReady, obj.Cluster.Status.InfrastructureReady),
		}
	}

	clusterObj := v1beta1.Cluster{}
	nsName := types.NamespacedName{
		Namespace: "default",
		Name:      "capi-quickstart",
	}
	err := e.client.Get(ctx, nsName, &clusterObj, &client.GetOptions{})
	if err != nil {
		response.Status = runtimehooksv1.ResponseStatusFailure
		response.Message = err.Error()
		return
	}
	log.Info("Cluster Obj get - successfull")

	response.CommonResponse = runtimehooksv1.CommonResponse{
		Status:  runtimehooksv1.ResponseStatusFailure,
		Message: fmt.Sprintf("Cluster is: %s, control plane ready: %v, infrastructure ready: %v", obj.Cluster.Status.Phase, obj.Cluster.Status.ControlPlaneReady, obj.Cluster.Status.InfrastructureReady),
	}
}

func (e *ExtensionHandler) DoBeforeClusterCreate(ctx context.Context, request *runtimehooksv1.BeforeClusterCreateRequest, response *runtimehooksv1.BeforeClusterCreateResponse) {
	log := ctrl.LoggerFrom(ctx)
	log.Info("BeforeClusterCreate is called")

	// Your implementation
	log.Info("It's my log", request.GetObjectKind().GroupVersionKind().Kind)
	log.Info(request.Cluster.GetName(), request.Cluster.GetNamespace())

	log.Info("start waiting")

	time.Sleep(100 * time.Millisecond)

	log.Info("finish waiting")
}
