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

go run main.go
```

## USER requests

```bash
curl --silent -X POST --data '{"name": "julio", "surname": "cesar", "email": "julio@mail.com", "phone": "+554199887766"}' --url 'http://localhost:3080/v1/user' | jq -r

curl --silent -X GET --url 'http://localhost:3080/v1/user/17f85d7f76e0935e50a781b9' | jq -r

curl --silent -X PUT --data '{"name":"julio10", "surname": "cesar", "email":"julio10", "phone": "+554199887766"}' --url 'http://localhost:3080/v1/user/17f85d7f76e0935e50a781b9' | jq -r

curl --silent -X DELETE --url 'http://localhost:3080/v1/user/17f85d7f76e0935e50a781b9' | jq -r

curl --silent -X GET --url 'http://localhost:3080/v1/user?page=0&size=50' | jq -r

# delete all
for UUID in $(curl --silent -X GET --url 'http://localhost:3080/v1/user?page=0&size=50' | jq -r '.data[].uuid'); do
  curl --silent -X DELETE --url "http://localhost:3080/v1/user/${UUID}" | jq -r
done



# upsert user rate
curl --silent -X POST --data '{"user_uuid": "17f85d7f76e0935e50a781b9", "book_uuid": "621ff675dc39ed5b1bce10b0", "rate": 8}' --url 'http://localhost:3080/v1/user/rate' | jq -r

# get user rate
curl --silent -X GET --url 'http://localhost:3080/v1/user/rate/17f85d7f76e0935e50a781b9' | jq -r

# delete user rate
curl --silent -X DELETE --url 'http://localhost:3080/v1/user/rate/17f85d7f76e0935e50a781b9' | jq -r

# list user rate
curl --silent -X GET --url 'http://localhost:3080/v1/user/rate?page=0&size=50' | jq -r

```

## BOOK requests

```bash

curl --silent -X POST --data '{"title": "Clean Architecture", "author": "Robert Martin", "genre": "Software", "image": "https://images-na.ssl-images-amazon.com/images/I/41-sN-mzwKL._SX258_BO1,204,203,200_QL70_ML2_.jpg"}' --url 'http://localhost:3080/v1/book' | jq -r

curl --silent -X GET --url 'http://localhost:3080/v1/book/i3o1d3p6o3i4z3z7j3k2n3f1' | jq -r

curl --silent -X PUT --data '{"title": "Architecture", "author": "Martin", "genre": "Software Architecture", "image": "https://images-na.ssl-images-amazon.com"}' --url 'http://localhost:3080/v1/book/i3o1d3p6o3i4z3z7j3k2n3f1' | jq -r

curl --silent -X DELETE --url 'http://localhost:3080/v1/book/i3o1d3p6o3i4z3z7j3k2n3f1' | jq -r

curl --silent -X GET --url 'http://localhost:3080/v1/book?page=0&size=50' | jq -r

# delete all
for UUID in $(curl --silent -X GET --url 'http://localhost:3080/v1/book?page=0&size=50' | jq -r '.data[].uuid'); do
  curl --silent -X DELETE --url "http://localhost:3080/v1/book/${UUID}" | jq -r
done

```

## Recommendation requests

```bash
# get recommendations
curl --silent -X GET --url 'http://localhost:3080/v1/recommendation/user/17f85d7f76e0935e50a781b9' | jq -r

```
