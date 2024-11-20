// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.20.1
// source: khulnasoft/language.proto

package khulnasoftrpc

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// LanguageRuntimeClient is the client API for LanguageRuntime service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LanguageRuntimeClient interface {
	// GetRequiredPlugins computes the complete set of anticipated plugins required by a program.
	GetRequiredPlugins(ctx context.Context, in *GetRequiredPluginsRequest, opts ...grpc.CallOption) (*GetRequiredPluginsResponse, error)
	// Run executes a program and returns its result.
	Run(ctx context.Context, in *RunRequest, opts ...grpc.CallOption) (*RunResponse, error)
	// GetPluginInfo returns generic information about this plugin, like its version.
	GetPluginInfo(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*PluginInfo, error)
	// InstallDependencies will install dependencies for the project, e.g. by running `npm install` for nodejs projects.
	InstallDependencies(ctx context.Context, in *InstallDependenciesRequest, opts ...grpc.CallOption) (LanguageRuntime_InstallDependenciesClient, error)
	// RuntimeOptionsPrompts returns a list of additional prompts to ask during `khulnasoft new`.
	RuntimeOptionsPrompts(ctx context.Context, in *RuntimeOptionsRequest, opts ...grpc.CallOption) (*RuntimeOptionsResponse, error)
	// About returns information about the runtime for this language.
	About(ctx context.Context, in *AboutRequest, opts ...grpc.CallOption) (*AboutResponse, error)
	// GetProgramDependencies returns the set of dependencies required by the program.
	GetProgramDependencies(ctx context.Context, in *GetProgramDependenciesRequest, opts ...grpc.CallOption) (*GetProgramDependenciesResponse, error)
	// RunPlugin executes a plugin program and returns its result asynchronously.
	RunPlugin(ctx context.Context, in *RunPluginRequest, opts ...grpc.CallOption) (LanguageRuntime_RunPluginClient, error)
	// GenerateProgram generates a given PCL program into a program for this language.
	GenerateProgram(ctx context.Context, in *GenerateProgramRequest, opts ...grpc.CallOption) (*GenerateProgramResponse, error)
	// GenerateProject generates a given PCL program into a project for this language.
	GenerateProject(ctx context.Context, in *GenerateProjectRequest, opts ...grpc.CallOption) (*GenerateProjectResponse, error)
	// GeneratePackage generates a given khulnasoft package into a package for this language.
	GeneratePackage(ctx context.Context, in *GeneratePackageRequest, opts ...grpc.CallOption) (*GeneratePackageResponse, error)
	// Pack packs a package into a language specific artifact.
	Pack(ctx context.Context, in *PackRequest, opts ...grpc.CallOption) (*PackResponse, error)
}

type languageRuntimeClient struct {
	cc grpc.ClientConnInterface
}

func NewLanguageRuntimeClient(cc grpc.ClientConnInterface) LanguageRuntimeClient {
	return &languageRuntimeClient{cc}
}

