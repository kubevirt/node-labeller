VERSION=`git describe --tags`

push-image: image
	docker push quay.io/ksimon/kubevirt-node-labeller:${VERSION}

image: test binary
	docker build -t quay.io/ksimon/kubevirt-node-labeller:${VERSION} .

binary: dep
	go build cmd/kubevirt-node-labeller/kubevirt-node-labeller.go

dep:
	dep ensure

clean:
	rm -f kubevirt-node-labeller

test:
	go test ./...

#deploy-manifest creates new (or replace old one) deploy manifest under folder deploy/<PACKAGE_VERSION>
deploy-manifest:
	ansible-playbook generate_template.yaml

.PHONY: push-image image binary clean test deploy-manifest
