// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package main

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

// GoServiceClient is the client API for GoService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GoServiceClient interface {
	Increment(ctx context.Context, in *EmptyMessage, opts ...grpc.CallOption) (*MasterMessage, error)
}

type goServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewGoServiceClient(cc grpc.ClientConnInterface) GoServiceClient {
	return &goServiceClient{cc}
}

func (c *goServiceClient) Increment(ctx context.Context, in *EmptyMessage, opts ...grpc.CallOption) (*MasterMessage, error) {
	out := new(MasterMessage)
	err := c.cc.Invoke(ctx, "/GoService/Increment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GoServiceServer is the server API for GoService service.
// All implementations must embed UnimplementedGoServiceServer
// for forward compatibility
type GoServiceServer interface {
	Increment(context.Context, *EmptyMessage) (*MasterMessage, error)
	mustEmbedUnimplementedGoServiceServer()
}

// UnimplementedGoServiceServer must be embedded to have forward compatible implementations.
type UnimplementedGoServiceServer struct {
}

func (UnimplementedGoServiceServer) Increment(context.Context, *EmptyMessage) (*MasterMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Increment not implemented")
}
func (UnimplementedGoServiceServer) mustEmbedUnimplementedGoServiceServer() {}

// UnsafeGoServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GoServiceServer will
// result in compilation errors.
type UnsafeGoServiceServer interface {
	mustEmbedUnimplementedGoServiceServer()
}

func RegisterGoServiceServer(s grpc.ServiceRegistrar, srv GoServiceServer) {
	s.RegisterService(&GoService_ServiceDesc, srv)
}

func _GoService_Increment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmptyMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GoServiceServer).Increment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/GoService/Increment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GoServiceServer).Increment(ctx, req.(*EmptyMessage))
	}
	return interceptor(ctx, in, info, handler)
}

// GoService_ServiceDesc is the grpc.ServiceDesc for GoService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GoService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "GoService",
	HandlerType: (*GoServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Increment",
			Handler:    _GoService_Increment_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "mastermodel.proto",
}

// JavaServiceClient is the client API for JavaService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type JavaServiceClient interface {
	Increment(ctx context.Context, in *EmptyMessage, opts ...grpc.CallOption) (*MasterMessage, error)
}

type javaServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewJavaServiceClient(cc grpc.ClientConnInterface) JavaServiceClient {
	return &javaServiceClient{cc}
}

func (c *javaServiceClient) Increment(ctx context.Context, in *EmptyMessage, opts ...grpc.CallOption) (*MasterMessage, error) {
	out := new(MasterMessage)
	err := c.cc.Invoke(ctx, "/JavaService/Increment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// JavaServiceServer is the server API for JavaService service.
// All implementations must embed UnimplementedJavaServiceServer
// for forward compatibility
type JavaServiceServer interface {
	Increment(context.Context, *EmptyMessage) (*MasterMessage, error)
	mustEmbedUnimplementedJavaServiceServer()
}

// UnimplementedJavaServiceServer must be embedded to have forward compatible implementations.
type UnimplementedJavaServiceServer struct {
}

func (UnimplementedJavaServiceServer) Increment(context.Context, *EmptyMessage) (*MasterMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Increment not implemented")
}
func (UnimplementedJavaServiceServer) mustEmbedUnimplementedJavaServiceServer() {}

// UnsafeJavaServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to JavaServiceServer will
// result in compilation errors.
type UnsafeJavaServiceServer interface {
	mustEmbedUnimplementedJavaServiceServer()
}

func RegisterJavaServiceServer(s grpc.ServiceRegistrar, srv JavaServiceServer) {
	s.RegisterService(&JavaService_ServiceDesc, srv)
}

func _JavaService_Increment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmptyMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JavaServiceServer).Increment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/JavaService/Increment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JavaServiceServer).Increment(ctx, req.(*EmptyMessage))
	}
	return interceptor(ctx, in, info, handler)
}

// JavaService_ServiceDesc is the grpc.ServiceDesc for JavaService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var JavaService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "JavaService",
	HandlerType: (*JavaServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Increment",
			Handler:    _JavaService_Increment_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "mastermodel.proto",
}

// DotnetServiceClient is the client API for DotnetService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DotnetServiceClient interface {
	Increment(ctx context.Context, in *EmptyMessage, opts ...grpc.CallOption) (*MasterMessage, error)
}

type dotnetServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDotnetServiceClient(cc grpc.ClientConnInterface) DotnetServiceClient {
	return &dotnetServiceClient{cc}
}

func (c *dotnetServiceClient) Increment(ctx context.Context, in *EmptyMessage, opts ...grpc.CallOption) (*MasterMessage, error) {
	out := new(MasterMessage)
	err := c.cc.Invoke(ctx, "/DotnetService/Increment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DotnetServiceServer is the server API for DotnetService service.
// All implementations must embed UnimplementedDotnetServiceServer
// for forward compatibility
type DotnetServiceServer interface {
	Increment(context.Context, *EmptyMessage) (*MasterMessage, error)
	mustEmbedUnimplementedDotnetServiceServer()
}

// UnimplementedDotnetServiceServer must be embedded to have forward compatible implementations.
type UnimplementedDotnetServiceServer struct {
}

func (UnimplementedDotnetServiceServer) Increment(context.Context, *EmptyMessage) (*MasterMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Increment not implemented")
}
func (UnimplementedDotnetServiceServer) mustEmbedUnimplementedDotnetServiceServer() {}

// UnsafeDotnetServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DotnetServiceServer will
// result in compilation errors.
type UnsafeDotnetServiceServer interface {
	mustEmbedUnimplementedDotnetServiceServer()
}

func RegisterDotnetServiceServer(s grpc.ServiceRegistrar, srv DotnetServiceServer) {
	s.RegisterService(&DotnetService_ServiceDesc, srv)
}

func _DotnetService_Increment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmptyMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DotnetServiceServer).Increment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/DotnetService/Increment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DotnetServiceServer).Increment(ctx, req.(*EmptyMessage))
	}
	return interceptor(ctx, in, info, handler)
}

// DotnetService_ServiceDesc is the grpc.ServiceDesc for DotnetService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DotnetService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "DotnetService",
	HandlerType: (*DotnetServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Increment",
			Handler:    _DotnetService_Increment_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "mastermodel.proto",
}
