#!/usr/bin/env bash

build_container(){
  docker build -t test/node-labeller:test .
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
    oc describe pod --selector=app=node-labeller
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