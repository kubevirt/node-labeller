FROM fedora:28

LABEL maintainer="ksimon@redhat.com"

ENV container docker

COPY cpu-node-labeller /usr/sbin/cpu-node-labeller


ENTRYPOINT [ "/usr/sbin/cpu-node-labeller"]
