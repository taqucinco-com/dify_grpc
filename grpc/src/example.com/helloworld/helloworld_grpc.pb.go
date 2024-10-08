// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: helloworld/helloworld.proto

package helloworld

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
	SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloResponse, error)
	SayHelloAgain(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (Greeter_SayHelloAgainClient, error)
	SayHelloToMany(ctx context.Context, opts ...grpc.CallOption) (Greeter_SayHelloToManyClient, error)
	SayChat(ctx context.Context, opts ...grpc.CallOption) (Greeter_SayChatClient, error)
}

type greeterClient struct {
	cc grpc.ClientConnInterface
}

func NewGreeterClient(cc grpc.ClientConnInterface) GreeterClient {
	return &greeterClient{cc}
}

func (c *greeterClient) SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloResponse, error) {
	out := new(HelloResponse)
	err := c.cc.Invoke(ctx, "/helloworld.Greeter/SayHello", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *greeterClient) SayHelloAgain(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (Greeter_SayHelloAgainClient, error) {
	stream, err := c.cc.NewStream(ctx, &Greeter_ServiceDesc.Streams[0], "/helloworld.Greeter/SayHelloAgain", opts...)
	if err != nil {
		return nil, err
	}
	x := &greeterSayHelloAgainClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Greeter_SayHelloAgainClient interface {
	Recv() (*HelloResponse, error)
	grpc.ClientStream
}

type greeterSayHelloAgainClient struct {
	grpc.ClientStream
}

func (x *greeterSayHelloAgainClient) Recv() (*HelloResponse, error) {
	m := new(HelloResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *greeterClient) SayHelloToMany(ctx context.Context, opts ...grpc.CallOption) (Greeter_SayHelloToManyClient, error) {
	stream, err := c.cc.NewStream(ctx, &Greeter_ServiceDesc.Streams[1], "/helloworld.Greeter/SayHelloToMany", opts...)
	if err != nil {
		return nil, err
	}
	x := &greeterSayHelloToManyClient{stream}
	return x, nil
}

type Greeter_SayHelloToManyClient interface {
	Send(*HelloRequest) error
	CloseAndRecv() (*HelloResponse, error)
	grpc.ClientStream
}

type greeterSayHelloToManyClient struct {
	grpc.ClientStream
}

func (x *greeterSayHelloToManyClient) Send(m *HelloRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *greeterSayHelloToManyClient) CloseAndRecv() (*HelloResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(HelloResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *greeterClient) SayChat(ctx context.Context, opts ...grpc.CallOption) (Greeter_SayChatClient, error) {
	stream, err := c.cc.NewStream(ctx, &Greeter_ServiceDesc.Streams[2], "/helloworld.Greeter/SayChat", opts...)
	if err != nil {
		return nil, err
	}
	x := &greeterSayChatClient{stream}
	return x, nil
}

type Greeter_SayChatClient interface {
	Send(*HelloRequest) error
	Recv() (*HelloResponse, error)
	grpc.ClientStream
}

type greeterSayChatClient struct {
	grpc.ClientStream
}

func (x *greeterSayChatClient) Send(m *HelloRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *greeterSayChatClient) Recv() (*HelloResponse, error) {
	m := new(HelloResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// GreeterServer is the server API for Greeter service.
// All implementations must embed UnimplementedGreeterServer
// for forward compatibility
type GreeterServer interface {
	SayHello(context.Context, *HelloRequest) (*HelloResponse, error)
	SayHelloAgain(*HelloRequest, Greeter_SayHelloAgainServer) error
	SayHelloToMany(Greeter_SayHelloToManyServer) error
	SayChat(Greeter_SayChatServer) error
	mustEmbedUnimplementedGreeterServer()
}

// UnimplementedGreeterServer must be embedded to have forward compatible implementations.
type UnimplementedGreeterServer struct {
}

func (UnimplementedGreeterServer) SayHello(context.Context, *HelloRequest) (*HelloResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SayHello not implemented")
}
func (UnimplementedGreeterServer) SayHelloAgain(*HelloRequest, Greeter_SayHelloAgainServer) error {
	return status.Errorf(codes.Unimplemented, "method SayHelloAgain not implemented")
}
func (UnimplementedGreeterServer) SayHelloToMany(Greeter_SayHelloToManyServer) error {
	return status.Errorf(codes.Unimplemented, "method SayHelloToMany not implemented")
}
func (UnimplementedGreeterServer) SayChat(Greeter_SayChatServer) error {
	return status.Errorf(codes.Unimplemented, "method SayChat not implemented")
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
		FullMethod: "/helloworld.Greeter/SayHello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreeterServer).SayHello(ctx, req.(*HelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Greeter_SayHelloAgain_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(HelloRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(GreeterServer).SayHelloAgain(m, &greeterSayHelloAgainServer{stream})
}

type Greeter_SayHelloAgainServer interface {
	Send(*HelloResponse) error
	grpc.ServerStream
}

type greeterSayHelloAgainServer struct {
	grpc.ServerStream
}

func (x *greeterSayHelloAgainServer) Send(m *HelloResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _Greeter_SayHelloToMany_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(GreeterServer).SayHelloToMany(&greeterSayHelloToManyServer{stream})
}

type Greeter_SayHelloToManyServer interface {
	SendAndClose(*HelloResponse) error
	Recv() (*HelloRequest, error)
	grpc.ServerStream
}

type greeterSayHelloToManyServer struct {
	grpc.ServerStream
}

func (x *greeterSayHelloToManyServer) SendAndClose(m *HelloResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *greeterSayHelloToManyServer) Recv() (*HelloRequest, error) {
	m := new(HelloRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Greeter_SayChat_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(GreeterServer).SayChat(&greeterSayChatServer{stream})
}

type Greeter_SayChatServer interface {
	Send(*HelloResponse) error
	Recv() (*HelloRequest, error)
	grpc.ServerStream
}

type greeterSayChatServer struct {
	grpc.ServerStream
}

func (x *greeterSayChatServer) Send(m *HelloResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *greeterSayChatServer) Recv() (*HelloRequest, error) {
	m := new(HelloRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Greeter_ServiceDesc is the grpc.ServiceDesc for Greeter service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Greeter_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "helloworld.Greeter",
	HandlerType: (*GreeterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SayHello",
			Handler:    _Greeter_SayHello_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "SayHelloAgain",
			Handler:       _Greeter_SayHelloAgain_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "SayHelloToMany",
			Handler:       _Greeter_SayHelloToMany_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "SayChat",
			Handler:       _Greeter_SayChat_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "helloworld/helloworld.proto",
}
