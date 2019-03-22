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
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/golang/glog"
	"kubevirt.io/kubevirt-node-labeller/pkg/client"
	"kubevirt.io/kubevirt-node-labeller/pkg/node"
)

func main() {
	fileDir := flag.String("fileDir", "/etc/kubernetes/node-feature-discovery/source.d/", "file folder")
	flag.Parse()

	glog.Infof("Running kubevirt-node-labeller")

	files, err := ioutil.ReadDir(*fileDir)
	if err != nil {
		glog.Fatalf("could not access directory with files: %s", err)
		os.Exit(1)
	}

	features := make(map[string]string)
	for _, file := range files {
		fileName := file.Name()
		err := runFile(*fileDir, fileName, features)
		if err != nil {
			glog.Warning("could not run file: " + fmt.Sprintf("%s, %s", fileName, err))
			continue
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

	oldLabels := node.GetNodeLabellerLabels(n)
	node.RemoveCPUModelNodeLabels(n, oldLabels)
	node.AddNodeLabels(n, features)

	err = node.UpdateNode(cli, n)
	if err != nil {
		glog.Fatalf("error while updating node: %s", err)
		os.Exit(1)
	}

	glog.Info("node updated!")
}

func runFile(fileDir, fileName string, features map[string]string) error {
	path := filepath.Join(fileDir, fileName)
	filestat, err := os.Stat(path)
	if err != nil {
		return err
	}

	//if file is regular and executable
	if filestat.Mode().IsRegular() && filestat.Mode()&0111 != 0 {
		cmd := exec.Command(path)
		var stdout, stderr bytes.Buffer
		cmd.Stdout = &stdout
		cmd.Stderr = &stderr

		// Run hook
		err = cmd.Run()

		lines := bytes.Split(stdout.Bytes(), []byte("\n"))
		for _, feature := range lines {
			if len(feature) == 0 {
				continue
			}

			features[string(feature)] = "true"
		}
	}
	return nil
}
