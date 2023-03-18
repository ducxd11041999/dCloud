#!/bin/bash

protoc -I/usr/local/include -I. \
  -I$GOPATH/src \
  -I$GOPATH/src/github.com/googleapis \
  -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway \
  --go-grpc_out require_unimplemented_servers=false:$GOPATH/src \
  --go_out paths=source_relative:$GOPATH/src \
  --openapiv2_out  $GOPATH/src \
  --openapiv2_opt logtostderr=true \
  --grpc-gateway_out  $GOPATH/src \
  --grpc-gateway_opt logtostderr=true \
  *.proto