FROM golang:1.20.4-alpine3.17
RUN apk update && apk add git curl alpine-sdk
RUN mkdir /go/src/app
WORKDIR /go/src/app
ADD . /go/src/app
# 必要なパッケージをインストール
# protoc
RUN apk add --no-cache protobuf-dev