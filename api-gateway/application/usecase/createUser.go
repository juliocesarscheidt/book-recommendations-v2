package usecase

import (
	"fmt"
	"time"
	"context"

	"github.com/juliocesarscheidt/apigateway/infra/adapter"
	userpb "github.com/juliocesarscheidt/apigateway/pb"
)

func CreateUser(req *userpb.CreateUserRequest, grpcClient adapter.GrpcClient) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	res, err := grpcClient.CreateUser(ctx, req)
	if err != nil {
		fmt.Errorf("Failed to call endpoint: %v", err)
		return "", err
	}

	return res.GetUuid(), nil
}
