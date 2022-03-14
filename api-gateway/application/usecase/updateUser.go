package usecase

import (
	"context"
	"fmt"

	"github.com/juliocesarscheidt/apigateway/infra/adapter"
	userpb "github.com/juliocesarscheidt/apigateway/pb"
)

func UpdateUser(req *userpb.UpdateUserRequest, grpcClient adapter.GrpcClient) (error) {
	_, err := grpcClient.UpdateUser(context.Background(), req)
	if err != nil {
		fmt.Errorf("Failed to call endpoint: %v", err)
		return err
	}

	return nil
}
