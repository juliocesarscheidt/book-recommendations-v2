package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/juliocesarscheidt/apigateway/infra/adapter"
	"github.com/juliocesarscheidt/apigateway/infra/router"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

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
	router.InjectRoutes(r, grpcClient, amqpClient)

	fmt.Println("Listening at 0.0.0.0:3080")
	log.Fatal(http.ListenAndServe("0.0.0.0:3080", r))
}
