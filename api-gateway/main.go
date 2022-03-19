package main

import (
	"os"
	"fmt"
	"log"
	"time"
	"net/http"

	"github.com/juliocesarscheidt/apigateway/infra/adapter"
	"github.com/juliocesarscheidt/apigateway/infra/router"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	var redisConnString = "localhost:6379"
	if os.Getenv("REDIS_CONN_STRING") != "" {
		redisConnString = os.Getenv("REDIS_CONN_STRING")
	}
	redisClient := adapter.GetRedisClient(redisConnString)
	defer redisClient.Client.Close()

	var grpcConnString = "localhost:50051"
	if os.Getenv("GRPC_CONN_STRING") != "" {
		grpcConnString = os.Getenv("GRPC_CONN_STRING")
	}
	grpcClient := adapter.GetGrpcClient(grpcConnString)
	defer grpcClient.Conn.Close()

	var amqpConnString = "amqp://guest:guest@localhost:5672/"
	if os.Getenv("AMQP_CONN_STRING") != "" {
		amqpConnString = os.Getenv("AMQP_CONN_STRING")
	}
	amqpClient := adapter.GetAmqpClient(amqpConnString)
	defer amqpClient.Conn.Close()
	defer amqpClient.Channel.Close()

	r := router.GetRouter()
	router.InjectRoutes(r, grpcClient, amqpClient, redisClient)

	srv := &http.Server{
		Handler:      r,
		Addr:         "0.0.0.0:3080",
		WriteTimeout: 60 * time.Second,
		ReadTimeout:  60 * time.Second,
		IdleTimeout:  60 * time.Second,
	}
	fmt.Println("[INFO] Server listening on 0.0.0.0:3080")
	log.Fatal(srv.ListenAndServe())
}
