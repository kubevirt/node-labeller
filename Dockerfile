FROM centos:7

LABEL maintainer="ksimon@redhat.com"

ENV container docker

RUN yum -y update
RUN yum clean all

COPY cpu-node-labeller /usr/sbin/cpu-node-labeller


ENTRYPOINT [ "/usr/sbin/cpu-node-labeller"]
