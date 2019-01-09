all: docker clean

push:
	docker push quay.io/ksimon/cpu-node-labeller:latest

docker: binary
	docker build -t quay.io/ksimon/cpu-node-labeller:latest .

binary: dep
	go build cmd/cpu-node-labeller/cpu-node-labeller.go

dep:
	dep ensure

clean:
	rm -f cpu-node-labeller

.PHONY: all push docker binary clean
