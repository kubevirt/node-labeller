all: docker clean

VERSION="0.0.4"

push-image:
	docker push quay.io/ksimon/kubevirt-cpu-node-labeller:${VERSION}

image: test binary
	docker build -t quay.io/ksimon/kubevirt-cpu-node-labeller:${VERSION} .

binary: dep
	go build cmd/cpu-node-labeller/cpu-node-labeller.go

dep:
	dep ensure

clean:
	rm -f cpu-node-labeller

test:
	go test ./...

.PHONY: all push-image image binary clean test
