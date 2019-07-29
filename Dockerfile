FROM golang:1.12 AS builder

MAINTAINER "The KubeVirt Project" <kubevirt-dev@googlegroups.com>

WORKDIR /go/src/kubevirt.io/node-labeller

ENV GOPATH=/go

COPY . .

RUN GO111MODULE=on go test ./...

RUN GO111MODULE=on CGO_ENABLED=0 GOOS=linux go build -o /node-labeller cmd/node-labeller/node-labeller.go

FROM registry.access.redhat.com/ubi8/ubi-minimal

COPY --from=builder /node-labeller /usr/sbin/node-labeller

ENTRYPOINT [ "/usr/sbin/node-labeller"]

