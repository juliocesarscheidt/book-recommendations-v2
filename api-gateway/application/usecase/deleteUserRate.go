package usecase

import (
	"fmt"
	"time"
	"context"

	"github.com/juliocesarscheidt/apigateway/infra/adapter"
	userpb "github.com/juliocesarscheidt/apigateway/pb"
)

func DeleteUserRate(req *userpb.DeleteUserRateRequest, grpcClient adapter.GrpcClient) (error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	_, err := grpcClient.DeleteUserRate(ctx, req)
	if err != nil {
		fmt.Errorf("Failed to call endpoint: %v", err)
		return err
	}

	return nil
}
