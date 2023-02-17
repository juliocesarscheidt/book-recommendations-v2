# API Gateway

```bash

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
export BUCKET_NAME=book-recommendations-v2-api-gw-bucket

go run main.go


go vet .

go fmt .

```

## Healthcheck requests

```bash
# healthcheck
curl --silent -X GET --url 'http://localhost:3080/api/healthcheck' | jq -r
```

## AUTH requests

```bash
# sign up
curl --silent -X POST --data '{"name": "julio", "surname": "scheidt", "email": "admin@email.com", "phone": "4199887766", "password": "PASSWORD"}' --url 'http://localhost:3080/api/auth/signup' | jq -r

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
curl --silent -X POST --data '{"name": "julio", "surname": "scheidt", "email": "admin@email.com", "phone": "4199887766", "password": "PASSWORD"}' -H "Authorization: Bearer ${TOKEN}" --url 'http://localhost:3080/api/user' | jq -r

# get user
curl --silent -X GET -H "Authorization: Bearer ${TOKEN}" --url 'http://localhost:3080/api/user/63eee27b6969b9409d2a2579' | jq -r

# update user
curl --silent -X PUT --data '{"name": "julio2", "surname": "scheidt", "email": "julio2@mail.com", "phone": "4199887766", "password": "PASSWORD"}' -H "Authorization: Bearer ${TOKEN}" --url 'http://localhost:3080/api/user/63eee27b6969b9409d2a2579' | jq -r

# delete user
curl --silent -X DELETE -H "Authorization: Bearer ${TOKEN}" --url 'http://localhost:3080/api/user/63eee27b6969b9409d2a2579' | jq -r

# list user
curl --silent -X GET -H "Authorization: Bearer ${TOKEN}" --url 'http://localhost:3080/api/user?page=0&size=50' | jq -r

# delete all
for UUID in $(curl --silent -X GET -H "Authorization: Bearer ${TOKEN}" --url 'http://localhost:3080/api/user?page=0&size=50' | jq -r '.data[].uuid'); do
  curl --silent -X DELETE -H "Authorization: Bearer ${TOKEN}" --url "http://localhost:3080/api/user/${UUID}" | jq -r
done


# upsert user rate
curl --silent -X POST --data '{"user_uuid": "63eee27b6969b9409d2a2579", "book_uuid": "z3j1m3c6w3f5j3r3r3a8v3f5", "rate": 8}' -H "Authorization: Bearer ${TOKEN}" --url 'http://localhost:3080/api/user/rating' | jq -r

# get user rate
curl --silent -X GET -H "Authorization: Bearer ${TOKEN}" --url 'http://localhost:3080/api/user/rating/63eee27b6969b9409d2a2579' | jq -r

# delete user rate
curl --silent -X DELETE -H "Authorization: Bearer ${TOKEN}" --url 'http://localhost:3080/api/user/rating/63eee27b6969b9409d2a2579' | jq -r

# list user rate
curl --silent -X GET -H "Authorization: Bearer ${TOKEN}" --url 'http://localhost:3080/api/user/rating?page=0&size=50' | jq -r
```

## BOOK requests

```bash
# create book
curl -i -X POST \
  -H "Content-Type: multipart/form-data" \
  -F "title=Clean Architecture" \
  -F "author=Robert Martin" \
  -F "genre=Software" \
  -F "image=@$PWD/miscellaneous/clean-arch-cover.jpg" \
  -H "Authorization: Bearer ${TOKEN}" \
  --url 'http://localhost:3080/api/book'

# get book
curl --silent -X GET -H "Authorization: Bearer ${TOKEN}" --url 'http://localhost:3080/api/book/z3j1m3c6w3f5j3r3r3a8v3f5' | jq -r

# get book presign url
curl --silent -X GET -H "Authorization: Bearer ${TOKEN}" --url 'http://localhost:3080/api/book/z3j1m3c6w3f5j3r3r3a8v3f5/image/url' | jq -r


# update book (without image)
curl --silent -X PUT --data '{"title": "Architecture", "author": "Martin", "genre": "Software Architecture"}' -H "Authorization: Bearer ${TOKEN}" --url 'http://localhost:3080/api/book/z3j1m3c6w3f5j3r3r3a8v3f5' | jq -r

# update book (with image)
curl -i -X PUT \
  -H "Content-Type: multipart/form-data" \
  -F "title=Clean Architecture" \
  -F "author=Robert Martin" \
  -F "genre=Software Architecture" \
  -F "image=@$PWD/miscellaneous/clean-arch-cover.jpg" \
  -H "Authorization: Bearer ${TOKEN}" \
  --url 'http://localhost:3080/api/book/z3j1m3c6w3f5j3r3r3a8v3f5/image'


# delete book
curl --silent -X DELETE -H "Authorization: Bearer ${TOKEN}" --url 'http://localhost:3080/api/book/z3j1m3c6w3f5j3r3r3a8v3f5' | jq -r

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
curl --silent -X GET -H "Authorization: Bearer ${TOKEN}" --url 'http://localhost:3080/api/recommendation/user/63eee27b6969b9409d2a2579' | jq -r
```
