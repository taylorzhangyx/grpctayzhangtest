// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.1
// source: tayzhangtest.proto

package taylorzhtestpb

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

// GreeterClient is the client API for Greeter service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GreeterClient interface {
	// Sends a greeting
	SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error)
	LoadTest(ctx context.Context, in *LoadTestRequest, opts ...grpc.CallOption) (*LoadTestReply, error)
	LoadTestMetrics(ctx context.Context, in *LoadTestMetricsRequest, opts ...grpc.CallOption) (*LoadTestMetricsReply, error)
	LoadTestClear(ctx context.Context, in *LoadTestClearRequest, opts ...grpc.CallOption) (*LoadTestClearReply, error)
}

type greeterClient struct {
	cc grpc.ClientConnInterface
}

func NewGreeterClient(cc grpc.ClientConnInterface) GreeterClient {
	return &greeterClient{cc}
}

func (c *greeterClient) SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error) {
	out := new(HelloReply)
	err := c.cc.Invoke(ctx, "/taylorzhtest.Greeter/SayHello", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *greeterClient) LoadTest(ctx context.Context, in *LoadTestRequest, opts ...grpc.CallOption) (*LoadTestReply, error) {
	out := new(LoadTestReply)
	err := c.cc.Invoke(ctx, "/taylorzhtest.Greeter/LoadTest", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *greeterClient) LoadTestMetrics(ctx context.Context, in *LoadTestMetricsRequest, opts ...grpc.CallOption) (*LoadTestMetricsReply, error) {
	out := new(LoadTestMetricsReply)
	err := c.cc.Invoke(ctx, "/taylorzhtest.Greeter/LoadTestMetrics", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *greeterClient) LoadTestClear(ctx context.Context, in *LoadTestClearRequest, opts ...grpc.CallOption) (*LoadTestClearReply, error) {
	out := new(LoadTestClearReply)
	err := c.cc.Invoke(ctx, "/taylorzhtest.Greeter/LoadTestClear", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GreeterServer is the server API for Greeter service.
// All implementations must embed UnimplementedGreeterServer
// for forward compatibility
type GreeterServer interface {
	// Sends a greeting
	SayHello(context.Context, *HelloRequest) (*HelloReply, error)
	LoadTest(context.Context, *LoadTestRequest) (*LoadTestReply, error)
	LoadTestMetrics(context.Context, *LoadTestMetricsRequest) (*LoadTestMetricsReply, error)
	LoadTestClear(context.Context, *LoadTestClearRequest) (*LoadTestClearReply, error)
	mustEmbedUnimplementedGreeterServer()
}

// UnimplementedGreeterServer must be embedded to have forward compatible implementations.
type UnimplementedGreeterServer struct {
}

func (UnimplementedGreeterServer) SayHello(context.Context, *HelloRequest) (*HelloReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SayHello not implemented")
}
func (UnimplementedGreeterServer) LoadTest(context.Context, *LoadTestRequest) (*LoadTestReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LoadTest not implemented")
}
func (UnimplementedGreeterServer) LoadTestMetrics(context.Context, *LoadTestMetricsRequest) (*LoadTestMetricsReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LoadTestMetrics not implemented")
}
func (UnimplementedGreeterServer) LoadTestClear(context.Context, *LoadTestClearRequest) (*LoadTestClearReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LoadTestClear not implemented")
}
func (UnimplementedGreeterServer) mustEmbedUnimplementedGreeterServer() {}

// UnsafeGreeterServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GreeterServer will
// result in compilation errors.
type UnsafeGreeterServer interface {
	mustEmbedUnimplementedGreeterServer()
}

func RegisterGreeterServer(s grpc.ServiceRegistrar, srv GreeterServer) {
	s.RegisterService(&Greeter_ServiceDesc, srv)
}

func _Greeter_SayHello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreeterServer).SayHello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/taylorzhtest.Greeter/SayHello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreeterServer).SayHello(ctx, req.(*HelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Greeter_LoadTest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoadTestRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreeterServer).LoadTest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/taylorzhtest.Greeter/LoadTest",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreeterServer).LoadTest(ctx, req.(*LoadTestRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Greeter_LoadTestMetrics_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoadTestMetricsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreeterServer).LoadTestMetrics(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/taylorzhtest.Greeter/LoadTestMetrics",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreeterServer).LoadTestMetrics(ctx, req.(*LoadTestMetricsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Greeter_LoadTestClear_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoadTestClearRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreeterServer).LoadTestClear(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/taylorzhtest.Greeter/LoadTestClear",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreeterServer).LoadTestClear(ctx, req.(*LoadTestClearRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Greeter_ServiceDesc is the grpc.ServiceDesc for Greeter service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Greeter_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "taylorzhtest.Greeter",
	HandlerType: (*GreeterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SayHello",
			Handler:    _Greeter_SayHello_Handler,
		},
		{
			MethodName: "LoadTest",
			Handler:    _Greeter_LoadTest_Handler,
		},
		{
			MethodName: "LoadTestMetrics",
			Handler:    _Greeter_LoadTestMetrics_Handler,
		},
		{
			MethodName: "LoadTestClear",
			Handler:    _Greeter_LoadTestClear_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "tayzhangtest.proto",
}