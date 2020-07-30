FROM alpine:latest
RUN apt update && apt install git

FROM golang:1.14.6-alpine3.12
RUN mkdir /go/src/work
WORKDIR /go/src/work
ADD . /go/src/work