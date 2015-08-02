FROM docker.sendgrid.net/sendgrid/dev_go

ADD . /opt/go/src/github.com/sendgrid/http-proxy
WORKDIR /opt/go/src/github.com/sendgrid/http-proxy
ENV GOPATH /opt/go/src/github.com/sendgrid/http-proxy/Godeps/_workspace:$GOPATH

RUN echo yes | yum install glibc-static.x86_64 zlib-static.x86_64

CMD ["./bin/start"]

EXPOSE 50260 \
       50263

ENV SERVICE_NAME=http-proxy \
    SERVICE_50263_TAGS=healthcheck \
    SERVICE_50263_CHECK_CMD=./bin/healthcheck \
    HTTP-PROXY_PORT=50260 \
    HTTP-PROXY_HEALTHCHECK_PORT=50263 \
    HTTP-PROXY_REQUEST_TIMEOUT=500 \
    HTTP-PROXY_MAINT_FILE=/var/tmp/http-proxy.maint
