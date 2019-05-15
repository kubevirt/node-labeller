FROM golang:1.10 AS builder

MAINTAINER "The KubeVirt Project" <kubevirt-dev@googlegroups.com>

WORKDIR /go/src/kubevirt.io/node-labeller

ENV GOPATH=/go

COPY . .

RUN go get -u github.com/golang/dep/cmd/dep && dep ensure

RUN go test ./...

RUN CGO_ENABLED=0 GOOS=linux go build -o /node-labeller cmd/node-labeller/node-labeller.go

FROM registry.access.redhat.com/ubi8/ubi-minimal

RUN mkdir -p /plugin/dest

COPY --from=builder /node-labeller /usr/sbin/node-labeller

ENTRYPOINT [ "/usr/sbin/node-labeller"]

