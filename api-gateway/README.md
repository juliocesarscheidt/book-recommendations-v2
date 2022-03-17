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

## AUTH requests

```bash
# sign up
curl --silent -X POST --data '{"name": "julio", "surname": "cesar", "email": "julio@mail.com", "phone": "+554199887766", "password": "PASSWORD"}' --url 'http://localhost:3080/v1/auth/signup' | jq -r

# sign in
curl --silent -X POST --data '{"email": "julio@mail.com", "password": "PASSWORD"}' --url 'http://localhost:3080/v1/auth/signin' | jq -r
```

## USER requests

```bash
# create user
curl --silent -X POST --data '{"name": "julio", "surname": "cesar", "email": "julio@mail.com", "phone": "+554199887766", "password": "PASSWORD"}' --url 'http://localhost:3080/v1/user' | jq -r

# get user
curl --silent -X GET --url 'http://localhost:3080/v1/user/17f9055dc9007b0d54d83270' | jq -r

# update user
curl --silent -X PUT --data '{"name":"julio10", "surname": "cesar", "email":"julio10", "phone": "+554199887766", "password": "PASSWORD"}' --url 'http://localhost:3080/v1/user/17f9055dc9007b0d54d83270' | jq -r

# delete user
curl --silent -X DELETE --url 'http://localhost:3080/v1/user/17f9055dc9007b0d54d83270' | jq -r

# list user
curl --silent -X GET -H 'Authorization: Bearer eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1dWlkIjoiMTdmOTYxZTQ1N2UwZjIxNWI2MjEyZDc3IiwibmFtZSI6Imp 1bGlvIiwic3VybmFtZSI6ImNlc2FyIiwiZW1haWwiOiJqdWxpb0BtYWlsLmNvbSIsInBob25lIjoiKzU1NDE5OTg4Nzc2NiIsImlhdCI6MTY0NzQ5MjI2NCwiZXhwIjoxNjQ3NDk1ODY0LCJhdWQiOiJ1cm46Ym9vay1yZWNvbW1lbmRhdGlvbnM6YXBpLWdhdGV3YXkiLCJpc3MiOiJ1cm46Ym9vay1y ZWNvbW1lbmRhdGlvbnM6dXNlcnMtbWljcm9zZXJ2aWNlIn0.UEC5p4VH5VyaWPA9lb8sYThtyhLVp3mLW8XMIy3F4S6UvkFiXkbq98yXLMwa0PxCgJzdTWrqwRpqcHcK25uNuTlkBuYzYhARyhGxW7fuRak-2QTgbjVR6VBxftIMfsw1JFGZpBEJ0v3L4SQsqZMw6N5SjLpTAdoJSzWLf2Mo2dqxOAtcj mqjCtSUBBwIOMzJFh9aQTLzLUJoxoreNE9GNl6zpCb2HifTlfpYReJGuTk77w71uYa3oOyKzEiUwrYOAOsukLcoSTUvha_a6z0aULta4zqiEyoQyPqMOEPIWoqd79NzBuWfLbc6JLdEug_CisLPKlYAbHI_Jy4QCsGWlA' --url 'http://localhost:3080/v1/user?page=0&size=50' | jq -r

# delete all
for UUID in $(curl --silent -X GET --url 'http://localhost:3080/v1/user?page=0&size=50' | jq -r '.data[].uuid'); do
  curl --silent -X DELETE --url "http://localhost:3080/v1/user/${UUID}" | jq -r
done



# upsert user rate
curl --silent -X POST --data '{"user_uuid": "17f9055dc9007b0d54d83270", "book_uuid": "621ff675dc39ed5b1bce10b0", "rate": 8}' --url 'http://localhost:3080/v1/user/rate' | jq -r

# get user rate
curl --silent -X GET --url 'http://localhost:3080/v1/user/rate/17f9055dc9007b0d54d83270' | jq -r

# delete user rate
curl --silent -X DELETE --url 'http://localhost:3080/v1/user/rate/17f9055dc9007b0d54d83270' | jq -r

# list user rate
curl --silent -X GET --url 'http://localhost:3080/v1/user/rate?page=0&size=50' | jq -r

```

## BOOK requests

```bash
# create book
curl --silent -X POST --data '{"title": "Clean Architecture", "author": "Robert Martin", "genre": "Software", "image": "https://images-na.ssl-images-amazon.com/images/I/41-sN-mzwKL._SX258_BO1,204,203,200_QL70_ML2_.jpg"}' --url 'http://localhost:3080/v1/book' | jq -r

# get book
curl --silent -X GET --url 'http://localhost:3080/v1/book/s3q1f3b6l3o4i3q7m3s3u3r8' | jq -r

# update book
curl --silent -X PUT --data '{"title": "Architecture", "author": "Martin", "genre": "Software Architecture", "image": "https://images-na.ssl-images-amazon.com"}' --url 'http://localhost:3080/v1/book/s3q1f3b6l3o4i3q7m3s3u3r8' | jq -r

# delete book
curl --silent -X DELETE --url 'http://localhost:3080/v1/book/s3q1f3b6l3o4i3q7m3s3u3r8' | jq -r

# list book
curl --silent -X GET --url 'http://localhost:3080/v1/book?page=0&size=50' | jq -r

# delete all
for UUID in $(curl --silent -X GET --url 'http://localhost:3080/v1/book?page=0&size=50' | jq -r '.data[].uuid'); do
  curl --silent -X DELETE --url "http://localhost:3080/v1/book/${UUID}" | jq -r
done

```

## Recommendation requests

```bash
# get recommendations
curl --silent -X GET --url 'http://localhost:3080/v1/recommendation/user/17f9055dc9007b0d54d83270' | jq -r

```
