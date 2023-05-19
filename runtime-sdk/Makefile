.PHONY: manifest
manifest:
	export CLUSTER_TOPOLOGY=true
	export EXP_RUNTIME_SDK=true
	export SERVICE_CIDR=["10.96.0.0/12"]
	export POD_CIDR=["192.168.0.0/16"]
	export SERVICE_DOMAIN="k8s.test"
	kind create cluster --config manifests/capi/kind-cluster-with-extramounts.yaml
	clusterctl init --infrastructure docker
	kubectl create ns runtimesdk || true

.PHONY: manifest-down
manifest-down:
	kubectl delete -f manifests/config
	kind delete clusters --all
