package usecase

import (
	"fmt"
	"time"
	"context"

	"github.com/juliocesarscheidt/apigateway/infra/adapter"
	userpb "github.com/juliocesarscheidt/apigateway/pb"
)

func GetUserRate(req *userpb.GetUserRateRequest, grpcClient adapter.GrpcClient) (*userpb.UserRate, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	res, err := grpcClient.GetUserRate(ctx, req)
	if err != nil {
		fmt.Errorf("Failed to call endpoint: %v", err)
		return &userpb.UserRate{}, err
	}

	bodyDTO := &userpb.UserRate{}
	if res.GetUserRate() != nil {
		bodyDTO = res.GetUserRate()
	}

	return bodyDTO, nil
}
