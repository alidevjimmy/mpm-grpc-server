// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package userpb

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
	Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error)
	// get all logs which are relevant to specified user with "id"
	GetUserLogs(ctx context.Context, in *GetUserLogsRequest, opts ...grpc.CallOption) (UserService_GetUserLogsClient, error)
	StoreNewToken(ctx context.Context, in *StoreNewTokenRequest, opts ...grpc.CallOption) (*StoreNewTokenResponse, error)
}

type userServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUserServiceClient(cc grpc.ClientConnInterface) UserServiceClient {
	return &userServiceClient{cc}
}

func (c *userServiceClient) Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error) {
	out := new(LoginResponse)
	err := c.cc.Invoke(ctx, "/user.UserService/Login", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetUserLogs(ctx context.Context, in *GetUserLogsRequest, opts ...grpc.CallOption) (UserService_GetUserLogsClient, error) {
	stream, err := c.cc.NewStream(ctx, &UserService_ServiceDesc.Streams[0], "/user.UserService/GetUserLogs", opts...)
	if err != nil {
		return nil, err
	}
	x := &userServiceGetUserLogsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type UserService_GetUserLogsClient interface {
	Recv() (*GetUserLogsResponse, error)
	grpc.ClientStream
}

type userServiceGetUserLogsClient struct {
	grpc.ClientStream
}

func (x *userServiceGetUserLogsClient) Recv() (*GetUserLogsResponse, error) {
	m := new(GetUserLogsResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *userServiceClient) StoreNewToken(ctx context.Context, in *StoreNewTokenRequest, opts ...grpc.CallOption) (*StoreNewTokenResponse, error) {
	out := new(StoreNewTokenResponse)
	err := c.cc.Invoke(ctx, "/user.UserService/StoreNewToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServiceServer is the server API for UserService service.
// All implementations must embed UnimplementedUserServiceServer
// for forward compatibility
type UserServiceServer interface {
	Login(context.Context, *LoginRequest) (*LoginResponse, error)
	// get all logs which are relevant to specified user with "id"
	GetUserLogs(*GetUserLogsRequest, UserService_GetUserLogsServer) error
	StoreNewToken(context.Context, *StoreNewTokenRequest) (*StoreNewTokenResponse, error)
	mustEmbedUnimplementedUserServiceServer()
}

// UnimplementedUserServiceServer must be embedded to have forward compatible implementations.
type UnimplementedUserServiceServer struct {
}

func (UnimplementedUserServiceServer) Login(context.Context, *LoginRequest) (*LoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedUserServiceServer) GetUserLogs(*GetUserLogsRequest, UserService_GetUserLogsServer) error {
	return status.Errorf(codes.Unimplemented, "method GetUserLogs not implemented")
}
func (UnimplementedUserServiceServer) StoreNewToken(context.Context, *StoreNewTokenRequest) (*StoreNewTokenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StoreNewToken not implemented")
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

func _UserService_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserService/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).Login(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetUserLogs_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GetUserLogsRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(UserServiceServer).GetUserLogs(m, &userServiceGetUserLogsServer{stream})
}

type UserService_GetUserLogsServer interface {
	Send(*GetUserLogsResponse) error
	grpc.ServerStream
}

type userServiceGetUserLogsServer struct {
	grpc.ServerStream
}

func (x *userServiceGetUserLogsServer) Send(m *GetUserLogsResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _UserService_StoreNewToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StoreNewTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).StoreNewToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserService/StoreNewToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).StoreNewToken(ctx, req.(*StoreNewTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// UserService_ServiceDesc is the grpc.ServiceDesc for UserService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "user.UserService",
	HandlerType: (*UserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Login",
			Handler:    _UserService_Login_Handler,
		},
		{
			MethodName: "StoreNewToken",
			Handler:    _UserService_StoreNewToken_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetUserLogs",
			Handler:       _UserService_GetUserLogs_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "src/pbs/userpb/user.proto",
}