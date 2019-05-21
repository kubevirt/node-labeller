VERSION=`git describe --tags`

push-image:
	docker push quay.io/kubevirt/node-labeller:${VERSION}

image:
	docker build -t quay.io/kubevirt/node-labeller:${VERSION} .

binary: dep
	go build cmd/node-labeller/node-labeller.go

dep:
	dep ensure

clean:
	rm -f node-labeller

test:
	go test ./...

#deploy-manifest creates new (or replace old one) deploy manifest under folder deploy/<PACKAGE_VERSION>
#PACKAGE_VERSION represents version of whole package(labeller, cpu plugin, kvm info, ...)
#LABELLER_VERSION represents version of kubevirt-node-labeller (https://github.com/kubevirt/node-labeller)
#CPU_PLUGIN_VERSION represents version of kubevirt-cpu-nfd-plugin (https://github.com/kubevirt/cpu-nfd-plugin)
#KVM_INFO_VERSION represents version of kvm-info-nfd-plugin (https://github.com/fromanirh/kvm-info-nfd-plugin)
#!!! ALL params are required !!!
#example: make deploy-manifest PACKAGE_VERSION=v0.0.5 LABELLER_VERSION=v0.0.5 CPU_PLUGIN_VERSION=v0.0.4 KVM_INFO_VERSION=v0.4.0
deploy-manifest:
	deploy/templates/release.sh $(PACKAGE_VERSION) $(LABELLER_VERSION) $(CPU_PLUGIN_VERSION) $(KVM_INFO_VERSION)

.PHONY: push-image image binary clean test deploy-manifest
