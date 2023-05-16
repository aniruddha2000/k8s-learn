.PHONY: manifest
manifest:
	kubectl create ns runtimesdk || true
	kubectl apply -f manifests/config

.PHONY: manifest-down
manifest-down:
	kubectl delete -f manifests/config
