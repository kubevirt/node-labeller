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

package main

import (
	"bytes"
	"flag"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/golang/glog"
	"github.com/ksimon1/cpu-node-labeller/pkg/client"
	"github.com/ksimon1/cpu-node-labeller/pkg/node"
)

func main() {
	fileName := flag.String("fileName", "cpu-model-nfd-plugin", "file Name")
	fileDir := flag.String("fileDir", "/etc/kubernetes/node-feature-discovery/source.d/", "file folder")
	flag.Parse()

	glog.Infof("Running cpu-node-labeller")

	path := filepath.Join(*fileDir, *fileName)
	filestat, err := os.Stat(path)
	if err != nil {
		glog.Fatalf("error while checking file: %s", err)
		os.Exit(1)
	}

	cpuModels := make(map[string]string)
	if filestat.Mode().IsRegular() {
		cmd := exec.Command(path)
		var stdout bytes.Buffer
		var stderr bytes.Buffer
		cmd.Stdout = &stdout
		cmd.Stderr = &stderr

		// Run hook
		err = cmd.Run()

		lines := bytes.Split(stdout.Bytes(), []byte("\n"))
		for _, cpuModel := range lines {
			if len(cpuModel) == 0 {
				continue
			}

			cpuModels[string(cpuModel)] = "true"
		}
	}

	cli, err := client.GetClient()
	if err != nil {
		glog.Fatalf("error while getting client: %s", err)
		os.Exit(1)
	}

	n, err := node.GetNode(cli)
	if err != nil {
		glog.Fatalf("error while getting node: %s", err)
		os.Exit(1)
	}

	node.RemoveCPUModelNodeLabels(n)
	node.AddNodeLabels(n, cpuModels)

	err = node.UpdateNode(cli, n)
	if err != nil {
		glog.Fatalf("error while updating node: %s", err)
		os.Exit(1)
	}

	glog.Info("updated node!")
}
