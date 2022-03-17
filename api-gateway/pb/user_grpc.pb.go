// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.5.1
// source: user.proto

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// UserServiceClient is the client API for UserService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserServiceClient interface {
	UserSignUp(ctx context.Context, in *UserSignUpRequest, opts ...grpc.CallOption) (*UserSignUpResponse, error)
	UserSignIn(ctx context.Context, in *UserSignInRequest, opts ...grpc.CallOption) (*UserSignInResponse, error)
	CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*CreateUserResponse, error)
	GetUser(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*GetUserResponse, error)
	UpdateUser(ctx context.Context, in *UpdateUserRequest, opts ...grpc.CallOption) (*UpdateUserResponse, error)
	DeleteUser(ctx context.Context, in *DeleteUserRequest, opts ...grpc.CallOption) (*DeleteUserResponse, error)
	ListUser(ctx context.Context, in *ListUserRequest, opts ...grpc.CallOption) (*ListUserResponse, error)
}

type userServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUserServiceClient(cc grpc.ClientConnInterface) UserServiceClient {
	return &userServiceClient{cc}
}

func (c *userServiceClient) UserSignUp(ctx context.Context, in *UserSignUpRequest, opts ...grpc.CallOption) (*UserSignUpResponse, error) {
	out := new(UserSignUpResponse)
	err := c.cc.Invoke(ctx, "/github.com.juliocesarscheidt.usersmicroservice.UserService/UserSignUp", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) UserSignIn(ctx context.Context, in *UserSignInRequest, opts ...grpc.CallOption) (*UserSignInResponse, error) {
	out := new(UserSignInResponse)
	err := c.cc.Invoke(ctx, "/github.com.juliocesarscheidt.usersmicroservice.UserService/UserSignIn", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*CreateUserResponse, error) {
	out := new(CreateUserResponse)
	err := c.cc.Invoke(ctx, "/github.com.juliocesarscheidt.usersmicroservice.UserService/CreateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetUser(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*GetUserResponse, error) {
	out := new(GetUserResponse)
	err := c.cc.Invoke(ctx, "/github.com.juliocesarscheidt.usersmicroservice.UserService/GetUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) UpdateUser(ctx context.Context, in *UpdateUserRequest, opts ...grpc.CallOption) (*UpdateUserResponse, error) {
	out := new(UpdateUserResponse)
	err := c.cc.Invoke(ctx, "/github.com.juliocesarscheidt.usersmicroservice.UserService/UpdateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) DeleteUser(ctx context.Context, in *DeleteUserRequest, opts ...grpc.CallOption) (*DeleteUserResponse, error) {
	out := new(DeleteUserResponse)
	err := c.cc.Invoke(ctx, "/github.com.juliocesarscheidt.usersmicroservice.UserService/DeleteUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) ListUser(ctx context.Context, in *ListUserRequest, opts ...grpc.CallOption) (*ListUserResponse, error) {
	out := new(ListUserResponse)
	err := c.cc.Invoke(ctx, "/github.com.juliocesarscheidt.usersmicroservice.UserService/ListUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServiceServer is the server API for UserService service.
// All implementations must embed UnimplementedUserServiceServer
// for forward compatibility
type UserServiceServer interface {
	UserSignUp(context.Context, *UserSignUpRequest) (*UserSignUpResponse, error)
	UserSignIn(context.Context, *UserSignInRequest) (*UserSignInResponse, error)
	CreateUser(context.Context, *CreateUserRequest) (*CreateUserResponse, error)
	GetUser(context.Context, *GetUserRequest) (*GetUserResponse, error)
	UpdateUser(context.Context, *UpdateUserRequest) (*UpdateUserResponse, error)
	DeleteUser(context.Context, *DeleteUserRequest) (*DeleteUserResponse, error)
	ListUser(context.Context, *ListUserRequest) (*ListUserResponse, error)
	mustEmbedUnimplementedUserServiceServer()
}

// UnimplementedUserServiceServer must be embedded to have forward compatible implementations.
type UnimplementedUserServiceServer struct {
}

func (UnimplementedUserServiceServer) UserSignUp(context.Context, *UserSignUpRequest) (*UserSignUpResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserSignUp not implemented")
}
func (UnimplementedUserServiceServer) UserSignIn(context.Context, *UserSignInRequest) (*UserSignInResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserSignIn not implemented")
}
func (UnimplementedUserServiceServer) CreateUser(context.Context, *CreateUserRequest) (*CreateUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUser not implemented")
}
func (UnimplementedUserServiceServer) GetUser(context.Context, *GetUserRequest) (*GetUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUser not implemented")
}
func (UnimplementedUserServiceServer) UpdateUser(context.Context, *UpdateUserRequest) (*UpdateUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUser not implemented")
}
func (UnimplementedUserServiceServer) DeleteUser(context.Context, *DeleteUserRequest) (*DeleteUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteUser not implemented")
}
func (UnimplementedUserServiceServer) ListUser(context.Context, *ListUserRequest) (*ListUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListUser not implemented")
}
func (UnimplementedUserServiceServer) mustEmbedUnimplementedUserServiceServer() {}

// UnsafeUserServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserServiceServer will
// result in compilation errors.
type UnsafeUserServiceServer interface {
	mustEmbedUnimplementedUserServiceServer()
}

func RegisterUserServiceServer(s grpc.ServiceRegistrar, srv UserServiceServer) {
	s.RegisterService(&UserService_ServiceDesc, srv)
}

func _UserService_UserSignUp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserSignUpRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).UserSignUp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/github.com.juliocesarscheidt.usersmicroservice.UserService/UserSignUp",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).UserSignUp(ctx, req.(*UserSignUpRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_UserSignIn_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserSignInRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).UserSignIn(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/github.com.juliocesarscheidt.usersmicroservice.UserService/UserSignIn",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).UserSignIn(ctx, req.(*UserSignInRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_CreateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).CreateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/github.com.juliocesarscheidt.usersmicroservice.UserService/CreateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).CreateUser(ctx, req.(*CreateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/github.com.juliocesarscheidt.usersmicroservice.UserService/GetUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetUser(ctx, req.(*GetUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_UpdateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).UpdateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/github.com.juliocesarscheidt.usersmicroservice.UserService/UpdateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).UpdateUser(ctx, req.(*UpdateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_DeleteUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).DeleteUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/github.com.juliocesarscheidt.usersmicroservice.UserService/DeleteUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).DeleteUser(ctx, req.(*DeleteUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_ListUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).ListUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/github.com.juliocesarscheidt.usersmicroservice.UserService/ListUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).ListUser(ctx, req.(*ListUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// UserService_ServiceDesc is the grpc.ServiceDesc for UserService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "github.com.juliocesarscheidt.usersmicroservice.UserService",
	HandlerType: (*UserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UserSignUp",
			Handler:    _UserService_UserSignUp_Handler,
		},
		{
			MethodName: "UserSignIn",
			Handler:    _UserService_UserSignIn_Handler,
		},
		{
			MethodName: "CreateUser",
			Handler:    _UserService_CreateUser_Handler,
		},
		{
			MethodName: "GetUser",
			Handler:    _UserService_GetUser_Handler,
		},
		{
			MethodName: "UpdateUser",
			Handler:    _UserService_UpdateUser_Handler,
		},
		{
			MethodName: "DeleteUser",
			Handler:    _UserService_DeleteUser_Handler,
		},
		{
			MethodName: "ListUser",
			Handler:    _UserService_ListUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user.proto",
}

// UserRateServiceClient is the client API for UserRateService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserRateServiceClient interface {
	UpsertUserRate(ctx context.Context, in *UpsertUserRateRequest, opts ...grpc.CallOption) (*UpsertUserRateResponse, error)
	GetUserRate(ctx context.Context, in *GetUserRateRequest, opts ...grpc.CallOption) (*GetUserRateResponse, error)
	DeleteUserRate(ctx context.Context, in *DeleteUserRateRequest, opts ...grpc.CallOption) (*DeleteUserRateResponse, error)
	ListUserRate(ctx context.Context, in *ListUserRateRequest, opts ...grpc.CallOption) (*ListUserRateResponse, error)
}

type userRateServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUserRateServiceClient(cc grpc.ClientConnInterface) UserRateServiceClient {
	return &userRateServiceClient{cc}
}

func (c *userRateServiceClient) UpsertUserRate(ctx context.Context, in *UpsertUserRateRequest, opts ...grpc.CallOption) (*UpsertUserRateResponse, error) {
	out := new(UpsertUserRateResponse)
	err := c.cc.Invoke(ctx, "/github.com.juliocesarscheidt.usersmicroservice.UserRateService/UpsertUserRate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userRateServiceClient) GetUserRate(ctx context.Context, in *GetUserRateRequest, opts ...grpc.CallOption) (*GetUserRateResponse, error) {
	out := new(GetUserRateResponse)
	err := c.cc.Invoke(ctx, "/github.com.juliocesarscheidt.usersmicroservice.UserRateService/GetUserRate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userRateServiceClient) DeleteUserRate(ctx context.Context, in *DeleteUserRateRequest, opts ...grpc.CallOption) (*DeleteUserRateResponse, error) {
	out := new(DeleteUserRateResponse)
	err := c.cc.Invoke(ctx, "/github.com.juliocesarscheidt.usersmicroservice.UserRateService/DeleteUserRate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userRateServiceClient) ListUserRate(ctx context.Context, in *ListUserRateRequest, opts ...grpc.CallOption) (*ListUserRateResponse, error) {
	out := new(ListUserRateResponse)
	err := c.cc.Invoke(ctx, "/github.com.juliocesarscheidt.usersmicroservice.UserRateService/ListUserRate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserRateServiceServer is the server API for UserRateService service.
// All implementations must embed UnimplementedUserRateServiceServer
// for forward compatibility
type UserRateServiceServer interface {
	UpsertUserRate(context.Context, *UpsertUserRateRequest) (*UpsertUserRateResponse, error)
	GetUserRate(context.Context, *GetUserRateRequest) (*GetUserRateResponse, error)
	DeleteUserRate(context.Context, *DeleteUserRateRequest) (*DeleteUserRateResponse, error)
	ListUserRate(context.Context, *ListUserRateRequest) (*ListUserRateResponse, error)
	mustEmbedUnimplementedUserRateServiceServer()
}

// UnimplementedUserRateServiceServer must be embedded to have forward compatible implementations.
type UnimplementedUserRateServiceServer struct {
}

func (UnimplementedUserRateServiceServer) UpsertUserRate(context.Context, *UpsertUserRateRequest) (*UpsertUserRateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpsertUserRate not implemented")
}
func (UnimplementedUserRateServiceServer) GetUserRate(context.Context, *GetUserRateRequest) (*GetUserRateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserRate not implemented")
}
func (UnimplementedUserRateServiceServer) DeleteUserRate(context.Context, *DeleteUserRateRequest) (*DeleteUserRateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteUserRate not implemented")
}
func (UnimplementedUserRateServiceServer) ListUserRate(context.Context, *ListUserRateRequest) (*ListUserRateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListUserRate not implemented")
}
func (UnimplementedUserRateServiceServer) mustEmbedUnimplementedUserRateServiceServer() {}

// UnsafeUserRateServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserRateServiceServer will
// result in compilation errors.
type UnsafeUserRateServiceServer interface {
	mustEmbedUnimplementedUserRateServiceServer()
}

func RegisterUserRateServiceServer(s grpc.ServiceRegistrar, srv UserRateServiceServer) {
	s.RegisterService(&UserRateService_ServiceDesc, srv)
}

func _UserRateService_UpsertUserRate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpsertUserRateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserRateServiceServer).UpsertUserRate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/github.com.juliocesarscheidt.usersmicroservice.UserRateService/UpsertUserRate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserRateServiceServer).UpsertUserRate(ctx, req.(*UpsertUserRateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserRateService_GetUserRate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserRateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserRateServiceServer).GetUserRate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/github.com.juliocesarscheidt.usersmicroservice.UserRateService/GetUserRate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserRateServiceServer).GetUserRate(ctx, req.(*GetUserRateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserRateService_DeleteUserRate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteUserRateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserRateServiceServer).DeleteUserRate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/github.com.juliocesarscheidt.usersmicroservice.UserRateService/DeleteUserRate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserRateServiceServer).DeleteUserRate(ctx, req.(*DeleteUserRateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserRateService_ListUserRate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListUserRateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserRateServiceServer).ListUserRate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/github.com.juliocesarscheidt.usersmicroservice.UserRateService/ListUserRate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserRateServiceServer).ListUserRate(ctx, req.(*ListUserRateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// UserRateService_ServiceDesc is the grpc.ServiceDesc for UserRateService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserRateService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "github.com.juliocesarscheidt.usersmicroservice.UserRateService",
	HandlerType: (*UserRateServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UpsertUserRate",
			Handler:    _UserRateService_UpsertUserRate_Handler,
		},
		{
			MethodName: "GetUserRate",
			Handler:    _UserRateService_GetUserRate_Handler,
		},
		{
			MethodName: "DeleteUserRate",
			Handler:    _UserRateService_DeleteUserRate_Handler,
		},
		{
			MethodName: "ListUserRate",
			Handler:    _UserRateService_ListUserRate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user.proto",
}
