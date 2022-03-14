package usecase

import (
	"fmt"
	"context"

	"github.com/juliocesarscheidt/apigateway/infra/adapter"
	userpb "github.com/juliocesarscheidt/apigateway/pb"
)

func UpsertUserRate(req *userpb.UpsertUserRateRequest, grpcClient adapter.GrpcClient) (string, error) {
	res, err := grpcClient.UpsertUserRate(context.Background(), req)
	if err != nil {
		fmt.Errorf("Failed to call endpoint: %v", err)
		return "", err
	}

	return res.GetUserUuid(), nil
}
