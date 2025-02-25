// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: slinky/service/v1/service.proto

package servicev1

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

const (
	Oracle_Prices_FullMethodName = "/slinky.service.v1.Oracle/Prices"
)

// OracleClient is the client API for Oracle service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OracleClient interface {
	// Prices defines a method for fetching the latest prices.
	Prices(ctx context.Context, in *QueryPricesRequest, opts ...grpc.CallOption) (*QueryPricesResponse, error)
}

type oracleClient struct {
	cc grpc.ClientConnInterface
}

func NewOracleClient(cc grpc.ClientConnInterface) OracleClient {
	return &oracleClient{cc}
}

func (c *oracleClient) Prices(ctx context.Context, in *QueryPricesRequest, opts ...grpc.CallOption) (*QueryPricesResponse, error) {
	out := new(QueryPricesResponse)
	err := c.cc.Invoke(ctx, Oracle_Prices_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OracleServer is the server API for Oracle service.
// All implementations must embed UnimplementedOracleServer
// for forward compatibility
type OracleServer interface {
	// Prices defines a method for fetching the latest prices.
	Prices(context.Context, *QueryPricesRequest) (*QueryPricesResponse, error)
	mustEmbedUnimplementedOracleServer()
}

// UnimplementedOracleServer must be embedded to have forward compatible implementations.
type UnimplementedOracleServer struct {
}

func (UnimplementedOracleServer) Prices(context.Context, *QueryPricesRequest) (*QueryPricesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Prices not implemented")
}
func (UnimplementedOracleServer) mustEmbedUnimplementedOracleServer() {}

// UnsafeOracleServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OracleServer will
// result in compilation errors.
type UnsafeOracleServer interface {
	mustEmbedUnimplementedOracleServer()
}

func RegisterOracleServer(s grpc.ServiceRegistrar, srv OracleServer) {
	s.RegisterService(&Oracle_ServiceDesc, srv)
}

func _Oracle_Prices_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryPricesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OracleServer).Prices(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Oracle_Prices_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OracleServer).Prices(ctx, req.(*QueryPricesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Oracle_ServiceDesc is the grpc.ServiceDesc for Oracle service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Oracle_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "slinky.service.v1.Oracle",
	HandlerType: (*OracleServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Prices",
			Handler:    _Oracle_Prices_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "slinky/service/v1/service.proto",
}
