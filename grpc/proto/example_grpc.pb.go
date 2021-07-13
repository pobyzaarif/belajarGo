// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package proto

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

// CalculateClient is the client API for Calculate service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CalculateClient interface {
	Add(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
	Multiply(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
}

type calculateClient struct {
	cc grpc.ClientConnInterface
}

func NewCalculateClient(cc grpc.ClientConnInterface) CalculateClient {
	return &calculateClient{cc}
}

func (c *calculateClient) Add(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/proto.Calculate/Add", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *calculateClient) Multiply(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/proto.Calculate/Multiply", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CalculateServer is the server API for Calculate service.
// All implementations must embed UnimplementedCalculateServer
// for forward compatibility
type CalculateServer interface {
	Add(context.Context, *Request) (*Response, error)
	Multiply(context.Context, *Request) (*Response, error)
	mustEmbedUnimplementedCalculateServer()
}

// UnimplementedCalculateServer must be embedded to have forward compatible implementations.
type UnimplementedCalculateServer struct {
}

func (UnimplementedCalculateServer) Add(context.Context, *Request) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Add not implemented")
}
func (UnimplementedCalculateServer) Multiply(context.Context, *Request) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Multiply not implemented")
}
func (UnimplementedCalculateServer) mustEmbedUnimplementedCalculateServer() {}

// UnsafeCalculateServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CalculateServer will
// result in compilation errors.
type UnsafeCalculateServer interface {
	mustEmbedUnimplementedCalculateServer()
}

func RegisterCalculateServer(s grpc.ServiceRegistrar, srv CalculateServer) {
	s.RegisterService(&Calculate_ServiceDesc, srv)
}

func _Calculate_Add_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalculateServer).Add(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Calculate/Add",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalculateServer).Add(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _Calculate_Multiply_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalculateServer).Multiply(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Calculate/Multiply",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalculateServer).Multiply(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

// Calculate_ServiceDesc is the grpc.ServiceDesc for Calculate service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Calculate_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Calculate",
	HandlerType: (*CalculateServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Add",
			Handler:    _Calculate_Add_Handler,
		},
		{
			MethodName: "Multiply",
			Handler:    _Calculate_Multiply_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "example.proto",
}