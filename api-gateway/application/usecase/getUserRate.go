package usecase

import (
	"fmt"
	"context"

	"github.com/juliocesarscheidt/apigateway/infra/adapter"
	userpb "github.com/juliocesarscheidt/apigateway/pb"
)

func GetUserRate(req *userpb.GetUserRateRequest, grpcClient adapter.GrpcClient) (*userpb.UserRate, error) {
	res, err := grpcClient.GetUserRate(context.Background(), req)
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
