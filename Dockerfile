FROM registry.access.redhat.com/ubi8/ubi-minimal

LABEL maintainer="ksimon@redhat.com"

ENV container docker

COPY kubevirt-node-labeller /usr/sbin/kubevirt-node-labeller


ENTRYPOINT [ "/usr/sbin/kubevirt-node-labeller"]
