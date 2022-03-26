# Book Recommendations

## Application architecture
![Architecture](./architecture/book-recommendations.drawio.png)

## Infrastructure architecture
![Architecture](./architecture/book-recommendations-infra.drawio.png)

## Up and Running

### Locally

```bash
# infra
docker-compose up -d postgres mongo redis rabbitmq
# infra logs
docker-compose logs -f --tail 100 postgres mongo redis rabbitmq

# services
docker-compose up -d --build \
  api-gateway \
  users-microservice \
  books-microservice \
  recommendations-microservice \
  client-microservice

# services logs
docker-compose logs -f --tail 100 \
  api-gateway users-microservice \
  books-microservice \
  recommendations-microservice \
  client-microservice
```

## Unit Tests

- [x] users-microservice
- [x] books-microservice
- [x] recommendations-microservice
