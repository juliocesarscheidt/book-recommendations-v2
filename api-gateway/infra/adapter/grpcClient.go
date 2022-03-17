package adapter

import (
	"log"
	"context"

	"google.golang.org/grpc"

	userpb "github.com/juliocesarscheidt/apigateway/pb"
)

type GrpcClient struct {
	Conn *grpc.ClientConn
	UserServiceClient userpb.UserServiceClient
	UserRateServiceClient userpb.UserRateServiceClient
}

func GetGrpcClient(grpcConnString string) GrpcClient {
	conn, userServiceClient, userRateServiceClient, _ := GetGrpcConnAndServiceClient(grpcConnString)

	return GrpcClient{
		Conn: conn,
		UserServiceClient: userServiceClient,
		UserRateServiceClient: userRateServiceClient,
	}
}

func GetGrpcConnAndServiceClient(grpcConnString string) (*grpc.ClientConn, userpb.UserServiceClient, userpb.UserRateServiceClient, error) {
	conn, err := grpc.Dial(grpcConnString, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect to gRPC endpoint: %v", err)
		return nil, nil, nil, err
	}

	return conn, userpb.NewUserServiceClient(conn), userpb.NewUserRateServiceClient(conn), nil
}

// auth
func (client GrpcClient) UserSignUp(ctx context.Context, in *userpb.UserSignUpRequest, opts ...grpc.CallOption) (*userpb.UserSignUpResponse, error) {
	return client.UserServiceClient.UserSignUp(ctx, in, opts...)
}

func (client GrpcClient) UserSignIn(ctx context.Context, in *userpb.UserSignInRequest, opts ...grpc.CallOption) (*userpb.UserSignInResponse, error) {
	return client.UserServiceClient.UserSignIn(ctx, in, opts...)
}

// user
func (client GrpcClient) CreateUser(ctx context.Context, in *userpb.CreateUserRequest, opts ...grpc.CallOption) (*userpb.CreateUserResponse, error) {
	return client.UserServiceClient.CreateUser(ctx, in, opts...)
}

func (client GrpcClient) GetUser(ctx context.Context, in *userpb.GetUserRequest, opts ...grpc.CallOption) (*userpb.GetUserResponse, error) {
	return client.UserServiceClient.GetUser(ctx, in, opts...)
}

func (client GrpcClient) UpdateUser(ctx context.Context, in *userpb.UpdateUserRequest, opts ...grpc.CallOption) (*userpb.UpdateUserResponse, error) {
	return client.UserServiceClient.UpdateUser(ctx, in, opts...)
}

func (client GrpcClient) DeleteUser(ctx context.Context, in *userpb.DeleteUserRequest, opts ...grpc.CallOption) (*userpb.DeleteUserResponse, error) {
	return client.UserServiceClient.DeleteUser(ctx, in, opts...)
}

func (client GrpcClient) ListUser(ctx context.Context, in *userpb.ListUserRequest, opts ...grpc.CallOption) (*userpb.ListUserResponse, error) {
	return client.UserServiceClient.ListUser(ctx, in, opts...)
}

// user rate
func (client GrpcClient) UpsertUserRate(ctx context.Context, in *userpb.UpsertUserRateRequest, opts ...grpc.CallOption) (*userpb.UpsertUserRateResponse, error) {
	return client.UserRateServiceClient.UpsertUserRate(ctx, in, opts...)
}

func (client GrpcClient) GetUserRate(ctx context.Context, in *userpb.GetUserRateRequest, opts ...grpc.CallOption) (*userpb.GetUserRateResponse, error) {
	return client.UserRateServiceClient.GetUserRate(ctx, in, opts...)
}

func (client GrpcClient) DeleteUserRate(ctx context.Context, in *userpb.DeleteUserRateRequest, opts ...grpc.CallOption) (*userpb.DeleteUserRateResponse, error) {
	return client.UserRateServiceClient.DeleteUserRate(ctx, in, opts...)
}

func (client GrpcClient) ListUserRate(ctx context.Context, in *userpb.ListUserRateRequest, opts ...grpc.CallOption) (*userpb.ListUserRateResponse, error) {
	return client.UserRateServiceClient.ListUserRate(ctx, in, opts...)
}
