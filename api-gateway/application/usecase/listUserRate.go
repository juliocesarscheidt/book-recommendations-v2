package usecase

import (
	"fmt"
	"context"

	"github.com/juliocesarscheidt/apigateway/infra/adapter"
	userpb "github.com/juliocesarscheidt/apigateway/pb"
)

func ListUserRate(req *userpb.ListUserRateRequest, grpcClient adapter.GrpcClient) ([]*userpb.UserRate, error) {
	res, err := grpcClient.ListUserRate(context.Background(), req)
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
