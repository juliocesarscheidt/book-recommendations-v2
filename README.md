# Book Recommendations

## Application architecture
![Architecture](./architecture/book-recommendations.drawio.png)

## Up and Running

### Locally

```bash

docker-compose config

# infra
docker-compose up -d postgres mongo redis
docker-compose logs -f --tail 100 postgres mongo redis

docker-compose up -d --build rabbitmq
docker-compose logs -f --tail 100 rabbitmq


# services
docker-compose up -d --build api-gateway
docker-compose logs -f --tail 100 api-gateway

docker-compose up -d --build users-microservice
docker-compose logs -f --tail 100 users-microservice

docker-compose up -d --build books-microservice
docker-compose logs -f --tail 100 books-microservice

docker-compose up -d --build recommendations-microservice
docker-compose logs -f --tail 100 recommendations-microservice


docker-compose logs -f --tail 100 \
  api-gateway users-microservice books-microservice recommendations-microservice


# front-end
docker-compose up -d --build client-microservice
docker-compose logs -f --tail 100 client-microservice

```

## Unit Tests

- [x] users-microservice
- [x] books-microservice
- [x] recommendations-microservice
