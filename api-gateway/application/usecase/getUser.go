package usecase

import (
	"fmt"
	"time"
	"context"

	"github.com/juliocesarscheidt/apigateway/infra/adapter"
	userpb "github.com/juliocesarscheidt/apigateway/pb"
)

func GetUser(req *userpb.GetUserRequest, grpcClient adapter.GrpcClient) (*userpb.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	res, err := grpcClient.GetUser(ctx, req)
	if err != nil {
		fmt.Errorf("Failed to call endpoint: %v", err)
		return &userpb.User{}, err
	}

	bodyDTO := &userpb.User{}
	if res.GetUser() != nil {
		bodyDTO = res.GetUser()
	}

	return bodyDTO, nil
}
