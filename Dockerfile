FROM centos:7

LABEL maintainer="ksimon@redhat.com"

ENV container docker

RUN yum -y update
RUN yum clean all

COPY kubevirt-node-labeller /usr/sbin/kubevirt-node-labeller


ENTRYPOINT [ "/usr/sbin/kubevirt-node-labeller"]
