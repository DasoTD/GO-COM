// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: service_Ecom.proto

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

// EcomClient is the client API for Ecom service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type EcomClient interface {
	CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*CreateUserResponse, error)
	//	rpc UpdateUser (UpdateUserRequest) returns (UpdateUserResponse) {
	//	    option (google.api.http) = {
	//	        patch: "/v1/update_user"
	//	        body: "*"
	//	    };
	//	    option (grpc.gateway.protoc_gen_openapiv2.options.openapiv2_operation) = {
	//	        description: "Use this API to update user";
	//	        summary: "Update user";
	//	    };
	//	}
	LoginUser(ctx context.Context, in *LoginUserRequest, opts ...grpc.CallOption) (*LoginUserResponse, error)
}

type ecomClient struct {
	cc grpc.ClientConnInterface
}

func NewEcomClient(cc grpc.ClientConnInterface) EcomClient {
	return &ecomClient{cc}
}

func (c *ecomClient) CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*CreateUserResponse, error) {
	out := new(CreateUserResponse)
	err := c.cc.Invoke(ctx, "/pb.Ecom/CreateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ecomClient) LoginUser(ctx context.Context, in *LoginUserRequest, opts ...grpc.CallOption) (*LoginUserResponse, error) {
	out := new(LoginUserResponse)
	err := c.cc.Invoke(ctx, "/pb.Ecom/LoginUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EcomServer is the server API for Ecom service.
// All implementations must embed UnimplementedEcomServer
// for forward compatibility
type EcomServer interface {
	CreateUser(context.Context, *CreateUserRequest) (*CreateUserResponse, error)
	//	rpc UpdateUser (UpdateUserRequest) returns (UpdateUserResponse) {
	//	    option (google.api.http) = {
	//	        patch: "/v1/update_user"
	//	        body: "*"
	//	    };
	//	    option (grpc.gateway.protoc_gen_openapiv2.options.openapiv2_operation) = {
	//	        description: "Use this API to update user";
	//	        summary: "Update user";
	//	    };
	//	}
	LoginUser(context.Context, *LoginUserRequest) (*LoginUserResponse, error)
	mustEmbedUnimplementedEcomServer()
}

// UnimplementedEcomServer must be embedded to have forward compatible implementations.
type UnimplementedEcomServer struct {
}

func (UnimplementedEcomServer) CreateUser(context.Context, *CreateUserRequest) (*CreateUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUser not implemented")
}
func (UnimplementedEcomServer) LoginUser(context.Context, *LoginUserRequest) (*LoginUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LoginUser not implemented")
}
func (UnimplementedEcomServer) mustEmbedUnimplementedEcomServer() {}

// UnsafeEcomServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to EcomServer will
// result in compilation errors.
type UnsafeEcomServer interface {
	mustEmbedUnimplementedEcomServer()
}

func RegisterEcomServer(s grpc.ServiceRegistrar, srv EcomServer) {
	s.RegisterService(&Ecom_ServiceDesc, srv)
}

func _Ecom_CreateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EcomServer).CreateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Ecom/CreateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EcomServer).CreateUser(ctx, req.(*CreateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Ecom_LoginUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EcomServer).LoginUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Ecom/LoginUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EcomServer).LoginUser(ctx, req.(*LoginUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Ecom_ServiceDesc is the grpc.ServiceDesc for Ecom service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Ecom_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.Ecom",
	HandlerType: (*EcomServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateUser",
			Handler:    _Ecom_CreateUser_Handler,
		},
		{
			MethodName: "LoginUser",
			Handler:    _Ecom_LoginUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service_Ecom.proto",
}