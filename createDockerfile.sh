#!/bin/sh
SSH_TARGET=`git remote -v|head -n1|awk '{print $2}'`
PROJECT=`basename ${SSH_TARGET}`
DIR=`echo $PROJECT|cut -d"." -f1`

cat <<! > Dockerfile
FROM gzg1984/dev_ubuntu:latest
LABEL maintainer="Maxpain <g.zg1984@gmail.com>"
RUN git clone http://github.com/gzg1984/${PROJECT}
WORKDIR /${DIR}
EXPOSE 22
EXPOSE 8888
ENTRYPOINT ["/${DIR}/docker_entrypoint.sh"]
!
