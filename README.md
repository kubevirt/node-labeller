# cpu-node-labeller

**cpu-node-labeller** 
 uses [cpu-model-nfd-plugin](https://github.com/ksimon1/cpu-model-nfd-plugin) to get all node supported cpu models and then it creates node labells  in format `feature.node.kubernetes.io/cpu-model-<cpuModel>` with these cpu models. Works only when [Kubevirt](https://github.com/kubevirt/kubevirt) is present in the cluster!

**Install**
 ```
 oc apply -f cpu-node-labeller.yaml
 ```