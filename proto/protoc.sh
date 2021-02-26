#!/bin/sh

set -xe

SERVER_OUTPUT_DIR=/output/server
CLIENT_OUTPUT_DIR=/output/client/lib/src/generated

protoc --version
protoc -I=/proto/protos pokemon.proto --go_out=plugins="grpc:${SERVER_OUTPUT_DIR}" \
  --dart_out="grpc:${CLIENT_OUTPUT_DIR}"
