#!/usr/bin/env bash

copy_files(){
  rm -rf deploy/$1; mkdir deploy/$1
  cp deploy/templates/roles.yaml deploy/$1/roles.yaml
  cp deploy/templates/kubevirt-node-labeller-ds.yaml deploy/$1/kubevirt-node-labeller-ds.yaml
}

update_files(){
  sed -i "s/<LABELLER_VERSION>/$2/g" deploy/$1/kubevirt-node-labeller-ds.yaml
  sed -i "s/<CPU_PLUGIN_VERSION>/$3/g" deploy/$1/kubevirt-node-labeller-ds.yaml
  sed -i "s/<KVM_INFO_VERSION>/$4/g" deploy/$1/kubevirt-node-labeller-ds.yaml
}
commit_files(){
  git add deploy/.
  git commit --message="Updated version of package to $1"
}

if [ $# -ne 4 ]; then 
  echo "Not enought parameters"
  exit 1
fi

copy_files $1
update_files $@
commit_files $1