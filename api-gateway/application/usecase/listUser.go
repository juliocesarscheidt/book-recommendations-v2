package usecase

import (
	"context"
	"fmt"

	"github.com/juliocesarscheidt/apigateway/infra/adapter"
	userpb "github.com/juliocesarscheidt/apigateway/pb"
)

func ListUser(req *userpb.ListUserRequest, grpcClient adapter.GrpcClient) ([]*userpb.User, error) {
	res, err := grpcClient.ListUser(context.Background(), req)
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
