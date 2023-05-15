## Setup

### Set of instructions executed so far

```shell
$ kind create cluster --config capi/kind-cluster-with-extramounts.yaml

$ clusterctl init --infrastructure docker

$ k create ns runtimesdk

$ k apply -f config
```

After this if I see the logs of runtimesdk extension pod

```shell
I0512 11:31:17.640527       1 main.go:105] "main: Starting Runtime Extension server"                                                                                    │
│ I0512 11:31:17.642045       1 server.go:149] "controller-runtime/webhook: Registering webhook" path="/hooks.runtime.cluster.x-k8s.io/v1alpha1/beforeclusterdelete/befor │
│ I0512 11:31:17.642125       1 server.go:149] "controller-runtime/webhook: Registering webhook" path="/hooks.runtime.cluster.x-k8s.io/v1alpha1/discovery"                │
│ I0512 11:31:17.642211       1 server.go:217] "controller-runtime/webhook/webhooks: Starting webhook server"                                                             │
│ I0512 11:31:17.643411       1 certwatcher.go:131] "controller-runtime/certwatcher: Updated current TLS certificate"                                                     │
│ I0512 11:31:17.647581       1 certwatcher.go:85] "controller-runtime/certwatcher: Starting certificate watcher"                                                         │
│ I0512 11:31:17.650214       1 server.go:271] "controller-runtime/webhook: Serving webhook server" host="" port=9443
```
