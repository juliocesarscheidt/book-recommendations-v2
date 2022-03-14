package usecase

import (
	"fmt"
	"context"

	"github.com/juliocesarscheidt/apigateway/infra/adapter"
	userpb "github.com/juliocesarscheidt/apigateway/pb"
)

func DeleteUserRate(req *userpb.DeleteUserRateRequest, grpcClient adapter.GrpcClient) (error) {
	_, err := grpcClient.DeleteUserRate(context.Background(), req)
	if err != nil {
		fmt.Errorf("Failed to call endpoint: %v", err)
		return err
	}

	return nil
}
