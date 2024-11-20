// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.20.1
// source: khulnasoft/codegen/mapper.proto

package codegen

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

// MapperClient is the client API for Mapper service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MapperClient interface {
	// GetMapping tries to find a mapping for the given provider.
	GetMapping(ctx context.Context, in *GetMappingRequest, opts ...grpc.CallOption) (*GetMappingResponse, error)
}

type mapperClient struct {
	cc grpc.ClientConnInterface
}

func NewMapperClient(cc grpc.ClientConnInterface) MapperClient {
	return &mapperClient{cc}
}

func (c *mapperClient) GetMapping(ctx context.Context, in *GetMappingRequest, opts ...grpc.CallOption) (*GetMappingResponse, error) {
	out := new(GetMappingResponse)
	err := c.cc.Invoke(ctx, "/codegen.Mapper/GetMapping", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MapperServer is the server API for Mapper service.
// All implementations must embed UnimplementedMapperServer
// for forward compatibility
type MapperServer interface {
	// GetMapping tries to find a mapping for the given provider.
	GetMapping(context.Context, *GetMappingRequest) (*GetMappingResponse, error)
	mustEmbedUnimplementedMapperServer()
}

// UnimplementedMapperServer must be embedded to have forward compatible implementations.
type UnimplementedMapperServer struct {
}

func (UnimplementedMapperServer) GetMapping(context.Context, *GetMappingRequest) (*GetMappingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMapping not implemented")
}
func (UnimplementedMapperServer) mustEmbedUnimplementedMapperServer() {}

// UnsafeMapperServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MapperServer will
// result in compilation errors.
type UnsafeMapperServer interface {
	mustEmbedUnimplementedMapperServer()
}

func RegisterMapperServer(s grpc.ServiceRegistrar, srv MapperServer) {
	s.RegisterService(&Mapper_ServiceDesc, srv)
}

func _Mapper_GetMapping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMappingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MapperServer).GetMapping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/codegen.Mapper/GetMapping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MapperServer).GetMapping(ctx, req.(*GetMappingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Mapper_ServiceDesc is the grpc.ServiceDesc for Mapper service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Mapper_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "codegen.Mapper",
	HandlerType: (*MapperServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetMapping",
			Handler:    _Mapper_GetMapping_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "khulnasoft/codegen/mapper.proto",
}
