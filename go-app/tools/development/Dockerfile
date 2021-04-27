FROM golang:1.16.3-alpine

ENV APP_DIR=/usr/src/app
WORKDIR $APP_DIR

RUN cd / && \
    go get github.com/githubnemo/CompileDaemon && \
    mv /go /localgo

ENV PATH="/localgo/bin:${PATH}"

ENTRYPOINT ["/usr/src/app/tools/development/run"]

