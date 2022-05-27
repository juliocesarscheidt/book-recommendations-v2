# Book Recommendations

This project is a mesh of microservices simulating a complex environment with different kinds of protocols to expose/interconnect services, like gRPC and AMQP.

Also it implements an API Gateway to proxy and orquestrate requests through internal services.

The services which use AMQP queue to receive commands, use the request/response pattern.

The main intention it is to calculate and recommend books based on users ratings.

## Application architecture
![Architecture](./architecture/book-recommendations.drawio.svg)

## Up and Running

It is possible to run all services locally using docker and docker-compose.

### Locally

```bash
# infrastructure
docker-compose up -d --build \
  postgres mongo redis rabbitmq

# services
docker-compose up -d --build \
  api-gateway \
  users-microservice \
  books-microservice \
  recommendations-microservice \
  client-microservice
```

### Cloud

The services are going to run on AWS, using ECS with Fargate for containers underlying infrastructure.

Also it will use services like RDS Postgres, DocumentDB, Amazon MQ for Rabbitmq, ElasticCache for Redis, for databases, event broker and cache respectively.

## Cloud architecture
![Architecture](./architecture/book-recommendations-infra.drawio.svg)

Instructions for deploy here: [Instructions](./infrastructure/README.md)

## CI/CD flow architecture
![Architecture](./architecture/book-recommendations-cicd.drawio.svg)
