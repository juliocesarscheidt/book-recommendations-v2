#!/bin/bash

python3 -m grpc_tools.protoc \
  -I protofiles/ \
  --python_out=pb/ \
  --grpc_python_out=pb/ \
  protofiles/user.proto
