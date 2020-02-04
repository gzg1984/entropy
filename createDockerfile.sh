#!/bin/sh
PROJECT_TARGET=`git remote -v|awk '{print $2}'`
echo ${PROJECT_TARGET}
base=`basename ${PROJECT_TARGET}`
echo ${base}

exit 0
cat <<! > Dockerfile
FROM gzg1984/dev_ubuntu:latest
LABEL maintainer="Maxpain <g.zg1984@gmail.com>"
RUN git clone ${PROJECT_TARGET}
WORKDIR /gang_dockerfile
RUN make
EXPOSE 22
EXPOSE 8888
ENTRYPOINT ["/gang_dockerfile/docker_entrypoint.sh"]
!
