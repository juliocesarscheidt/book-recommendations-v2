# Recommendations Microservice

## gRPC

```bash
python3 -m pip install grpcio
python3 -m pip install grpcio-tools

pip freeze | grep grpcio
cat >> requirements.txt <<EOF
grpcio==1.44.0
grpcio-tools==1.44.0
EOF

cd src/recommendations/
mkdir -p pb

# generate protobuf files
python3 -m grpc_tools.protoc \
  -I protofiles/ \
  --python_out=pb/ \
  --grpc_python_out=pb/ \
  protofiles/user.proto
```

```bash
export AMQP_CONN_STRING=amqp://rabbitmq:admin@127.0.0.1:5872/
export REDIS_CONN_STRING=127.0.0.1:6379
export GRPC_CONN_STRING=127.0.0.1:50051

python3 -u src/recommendations/main.py


# black linter
make lint

make test

make report

```
