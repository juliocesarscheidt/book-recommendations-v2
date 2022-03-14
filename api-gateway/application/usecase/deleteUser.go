package usecase

import (
	"fmt"
	"context"

	"github.com/juliocesarscheidt/apigateway/infra/adapter"
	userpb "github.com/juliocesarscheidt/apigateway/pb"
)

func DeleteUser(req *userpb.DeleteUserRequest, grpcClient adapter.GrpcClient) (error) {
	_, err := grpcClient.DeleteUser(context.Background(), req)
	if err != nil {
		fmt.Errorf("Failed to call endpoint: %v", err)
		return err
	}

	return nil
}
