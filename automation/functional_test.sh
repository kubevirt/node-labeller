#!/usr/bin/env bash

build_container(){
  docker build -t test/node-labeller:test .
}

get_labeller_pods() {
  oc get pods --field-selector=status.phase!=Running,status.phase!=Succeeded 2>/dev/null| grep node-labeller;
}


deploy_labeller(){
  oc apply -f _out/
  sleep 10
  oc describe ds

  while [[ "$( get_labeller_pods | wc -l)" -ge 1 ]];
  do
    oc get pods
    oc describe pod --selector=app=kubevirt-node-labeller
    sleep 6;
  done
  oc get pods

  oc describe nodes

  exit 0
}

build_container
deploy_labeller
