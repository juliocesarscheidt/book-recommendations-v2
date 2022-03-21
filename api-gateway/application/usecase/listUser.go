package usecase

import (
	"fmt"
	"time"
	"context"

	"github.com/juliocesarscheidt/apigateway/infra/adapter"
	userpb "github.com/juliocesarscheidt/apigateway/pb"
)

func ListUser(req *userpb.ListUserRequest, grpcClient adapter.GrpcClient) ([]*userpb.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	res, err := grpcClient.ListUser(ctx, req)
	if err != nil {
		fmt.Errorf("Failed to call endpoint: %v", err)
		return []*userpb.User{}, err
	}

	bodyDTO := []*userpb.User{}
	if res.GetUsers() != nil {
		bodyDTO = res.GetUsers()
	}

	return bodyDTO, nil
}
