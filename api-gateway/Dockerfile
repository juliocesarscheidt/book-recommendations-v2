FROM golang:1.16-alpine as builder
LABEL maintainer="Julio Cesar <julio@blackdevs.com.br>"

WORKDIR /go/src/app

COPY go.mod go.sum /go/src/app/
RUN go mod download

COPY . /go/src/app/

RUN GOOS=linux \
    GOARCH=amd64 \
    GO111MODULE=on \
    CGO_ENABLED=0 \
    go build -o /go/src/app/main .

EXPOSE 3080

ENTRYPOINT [ "/go/src/app/main" ]
