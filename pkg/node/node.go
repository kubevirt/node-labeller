/*
 * This file is part of the KubeVirt project
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 * Copyright 2019 Red Hat, Inc.
 */

package node

import (
	"os"
	"strings"

	v1 "k8s.io/api/core/v1"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	k8sclient "k8s.io/client-go/kubernetes"
)

const (
	nodeNameEnv       = "NODE_NAME"
	labelNamespace    = "feature.node.kubernetes.io"
	labellerNamespace = "node-labeller"
)

// GetNode gets node by name
func GetNode(client *k8sclient.Clientset) (*v1.Node, error) {
	nodeName := os.Getenv(nodeNameEnv)

	node, err := client.Core().Nodes().Get(nodeName, meta_v1.GetOptions{})
	if err != nil {
		return nil, err
	}

	return node, nil
}

// AddNodeLabels adds labels to node
func AddNodeLabels(node *v1.Node, labels map[string]string) {
	for name, value := range labels {
		node.Labels[labelNamespace+name] = value
		node.Annotations[labellerNamespace+"-"+labelNamespace+name] = value
	}
}

// GetNodeLabellerLabels gets all labels which were created by cpu-node-labeller
func GetNodeLabellerLabels(node *v1.Node) map[string]bool {
	labellerLabels := make(map[string]bool)
	for key := range node.Annotations {
		if strings.Contains(key, labellerNamespace) {
			delete(node.Annotations, key)
			labellerLabels[key] = true
		}
	}
	return labellerLabels
}

// RemoveCPUModelNodeLabels removes labels from node which were created by cpu-node-labeller
func RemoveCPUModelNodeLabels(node *v1.Node, oldLabels map[string]bool) {
	for label := range node.Labels {
		if ok := oldLabels[label]; ok || strings.Contains(label, labelNamespace+"/cpu-model-") {
			delete(node.Labels, label)
		}
	}
}

// UpdateNode updates node
func UpdateNode(client *k8sclient.Clientset, node *v1.Node) error {
	_, err := client.Core().Nodes().Update(node)
	if err != nil {
		return err
	}

	return nil
}
