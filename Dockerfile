FROM gzg1984/dev_ubuntu:latest
LABEL maintainer="Maxpain <g.zg1984@gmail.com>"
RUN git clone http://github.com/gzg1984/entropy.git
WORKDIR /entropy
EXPOSE 22
EXPOSE 8888
ENTRYPOINT ["/entropy/docker_entrypoint.sh"]
