VERSION=`git describe --tags`

push-image:
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

.PHONY: push-image image binary clean test
