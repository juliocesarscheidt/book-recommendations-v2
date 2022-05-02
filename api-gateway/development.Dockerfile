FROM golang:1.16-alpine as runtime
LABEL maintainer="Julio Cesar <julio@blackdevs.com.br>"

WORKDIR /go/src/app

RUN apk update && apk add build-base

COPY go.mod go.sum /go/src/app/
RUN go mod download
COPY . /go/src/app/

EXPOSE 3080

CMD [ "go", "run", "main.go" ]
