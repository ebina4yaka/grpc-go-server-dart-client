
FROM golang:1.16.0

ENV DEBIAN_FRONTEND=noninteractive

ARG PROTO_VERSION=3.15.3

WORKDIR /proto

COPY ./proto .

RUN mkdir /output /output/server /output/client

RUN apt-get -qq update && apt-get -qq install -y \
  unzip

RUN curl -sSL https://github.com/protocolbuffers/protobuf/releases/download/v${PROTO_VERSION}/protoc-${PROTO_VERSION}-linux-x86_64.zip -o protoc.zip && \
  unzip -qq protoc.zip && \
  cp ./bin/protoc /usr/local/bin/protoc && \
  cp -r ./include /usr/local

# Go
RUN go get -u github.com/golang/protobuf/protoc-gen-go

# Dart
RUN apt-get install apt-transport-https
RUN sh -c 'curl https://dl-ssl.google.com/linux/linux_signing_key.pub | apt-key add -'
RUN sh -c 'curl https://storage.googleapis.com/download.dartlang.org/linux/debian/dart_testing.list > /etc/apt/sources.list.d/dart_testing.list'
RUN apt-get update
RUN apt-get install dart -y
ENV PATH="${PATH}:/usr/lib/dart/bin/"
ENV PATH="${PATH}:/root/.pub-cache/bin"
RUN pub global activate protoc_plugin
