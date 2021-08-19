// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package fetcher

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

// FetcherClient is the client API for Fetcher service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FetcherClient interface {
	Fetch(ctx context.Context, in *FetchRequest, opts ...grpc.CallOption) (*FetchReply, error)
}

type fetcherClient struct {
	cc grpc.ClientConnInterface
}

func NewFetcherClient(cc grpc.ClientConnInterface) FetcherClient {
	return &fetcherClient{cc}
}

func (c *fetcherClient) Fetch(ctx context.Context, in *FetchRequest, opts ...grpc.CallOption) (*FetchReply, error) {
	out := new(FetchReply)
	err := c.cc.Invoke(ctx, "/fetcher.Fetcher/Fetch", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FetcherServer is the server API for Fetcher service.
// All implementations must embed UnimplementedFetcherServer
// for forward compatibility
type FetcherServer interface {
	Fetch(context.Context, *FetchRequest) (*FetchReply, error)
	mustEmbedUnimplementedFetcherServer()
}

// UnimplementedFetcherServer must be embedded to have forward compatible implementations.
type UnimplementedFetcherServer struct {
}

func (UnimplementedFetcherServer) Fetch(context.Context, *FetchRequest) (*FetchReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Fetch not implemented")
}
func (UnimplementedFetcherServer) mustEmbedUnimplementedFetcherServer() {}

// UnsafeFetcherServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FetcherServer will
// result in compilation errors.
type UnsafeFetcherServer interface {
	mustEmbedUnimplementedFetcherServer()
}

func RegisterFetcherServer(s grpc.ServiceRegistrar, srv FetcherServer) {
	s.RegisterService(&Fetcher_ServiceDesc, srv)
}

func _Fetcher_Fetch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FetchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FetcherServer).Fetch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fetcher.Fetcher/Fetch",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FetcherServer).Fetch(ctx, req.(*FetchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Fetcher_ServiceDesc is the grpc.ServiceDesc for Fetcher service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Fetcher_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "fetcher.Fetcher",
	HandlerType: (*FetcherServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Fetch",
			Handler:    _Fetcher_Fetch_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "fetcher.proto",
}
