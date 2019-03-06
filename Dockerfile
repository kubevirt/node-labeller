FROM centos:7

LABEL maintainer="ksimon@redhat.com"

ENV container docker

COPY cpu-node-labeller /usr/sbin/cpu-node-labeller


ENTRYPOINT [ "/usr/sbin/cpu-node-labeller"]
