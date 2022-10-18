// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.5
// source: secrets.proto

package secrets

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

// SecretsClient is the client API for Secrets service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SecretsClient interface {
	CreateSecret(ctx context.Context, in *CreateSecretRequest, opts ...grpc.CallOption) (*CreateSecretResponse, error)
	GetSecret(ctx context.Context, in *GetSecretRequest, opts ...grpc.CallOption) (*GetSecretResponse, error)
	GetSecrets(ctx context.Context, in *GetSecretsRequest, opts ...grpc.CallOption) (*GetSecretsResponse, error)
	DeleteSecret(ctx context.Context, in *DeleteSecretRequest, opts ...grpc.CallOption) (*DeleteSecretResponse, error)
}

type secretsClient struct {
	cc grpc.ClientConnInterface
}

func NewSecretsClient(cc grpc.ClientConnInterface) SecretsClient {
	return &secretsClient{cc}
}

func (c *secretsClient) CreateSecret(ctx context.Context, in *CreateSecretRequest, opts ...grpc.CallOption) (*CreateSecretResponse, error) {
	out := new(CreateSecretResponse)
	err := c.cc.Invoke(ctx, "/secrets.Secrets/CreateSecret", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *secretsClient) GetSecret(ctx context.Context, in *GetSecretRequest, opts ...grpc.CallOption) (*GetSecretResponse, error) {
	out := new(GetSecretResponse)
	err := c.cc.Invoke(ctx, "/secrets.Secrets/GetSecret", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *secretsClient) GetSecrets(ctx context.Context, in *GetSecretsRequest, opts ...grpc.CallOption) (*GetSecretsResponse, error) {
	out := new(GetSecretsResponse)
	err := c.cc.Invoke(ctx, "/secrets.Secrets/GetSecrets", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *secretsClient) DeleteSecret(ctx context.Context, in *DeleteSecretRequest, opts ...grpc.CallOption) (*DeleteSecretResponse, error) {
	out := new(DeleteSecretResponse)
	err := c.cc.Invoke(ctx, "/secrets.Secrets/DeleteSecret", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SecretsServer is the server API for Secrets service.
// All implementations must embed UnimplementedSecretsServer
// for forward compatibility
type SecretsServer interface {
	CreateSecret(context.Context, *CreateSecretRequest) (*CreateSecretResponse, error)
	GetSecret(context.Context, *GetSecretRequest) (*GetSecretResponse, error)
	GetSecrets(context.Context, *GetSecretsRequest) (*GetSecretsResponse, error)
	DeleteSecret(context.Context, *DeleteSecretRequest) (*DeleteSecretResponse, error)
	mustEmbedUnimplementedSecretsServer()
}

// UnimplementedSecretsServer must be embedded to have forward compatible implementations.
type UnimplementedSecretsServer struct {
}

func (UnimplementedSecretsServer) CreateSecret(context.Context, *CreateSecretRequest) (*CreateSecretResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateSecret not implemented")
}
func (UnimplementedSecretsServer) GetSecret(context.Context, *GetSecretRequest) (*GetSecretResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSecret not implemented")
}
func (UnimplementedSecretsServer) GetSecrets(context.Context, *GetSecretsRequest) (*GetSecretsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSecrets not implemented")
}
func (UnimplementedSecretsServer) DeleteSecret(context.Context, *DeleteSecretRequest) (*DeleteSecretResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteSecret not implemented")
}
func (UnimplementedSecretsServer) mustEmbedUnimplementedSecretsServer() {}

// UnsafeSecretsServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SecretsServer will
// result in compilation errors.
type UnsafeSecretsServer interface {
	mustEmbedUnimplementedSecretsServer()
}

func RegisterSecretsServer(s grpc.ServiceRegistrar, srv SecretsServer) {
	s.RegisterService(&Secrets_ServiceDesc, srv)
}

func _Secrets_CreateSecret_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateSecretRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SecretsServer).CreateSecret(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/secrets.Secrets/CreateSecret",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SecretsServer).CreateSecret(ctx, req.(*CreateSecretRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Secrets_GetSecret_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSecretRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SecretsServer).GetSecret(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/secrets.Secrets/GetSecret",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SecretsServer).GetSecret(ctx, req.(*GetSecretRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Secrets_GetSecrets_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSecretsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SecretsServer).GetSecrets(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/secrets.Secrets/GetSecrets",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SecretsServer).GetSecrets(ctx, req.(*GetSecretsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Secrets_DeleteSecret_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteSecretRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SecretsServer).DeleteSecret(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/secrets.Secrets/DeleteSecret",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SecretsServer).DeleteSecret(ctx, req.(*DeleteSecretRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Secrets_ServiceDesc is the grpc.ServiceDesc for Secrets service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Secrets_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "secrets.Secrets",
	HandlerType: (*SecretsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateSecret",
			Handler:    _Secrets_CreateSecret_Handler,
		},
		{
			MethodName: "GetSecret",
			Handler:    _Secrets_GetSecret_Handler,
		},
		{
			MethodName: "GetSecrets",
			Handler:    _Secrets_GetSecrets_Handler,
		},
		{
			MethodName: "DeleteSecret",
			Handler:    _Secrets_DeleteSecret_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "secrets.proto",
}
