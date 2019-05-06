#!/usr/bin/env bash

build_container(){
  go get -u github.com/golang/dep/cmd/dep
  dep ensure
  make test
  go build cmd/kubevirt-node-labeller/kubevirt-node-labeller.go
  docker build -t test/kubevirt-node-labeller:test .
}

get_labeller_pods() {
  oc get pods --field-selector=status.phase!=Running,status.phase!=Succeeded 2>/dev/null| grep node-labeller;
}


deploy_labeller(){
  oc create -f deploy/test
  sleep 10
  oc describe ds

  while [[ "$( get_labeller_pods | wc -l)" -ge 1 ]];
  do
    oc get pods
    sleep 6;
  done
  oc get pods

  oc describe nodes
}

check_result(){
  if [ $(oc get nodes -o json | jq '.items[0].metadata.labels' | grep "cpu-model" | wc -l) -eq 0 ]; then
    echo "It should report cpu-models"
    exit 1
  fi

  if [ $(oc get nodes -o json | jq '.items[0].metadata.labels' | grep "cpu-feature" | wc -l) -eq 0 ]; then
    echo "It should report cpu-features"
    exit 1
  fi
}

build_container
deploy_labeller
check_result