func (c *languageRuntimeClient) GetRequiredPlugins(ctx context.Context, in *GetRequiredPluginsRequest, opts ...grpc.CallOption) (*GetRequiredPluginsResponse, error) {
	out := new(GetRequiredPluginsResponse)
	err := c.cc.Invoke(ctx, "/khulnasoftrpc.LanguageRuntime/GetRequiredPlugins", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *languageRuntimeClient) Run(ctx context.Context, in *RunRequest, opts ...grpc.CallOption) (*RunResponse, error) {
	out := new(RunResponse)
	err := c.cc.Invoke(ctx, "/khulnasoftrpc.LanguageRuntime/Run", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *languageRuntimeClient) GetPluginInfo(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*PluginInfo, error) {
	out := new(PluginInfo)
	err := c.cc.Invoke(ctx, "/khulnasoftrpc.LanguageRuntime/GetPluginInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *languageRuntimeClient) InstallDependencies(ctx context.Context, in *InstallDependenciesRequest, opts ...grpc.CallOption) (LanguageRuntime_InstallDependenciesClient, error) {
	stream, err := c.cc.NewStream(ctx, &LanguageRuntime_ServiceDesc.Streams[0], "/khulnasoftrpc.LanguageRuntime/InstallDependencies", opts...)
	if err != nil {
		return nil, err
	}
	x := &languageRuntimeInstallDependenciesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type LanguageRuntime_InstallDependenciesClient interface {
	Recv() (*InstallDependenciesResponse, error)
	grpc.ClientStream
}

type languageRuntimeInstallDependenciesClient struct {
	grpc.ClientStream
}

func (x *languageRuntimeInstallDependenciesClient) Recv() (*InstallDependenciesResponse, error) {
	m := new(InstallDependenciesResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *languageRuntimeClient) RuntimeOptionsPrompts(ctx context.Context, in *RuntimeOptionsRequest, opts ...grpc.CallOption) (*RuntimeOptionsResponse, error) {
	out := new(RuntimeOptionsResponse)
	err := c.cc.Invoke(ctx, "/khulnasoftrpc.LanguageRuntime/RuntimeOptionsPrompts", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *languageRuntimeClient) About(ctx context.Context, in *AboutRequest, opts ...grpc.CallOption) (*AboutResponse, error) {
	out := new(AboutResponse)
	err := c.cc.Invoke(ctx, "/khulnasoftrpc.LanguageRuntime/About", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *languageRuntimeClient) GetProgramDependencies(ctx context.Context, in *GetProgramDependenciesRequest, opts ...grpc.CallOption) (*GetProgramDependenciesResponse, error) {
	out := new(GetProgramDependenciesResponse)
	err := c.cc.Invoke(ctx, "/khulnasoftrpc.LanguageRuntime/GetProgramDependencies", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *languageRuntimeClient) RunPlugin(ctx context.Context, in *RunPluginRequest, opts ...grpc.CallOption) (LanguageRuntime_RunPluginClient, error) {
	stream, err := c.cc.NewStream(ctx, &LanguageRuntime_ServiceDesc.Streams[1], "/khulnasoftrpc.LanguageRuntime/RunPlugin", opts...)
	if err != nil {
		return nil, err
	}
	x := &languageRuntimeRunPluginClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type LanguageRuntime_RunPluginClient interface {
	Recv() (*RunPluginResponse, error)
	grpc.ClientStream
}

type languageRuntimeRunPluginClient struct {
	grpc.ClientStream
}

func (x *languageRuntimeRunPluginClient) Recv() (*RunPluginResponse, error) {
	m := new(RunPluginResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *languageRuntimeClient) GenerateProgram(ctx context.Context, in *GenerateProgramRequest, opts ...grpc.CallOption) (*GenerateProgramResponse, error) {
	out := new(GenerateProgramResponse)
	err := c.cc.Invoke(ctx, "/khulnasoftrpc.LanguageRuntime/GenerateProgram", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *languageRuntimeClient) GenerateProject(ctx context.Context, in *GenerateProjectRequest, opts ...grpc.CallOption) (*GenerateProjectResponse, error) {
	out := new(GenerateProjectResponse)
	err := c.cc.Invoke(ctx, "/khulnasoftrpc.LanguageRuntime/GenerateProject", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *languageRuntimeClient) GeneratePackage(ctx context.Context, in *GeneratePackageRequest, opts ...grpc.CallOption) (*GeneratePackageResponse, error) {
	out := new(GeneratePackageResponse)
	err := c.cc.Invoke(ctx, "/khulnasoftrpc.LanguageRuntime/GeneratePackage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *languageRuntimeClient) Pack(ctx context.Context, in *PackRequest, opts ...grpc.CallOption) (*PackResponse, error) {
	out := new(PackResponse)
	err := c.cc.Invoke(ctx, "/khulnasoftrpc.LanguageRuntime/Pack", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LanguageRuntimeServer is the server API for LanguageRuntime service.
// All implementations must embed UnimplementedLanguageRuntimeServer
// for forward compatibility
type LanguageRuntimeServer interface {
	// GetRequiredPlugins computes the complete set of anticipated plugins required by a program.
	GetRequiredPlugins(context.Context, *GetRequiredPluginsRequest) (*GetRequiredPluginsResponse, error)
	// Run executes a program and returns its result.
	Run(context.Context, *RunRequest) (*RunResponse, error)
	// GetPluginInfo returns generic information about this plugin, like its version.
	GetPluginInfo(context.Context, *emptypb.Empty) (*PluginInfo, error)
	// InstallDependencies will install dependencies for the project, e.g. by running `npm install` for nodejs projects.
	InstallDependencies(*InstallDependenciesRequest, LanguageRuntime_InstallDependenciesServer) error
	// RuntimeOptionsPrompts returns a list of additional prompts to ask during `khulnasoft new`.
	RuntimeOptionsPrompts(context.Context, *RuntimeOptionsRequest) (*RuntimeOptionsResponse, error)
	// About returns information about the runtime for this language.
	About(context.Context, *AboutRequest) (*AboutResponse, error)
	// GetProgramDependencies returns the set of dependencies required by the program.
	GetProgramDependencies(context.Context, *GetProgramDependenciesRequest) (*GetProgramDependenciesResponse, error)
	// RunPlugin executes a plugin program and returns its result asynchronously.
	RunPlugin(*RunPluginRequest, LanguageRuntime_RunPluginServer) error
	// GenerateProgram generates a given PCL program into a program for this language.
	GenerateProgram(context.Context, *GenerateProgramRequest) (*GenerateProgramResponse, error)
	// GenerateProject generates a given PCL program into a project for this language.
	GenerateProject(context.Context, *GenerateProjectRequest) (*GenerateProjectResponse, error)
	// GeneratePackage generates a given khulnasoft package into a package for this language.
	GeneratePackage(context.Context, *GeneratePackageRequest) (*GeneratePackageResponse, error)
	// Pack packs a package into a language specific artifact.
	Pack(context.Context, *PackRequest) (*PackResponse, error)
	mustEmbedUnimplementedLanguageRuntimeServer()
}

// UnimplementedLanguageRuntimeServer must be embedded to have forward compatible implementations.
type UnimplementedLanguageRuntimeServer struct {
}

func (UnimplementedLanguageRuntimeServer) GetRequiredPlugins(context.Context, *GetRequiredPluginsRequest) (*GetRequiredPluginsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRequiredPlugins not implemented")
}
func (UnimplementedLanguageRuntimeServer) Run(context.Context, *RunRequest) (*RunResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Run not implemented")
}
func (UnimplementedLanguageRuntimeServer) GetPluginInfo(context.Context, *emptypb.Empty) (*PluginInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPluginInfo not implemented")
}
func (UnimplementedLanguageRuntimeServer) InstallDependencies(*InstallDependenciesRequest, LanguageRuntime_InstallDependenciesServer) error {
	return status.Errorf(codes.Unimplemented, "method InstallDependencies not implemented")
}
func (UnimplementedLanguageRuntimeServer) RuntimeOptionsPrompts(context.Context, *RuntimeOptionsRequest) (*RuntimeOptionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RuntimeOptionsPrompts not implemented")
}
func (UnimplementedLanguageRuntimeServer) About(context.Context, *AboutRequest) (*AboutResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method About not implemented")
}
func (UnimplementedLanguageRuntimeServer) GetProgramDependencies(context.Context, *GetProgramDependenciesRequest) (*GetProgramDependenciesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProgramDependencies not implemented")
}
func (UnimplementedLanguageRuntimeServer) RunPlugin(*RunPluginRequest, LanguageRuntime_RunPluginServer) error {
	return status.Errorf(codes.Unimplemented, "method RunPlugin not implemented")
}
func (UnimplementedLanguageRuntimeServer) GenerateProgram(context.Context, *GenerateProgramRequest) (*GenerateProgramResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GenerateProgram not implemented")
}
func (UnimplementedLanguageRuntimeServer) GenerateProject(context.Context, *GenerateProjectRequest) (*GenerateProjectResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GenerateProject not implemented")
}
func (UnimplementedLanguageRuntimeServer) GeneratePackage(context.Context, *GeneratePackageRequest) (*GeneratePackageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GeneratePackage not implemented")
}
func (UnimplementedLanguageRuntimeServer) Pack(context.Context, *PackRequest) (*PackResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Pack not implemented")
}
func (UnimplementedLanguageRuntimeServer) mustEmbedUnimplementedLanguageRuntimeServer() {}

// UnsafeLanguageRuntimeServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LanguageRuntimeServer will
// result in compilation errors.
type UnsafeLanguageRuntimeServer interface {
	mustEmbedUnimplementedLanguageRuntimeServer()
}

func RegisterLanguageRuntimeServer(s grpc.ServiceRegistrar, srv LanguageRuntimeServer) {
	s.RegisterService(&LanguageRuntime_ServiceDesc, srv)
}

func _LanguageRuntime_GetRequiredPlugins_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequiredPluginsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LanguageRuntimeServer).GetRequiredPlugins(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/khulnasoftrpc.LanguageRuntime/GetRequiredPlugins",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LanguageRuntimeServer).GetRequiredPlugins(ctx, req.(*GetRequiredPluginsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LanguageRuntime_Run_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RunRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LanguageRuntimeServer).Run(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/khulnasoftrpc.LanguageRuntime/Run",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LanguageRuntimeServer).Run(ctx, req.(*RunRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LanguageRuntime_GetPluginInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LanguageRuntimeServer).GetPluginInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/khulnasoftrpc.LanguageRuntime/GetPluginInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LanguageRuntimeServer).GetPluginInfo(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _LanguageRuntime_InstallDependencies_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(InstallDependenciesRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(LanguageRuntimeServer).InstallDependencies(m, &languageRuntimeInstallDependenciesServer{stream})
}

type LanguageRuntime_InstallDependenciesServer interface {
	Send(*InstallDependenciesResponse) error
	grpc.ServerStream
}

type languageRuntimeInstallDependenciesServer struct {
	grpc.ServerStream
}

func (x *languageRuntimeInstallDependenciesServer) Send(m *InstallDependenciesResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _LanguageRuntime_RuntimeOptionsPrompts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RuntimeOptionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LanguageRuntimeServer).RuntimeOptionsPrompts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/khulnasoftrpc.LanguageRuntime/RuntimeOptionsPrompts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LanguageRuntimeServer).RuntimeOptionsPrompts(ctx, req.(*RuntimeOptionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LanguageRuntime_About_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AboutRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LanguageRuntimeServer).About(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/khulnasoftrpc.LanguageRuntime/About",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LanguageRuntimeServer).About(ctx, req.(*AboutRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LanguageRuntime_GetProgramDependencies_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetProgramDependenciesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LanguageRuntimeServer).GetProgramDependencies(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/khulnasoftrpc.LanguageRuntime/GetProgramDependencies",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LanguageRuntimeServer).GetProgramDependencies(ctx, req.(*GetProgramDependenciesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LanguageRuntime_RunPlugin_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(RunPluginRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(LanguageRuntimeServer).RunPlugin(m, &languageRuntimeRunPluginServer{stream})
}

type LanguageRuntime_RunPluginServer interface {
	Send(*RunPluginResponse) error
	grpc.ServerStream
}

type languageRuntimeRunPluginServer struct {
	grpc.ServerStream
}

func (x *languageRuntimeRunPluginServer) Send(m *RunPluginResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _LanguageRuntime_GenerateProgram_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GenerateProgramRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LanguageRuntimeServer).GenerateProgram(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/khulnasoftrpc.LanguageRuntime/GenerateProgram",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LanguageRuntimeServer).GenerateProgram(ctx, req.(*GenerateProgramRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LanguageRuntime_GenerateProject_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GenerateProjectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LanguageRuntimeServer).GenerateProject(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/khulnasoftrpc.LanguageRuntime/GenerateProject",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LanguageRuntimeServer).GenerateProject(ctx, req.(*GenerateProjectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LanguageRuntime_GeneratePackage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GeneratePackageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LanguageRuntimeServer).GeneratePackage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/khulnasoftrpc.LanguageRuntime/GeneratePackage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LanguageRuntimeServer).GeneratePackage(ctx, req.(*GeneratePackageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LanguageRuntime_Pack_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PackRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LanguageRuntimeServer).Pack(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/khulnasoftrpc.LanguageRuntime/Pack",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LanguageRuntimeServer).Pack(ctx, req.(*PackRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// LanguageRuntime_ServiceDesc is the grpc.ServiceDesc for LanguageRuntime service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var LanguageRuntime_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "khulnasoftrpc.LanguageRuntime",
	HandlerType: (*LanguageRuntimeServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetRequiredPlugins",
			Handler:    _LanguageRuntime_GetRequiredPlugins_Handler,
		},
		{
			MethodName: "Run",
			Handler:    _LanguageRuntime_Run_Handler,
		},
		{
			MethodName: "GetPluginInfo",
			Handler:    _LanguageRuntime_GetPluginInfo_Handler,
		},
		{
			MethodName: "RuntimeOptionsPrompts",
			Handler:    _LanguageRuntime_RuntimeOptionsPrompts_Handler,
		},
		{
			MethodName: "About",
			Handler:    _LanguageRuntime_About_Handler,
		},
		{
			MethodName: "GetProgramDependencies",
			Handler:    _LanguageRuntime_GetProgramDependencies_Handler,
		},
		{
			MethodName: "GenerateProgram",
			Handler:    _LanguageRuntime_GenerateProgram_Handler,
		},
		{
			MethodName: "GenerateProject",
			Handler:    _LanguageRuntime_GenerateProject_Handler,
		},
		{
			MethodName: "GeneratePackage",
			Handler:    _LanguageRuntime_GeneratePackage_Handler,
		},
		{
			MethodName: "Pack",
			Handler:    _LanguageRuntime_Pack_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "InstallDependencies",
			Handler:       _LanguageRuntime_InstallDependencies_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "RunPlugin",
			Handler:       _LanguageRuntime_RunPlugin_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "khulnasoft/language.proto",
}
