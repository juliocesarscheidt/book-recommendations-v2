# Books Microservice

```bash

make package

make build

make test

make verify

make report


# run
docker-compose up -d --build books-microservice
docker-compose logs -f --tail 100 books-microservice

```

> application.properties

```bash
spring.application.name=books-microservice

# spring datasource config
spring.datasource.url=jdbc:postgresql://${POSTGRES_HOST}:${POSTGRES_PORT}/${POSTGRES_DB}
spring.datasource.driverClassName=org.postgresql.Driver
spring.datasource.username=${POSTGRES_USER}
spring.datasource.password=${POSTGRES_PASS}

# spring jpa config
spring.jpa.properties.hibernate.generate_statistics: false
spring.jpa.properties.hibernate.format_sql: true
spring.jpa.show-sql: true
spring.jpa.open-in-view: false

# jackson config
spring.jackson.default-property-inclusion = NON_NULL

# connections config
amqp.conn_string=${AMQP_CONN_STRING}
amqp.books_queue=books_queue

# pool config
spring.datasource.hikari.connectionTimeout=30000
spring.datasource.hikari.maximumPoolSize=10
```
