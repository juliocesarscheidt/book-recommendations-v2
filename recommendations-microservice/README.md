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

python3 src/recommendations/main.py


# black linter
make lint

make test

make report

```

## Redis

```bash
docker container exec -it redis sh

docker container exec -it redis redis-cli KEYS \*

docker container exec -it redis redis-cli GET "/recommendations/17f9ad3783d07a4f3e3db8ce"

docker container exec -it redis redis-cli GET "/user/bearer/17f9ad3783d07a4f3e3db8ce"

docker container exec -it redis redis-cli

$ redis-cli

KEYS *

GET "/recommendations/17f9ad3783d07a4f3e3db8ce"

DEL "/recommendations/17f9ad3783d07a4f3e3db8ce"

FLUSHALL

FLUSHDB

QUIT
```
