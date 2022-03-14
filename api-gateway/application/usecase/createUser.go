package usecase

import (
	"context"
	"fmt"

	"github.com/juliocesarscheidt/apigateway/infra/adapter"
	userpb "github.com/juliocesarscheidt/apigateway/pb"
)

func CreateUser(req *userpb.CreateUserRequest, grpcClient adapter.GrpcClient) (string, error) {
	res, err := grpcClient.CreateUser(context.Background(), req)
	if err != nil {
		fmt.Errorf("Failed to call endpoint: %v", err)
		return "", err
	}

	return res.GetUuid(), nil
}
