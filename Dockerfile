ARG BASE_IMAGE=docker.io/library/debian:bullseye-slim

FROM $BASE_IMAGE as app-base

ENV TINI_VERSION v0.19.0
ADD https://github.com/krallin/tini/releases/download/${TINI_VERSION}/tini /tini
RUN chmod +x /tini

ENV DEBIAN_FRONTEND=noninteractive

USER root

RUN groupadd -g 999 books && \
    useradd -r -u 999 -g books books && \
    mkdir -p /home/books && \
    chown books:0 /home/books && \
    chmod g=u /home/books && \
    apt-get update && \
    apt-get clean && \
    rm -rf /var/lib/apt/lists/* /tmp/* /var/tmp/*
 
 ENV USER=books
 USER 999
 WORKDIR /home/books


FROM golang:1.16.6-alpine3.14 as books-build

WORKDIR /go/src/github.com/thiago18l/restful-gin-api

COPY go.mod go.mod
COPY go.sum go.sum

RUN go mod download

COPY ./src ./src
RUN cd ./src && CGO_ENABLE=0 GOOS=linux go build -tags netgo -a -v -installsuffix cgo -o main .


FROM app-base

COPY --from=books-build /go/src/github.com/thiago18l/restful-gin-api/src/main /home/books/main

EXPOSE 8080
ENTRYPOINT [ "/tini", "--" ]
CMD [ "./main" ]

USER 999