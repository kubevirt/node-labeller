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
	"testing"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func TestAddNodeLabels(t *testing.T) {
	node := &v1.Node{
		ObjectMeta: metav1.ObjectMeta{
			Labels:      map[string]string{},
			Annotations: map[string]string{},
		},
	}
	labels := map[string]string{
		"fakeLabel": "true",
		"testLabel": "true",
	}
	AddNodeLabels(node, labels)

	if len(node.Labels) != 2 {
		t.Error("node should contain 2 labels")
	}

	if len(node.Annotations) != 2 {
		t.Error("node should contain 2 anotations")
	}
}

func TestGetNodeLabellerLabels(t *testing.T) {
	node := &v1.Node{
		ObjectMeta: metav1.ObjectMeta{
			Annotations: map[string]string{
				"fakeLabel": "true",
				"testLabel": "true",
			},
		},
	}
	labels := GetNodeLabellerLabels(node)

	if len(labels) != 0 {
		t.Error("it should return 0 labels")
	}

	//------------------------------------------

	node.Annotations = map[string]string{
		labellerNamespace + "-fakeLabel": "true",
		"testLabel":                      "true",
	}
	labels = GetNodeLabellerLabels(node)

	if len(labels) != 1 {
		t.Error("it should return 1 label")
	}

	//------------------------------------------

	node.Annotations = map[string]string{
		labellerNamespace + "-fakeLabel": "true",
		labellerNamespace + "-testLabel": "true",
	}
	labels = GetNodeLabellerLabels(node)

	if len(labels) != 2 {
		t.Error("it should return 2 labels")
	}

}

func TestRemoveCPUModelNodeLabels(t *testing.T) {
	labelsToRemove := map[string]bool{
		labellerNamespace + "-fakeLabel": true,
	}
	labels := map[string]string{
		labellerNamespace + "-fakeLabel": "true",
		"testLabel":                      "true",
	}
	node := &v1.Node{
		ObjectMeta: metav1.ObjectMeta{
			Labels:      labels,
			Annotations: labels,
		},
	}
	RemoveCPUModelNodeLabels(node, labelsToRemove)

	if len(node.Labels) != 1 {
		t.Error("it should return 1 labels")
	}

	if _, ok := node.Labels["testLabel"]; !ok {
		t.Error("it should not delete label without special tag")
	}

	//------------------------------------------

	labelsToRemove = map[string]bool{
		labellerNamespace + "-non-existing": true,
	}
	labels = map[string]string{
		"fakeLabel": "true",
		"testLabel": "true",
	}
	node = &v1.Node{
		ObjectMeta: metav1.ObjectMeta{
			Labels:      labels,
			Annotations: labels,
		},
	}
	RemoveCPUModelNodeLabels(node, labelsToRemove)

	if len(node.Labels) != 2 {
		t.Error("it should return 2 labels")
	}

	//------------------------------------------

	labelsToRemove = map[string]bool{
		labellerNamespace + "-fakeLabel": true,
		labellerNamespace + "-testLabel": true,
	}
	labels = map[string]string{
		labellerNamespace + "-fakeLabel": "true",
		labellerNamespace + "-testLabel": "true",
	}
	node = &v1.Node{
		ObjectMeta: metav1.ObjectMeta{
			Labels:      labels,
			Annotations: labels,
		},
	}
	RemoveCPUModelNodeLabels(node, labelsToRemove)

	if len(node.Labels) != 0 {
		t.Error("it should return 0 labels")
	}
}
