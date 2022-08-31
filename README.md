# node-labeller

**NOTE:** All the functionality provided by ths component has been merged into KubeVirt (kubevirt/kubevirt) and thus this component is now obsolete.

**node-labeller** 
 uses [cpu-nfd-plugin](https://github.com/kubevirt/cpu-nfd-plugin) to get all node supported cpu models, features and then it creates node labells  in format `feature.node.kubernetes.io/cpu-model-<cpuModel>` with these cpu models and `feature.node.kubernetes.io/cpu-feature-<cpuFeature>`. Works only when [Kubevirt](https://github.com/kubevirt/kubevirt) is present in the cluster!

**Install**
 ```
 oc apply -f deploy/<version>/
 ```
