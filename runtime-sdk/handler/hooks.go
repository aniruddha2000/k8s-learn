package handler

import (
	"context"
	"time"

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

	// Your implementation
	log.Info("It's my log", request.GetObjectKind().GroupVersionKind().Kind)
	log.Info(request.Cluster.GetName(), request.Cluster.GetNamespace())

	log.Info("start waiting")

	time.Sleep(100 * time.Millisecond)

	log.Info("finish waiting")
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
