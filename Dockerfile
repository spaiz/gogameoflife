FROM golang:1.8

ADD . ${GOPATH}/src/github.com/spaiz/gogameoflife/
ADD ./conf/docker/entrypoint.sh /entrypoint.sh
RUN chmod +x /entrypoint.sh
ENTRYPOINT ["/entrypoint.sh"]