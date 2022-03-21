package usecase

import (
	"fmt"
	"time"
	"context"

	"github.com/juliocesarscheidt/apigateway/infra/adapter"
	userpb "github.com/juliocesarscheidt/apigateway/pb"
)

func ListUserRate(req *userpb.ListUserRateRequest, grpcClient adapter.GrpcClient) ([]*userpb.UserRate, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	res, err := grpcClient.ListUserRate(ctx, req)
	if err != nil {
		fmt.Errorf("Failed to call endpoint: %v", err)
		return []*userpb.UserRate{}, err
	}

	bodyDTO := []*userpb.UserRate{}
	if res.GetUserRates() != nil {
		bodyDTO = res.GetUserRates()
	}

	return bodyDTO, nil
}
