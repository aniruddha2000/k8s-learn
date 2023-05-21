.PHONY: manifest
manifest:
	kind create cluster --config manifests/capi/kind-cluster-with-extramounts.yaml
	clusterctl init --infrastructure docker
	kubectl create ns runtimesdk || true

.PHONY: manifest-down
manifest-down:
	kubectl delete -f manifests/config
	kind delete clusters --all
