// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.20.1
// source: khulnasoft/converter.proto

package khulnasoftrpc

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

// ConverterClient is the client API for Converter service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ConverterClient interface {
	// ConvertState converts state from the target ecosystem into a form that can be imported into Pulumi.
	ConvertState(ctx context.Context, in *ConvertStateRequest, opts ...grpc.CallOption) (*ConvertStateResponse, error)
	// ConvertProgram converts a program from the target ecosystem into a form that can be used with Pulumi.
	ConvertProgram(ctx context.Context, in *ConvertProgramRequest, opts ...grpc.CallOption) (*ConvertProgramResponse, error)
}

type converterClient struct {
	cc grpc.ClientConnInterface
}

func NewConverterClient(cc grpc.ClientConnInterface) ConverterClient {
	return &converterClient{cc}
}

func (c *converterClient) ConvertState(ctx context.Context, in *ConvertStateRequest, opts ...grpc.CallOption) (*ConvertStateResponse, error) {
	out := new(ConvertStateResponse)
	err := c.cc.Invoke(ctx, "/khulnasoftrpc.Converter/ConvertState", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *converterClient) ConvertProgram(ctx context.Context, in *ConvertProgramRequest, opts ...grpc.CallOption) (*ConvertProgramResponse, error) {
	out := new(ConvertProgramResponse)
	err := c.cc.Invoke(ctx, "/khulnasoftrpc.Converter/ConvertProgram", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ConverterServer is the server API for Converter service.
// All implementations must embed UnimplementedConverterServer
// for forward compatibility
type ConverterServer interface {
	// ConvertState converts state from the target ecosystem into a form that can be imported into Pulumi.
	ConvertState(context.Context, *ConvertStateRequest) (*ConvertStateResponse, error)
	// ConvertProgram converts a program from the target ecosystem into a form that can be used with Pulumi.
	ConvertProgram(context.Context, *ConvertProgramRequest) (*ConvertProgramResponse, error)
	mustEmbedUnimplementedConverterServer()
}

// UnimplementedConverterServer must be embedded to have forward compatible implementations.
type UnimplementedConverterServer struct {
}

func (UnimplementedConverterServer) ConvertState(context.Context, *ConvertStateRequest) (*ConvertStateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ConvertState not implemented")
}
func (UnimplementedConverterServer) ConvertProgram(context.Context, *ConvertProgramRequest) (*ConvertProgramResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ConvertProgram not implemented")
}
func (UnimplementedConverterServer) mustEmbedUnimplementedConverterServer() {}

// UnsafeConverterServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ConverterServer will
// result in compilation errors.
type UnsafeConverterServer interface {
	mustEmbedUnimplementedConverterServer()
}

func RegisterConverterServer(s grpc.ServiceRegistrar, srv ConverterServer) {
	s.RegisterService(&Converter_ServiceDesc, srv)
}

func _Converter_ConvertState_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConvertStateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConverterServer).ConvertState(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/khulnasoftrpc.Converter/ConvertState",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConverterServer).ConvertState(ctx, req.(*ConvertStateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Converter_ConvertProgram_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConvertProgramRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConverterServer).ConvertProgram(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/khulnasoftrpc.Converter/ConvertProgram",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConverterServer).ConvertProgram(ctx, req.(*ConvertProgramRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Converter_ServiceDesc is the grpc.ServiceDesc for Converter service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Converter_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "khulnasoftrpc.Converter",
	HandlerType: (*ConverterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ConvertState",
			Handler:    _Converter_ConvertState_Handler,
		},
		{
			MethodName: "ConvertProgram",
			Handler:    _Converter_ConvertProgram_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "khulnasoft/converter.proto",
}
