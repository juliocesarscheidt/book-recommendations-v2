# API Gateway

```bash
go fmt .


go get github.com/golang/protobuf
go get google.golang.org/grpc
go get google.golang.org/grpc/cmd/protoc-gen-go-grpc


export GO111MODULE=on
go mod init github.com/juliocesarscheidt/apigateway
go mod tidy

go mod download


export GRPC_CONN_STRING=127.0.0.1:50051
export AMQP_CONN_STRING=amqp://rabbitmq:admin@127.0.0.1:5872/
export REDIS_CONN_STRING=127.0.0.1:6379

go run main.go

```

## Healthcheck requests

```bash
# healthcheck
curl --silent -X GET --url 'http://localhost:3080/api/healthcheck' | jq -r
```

## AUTH requests

```bash
# sign up
curl --silent -X POST --data '{"name": "julio", "surname": "cesar", "email": "admin@email.com", "phone": "4199887766", "password": "PASSWORD"}' --url 'http://localhost:3080/api/auth/signup' | jq -r

# sign in
curl --silent -X POST --data '{"email": "admin@email.com", "password": "PASSWORD"}' --url 'http://localhost:3080/api/auth/signin' | jq -r

export TOKEN=$(curl --silent -X POST --data '{"email": "admin@email.com", "password": "PASSWORD"}' --url 'http://localhost:3080/api/auth/signin' | jq -r '.data')
echo "${TOKEN}"
```

## USER requests

```bash
# get current user info
curl --silent -X GET -H "Authorization: Bearer ${TOKEN}" --url 'http://localhost:3080/api/user/me' | jq -r

# create user
curl --silent -X POST --data '{"name": "julio", "surname": "cesar", "email": "admin@email.com", "phone": "4199887766", "password": "PASSWORD"}' -H "Authorization: Bearer ${TOKEN}" --url 'http://localhost:3080/api/user' | jq -r

# get user
curl --silent -X GET -H "Authorization: Bearer ${TOKEN}" --url 'http://localhost:3080/api/user/17fc6972e270d41bcd45e9e8' | jq -r

# update user
curl --silent -X PUT --data '{"name": "julio2", "surname": "cesar", "email": "julio2@mail.com", "phone": "4199887766", "password": "PASSWORD"}' -H "Authorization: Bearer ${TOKEN}" --url 'http://localhost:3080/api/user/17fc6972e270d41bcd45e9e8' | jq -r

# delete user
curl --silent -X DELETE -H "Authorization: Bearer ${TOKEN}" --url 'http://localhost:3080/api/user/17fc6972e270d41bcd45e9e8' | jq -r

# list user
curl --silent -X GET -H "Authorization: Bearer ${TOKEN}" --url 'http://localhost:3080/api/user?page=0&size=50' | jq -r

# delete all
for UUID in $(curl --silent -X GET -H "Authorization: Bearer ${TOKEN}" --url 'http://localhost:3080/api/user?page=0&size=50' | jq -r '.data[].uuid'); do
  curl --silent -X DELETE -H "Authorization: Bearer ${TOKEN}" --url "http://localhost:3080/api/user/${UUID}" | jq -r
done



# upsert user rate
curl --silent -X POST --data '{"user_uuid": "17fc6972e270d41bcd45e9e8", "book_uuid": "d3n1t3q6w3f4x3x7c3d7x3l3", "rate": 8}' -H "Authorization: Bearer ${TOKEN}" --url 'http://localhost:3080/api/user/rating' | jq -r

# get user rate
curl --silent -X GET -H "Authorization: Bearer ${TOKEN}" --url 'http://localhost:3080/api/user/rating/17fc6972e270d41bcd45e9e8' | jq -r

# delete user rate
curl --silent -X DELETE -H "Authorization: Bearer ${TOKEN}" --url 'http://localhost:3080/api/user/rating/17fc6972e270d41bcd45e9e8' | jq -r

# list user rate
curl --silent -X GET -H "Authorization: Bearer ${TOKEN}" --url 'http://localhost:3080/api/user/rating?page=0&size=50' | jq -r

```

## BOOK requests

```bash
# create book
curl --silent -X POST --data '{"title": "Clean Architecture", "author": "Robert Martin", "genre": "Software", "image": "https://images-na.ssl-images-amazon.com/images/I/41-sN-mzwKL._SX258_BO1,204,203,200_QL70_ML2_.jpg"}' -H "Authorization: Bearer ${TOKEN}" --url 'http://localhost:3080/api/book' | jq -r

# get book
curl --silent -X GET -H "Authorization: Bearer ${TOKEN}" --url 'http://localhost:3080/api/book/d3n1t3q6w3f4x3x7c3d7x3l3' | jq -r

# update book
curl --silent -X PUT --data '{"title": "Architecture", "author": "Martin", "genre": "Software Architecture", "image": "https://images-na.ssl-images-amazon.com"}' -H "Authorization: Bearer ${TOKEN}" --url 'http://localhost:3080/api/book/d3n1t3q6w3f4x3x7c3d7x3l3' | jq -r

# delete book
curl --silent -X DELETE -H "Authorization: Bearer ${TOKEN}" --url 'http://localhost:3080/api/book/d3n1t3q6w3f4x3x7c3d7x3l3' | jq -r

# list book
curl --silent -X GET -H "Authorization: Bearer ${TOKEN}" --url 'http://localhost:3080/api/book?page=0&size=50' | jq -r

# delete all
for UUID in $(curl --silent -X GET -H "Authorization: Bearer ${TOKEN}" --url 'http://localhost:3080/api/book?page=0&size=50' | jq -r '.data[].uuid'); do
  curl --silent -X DELETE -H "Authorization: Bearer ${TOKEN}" --url "http://localhost:3080/api/book/${UUID}" | jq -r
done

```

## Recommendation requests

```bash
# get recommendations
curl --silent -X GET -H "Authorization: Bearer ${TOKEN}" --url 'http://localhost:3080/api/recommendation/user/17fc6972e270d41bcd45e9e8' | jq -r

```
