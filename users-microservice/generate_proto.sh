#!/bin/bash

protoc \
  --js_out=pb \
  --proto_path=./protofiles \
  ./protofiles/*.proto
