# Users Microservice

```bash

yarn run lint

yarn run test
yarn run coverage

export MONGO_CONN_STRING=mongodb://root:admin@172.16.0.3:28017
export MONGO_DATABASE=user_db
export REDIS_CONN_STRING=127.0.0.1:6379

node index.js


# generate Public and Private keys for service
openssl genrsa -out keys/jwtencyptionkey.pem 2048
openssl rsa -pubout -in keys/jwtencyptionkey.pem -out keys/jwtencyptionkey.pub

chmod 400 keys/jwtencyptionkey.{pem,pub}
```
