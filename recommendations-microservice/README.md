# Recommendations Microservice

```bash
export AMQP_CONN_STRING=amqp://rabbitmq:admin@127.0.0.1:5872/

export REDIS_HOST=127.0.0.1
export REDIS_PORT=6379

export API_GATEWAY_CONN_STRING=http://127.0.0.1:3080

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

docker container exec -it redis redis-cli GET "/recommendations/17f182b0dfc002737d507f74"

docker container exec -it redis redis-cli

$ redis-cli

KEYS *

GET "/recommendations/17f182b0dfc002737d507f74"

DEL "/recommendations/17f182b0dfc002737d507f74"

FLUSHALL
```
