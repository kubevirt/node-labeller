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
deploy-manifest:
	ansible-playbook generate_deploy_manifest.yaml

.PHONY: push-image image binary clean test deploy-manifest
