package handler

import (
	"context"
	"fmt"

	"github.com/pkg/errors"
	corev1 "k8s.io/api/core/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/klog/v2"
	clusterv1 "sigs.k8s.io/cluster-api/api/v1beta1"
	runtimecatalog "sigs.k8s.io/cluster-api/exp/runtime/catalog"
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

}

func (e *ExtensionHandler) DoAfterControlPlaneInitialized(ctx context.Context, request *runtimehooksv1.AfterControlPlaneInitializedRequest, response *runtimehooksv1.AfterControlPlaneInitializedResponse) {
	log := ctrl.LoggerFrom(ctx)
	log.Info("DoAfterControlPlaneInitialized is called")
	log.Info("Namespace:", request.Cluster.GetNamespace(), "ClusterName: ", request.Cluster.GetName())

	// Your implementation
	ok, err := e.checkConfigMap(ctx, &request.Cluster)
	if !ok {
		log.Info("not ok")
		if err != nil {
			log.Info("with error")
			response.Status = runtimehooksv1.ResponseStatusFailure
			response.Message = err.Error()
			return
		}
		log.Info("without error")
		if err := e.createConfigMap(ctx, &request.Cluster, runtimehooksv1.AfterControlPlaneInitialized, request.GetSettings(), response); err != nil {
			response.Status = runtimehooksv1.ResponseStatusFailure
			response.Message = err.Error()
			return
		}
	}
	log.Info("everything is ok")
}

func (e *ExtensionHandler) checkConfigMap(ctx context.Context, cluster *clusterv1.Cluster) (bool, error) {
	log := ctrl.LoggerFrom(ctx)
	log.Info("Checking for ConfigMap")

	configMap := &corev1.ConfigMap{}
	configMapName := fmt.Sprintf("%s-test-extension-hookresponse", cluster.Name)
	nsName := client.ObjectKey{Namespace: cluster.Namespace, Name: cluster.Name}
	if err := e.client.Get(ctx, nsName, configMap); err != nil {
		if apierrors.IsNotFound(err) {
			log.Info("ConfigMap not found")
			return false, nil
		}
		log.Error(err, "ConfigMap not found with an error")
		return false, errors.Wrapf(err, "failed to read the ConfigMap %s", klog.KRef(cluster.Namespace, configMapName))
	}
	log.Info("ConfigMap found")
	return true, nil
}

func (e *ExtensionHandler) createConfigMap(ctx context.Context, cluster *clusterv1.Cluster, hook runtimecatalog.Hook, settings map[string]string, response runtimehooksv1.ResponseObject) error {
	log := ctrl.LoggerFrom(ctx)
	log.Info("Creating ConfigMap")

	configMapName := fmt.Sprintf("%s-test-extension-hookresponse", cluster.Name)
	configMap := e.getConfigMap(cluster)
	if err := e.client.Create(ctx, configMap); err != nil {
		log.Error(err, "creating config map")
		return errors.Wrapf(err, "failed to create the ConfigMap %s", klog.KRef(cluster.Namespace, configMapName))
	}
	log.Info("configmap created successfully")
	return nil
}

func (e *ExtensionHandler) getConfigMap(cluster *clusterv1.Cluster) *corev1.ConfigMap {
	return &corev1.ConfigMap{
		ObjectMeta: metav1.ObjectMeta{
			Name:      fmt.Sprintf("%s-test-extension-hookresponses", cluster.Name),
			Namespace: cluster.Namespace,
		},
		Data: map[string]string{
			"AfterControlPlaneInitialized-preloadedResponse": `{"Status": "Success"}`,
		},
	}
}
