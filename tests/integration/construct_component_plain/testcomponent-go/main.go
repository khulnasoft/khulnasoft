// Copyright 2016-2021, Pulumi Corporation.  All rights reserved.
//go:build !all
// +build !all

package main

import (
	"context"
	"errors"
	"fmt"

	"github.com/khulnasoft/khulnasoft/pkg/v3/resource/provider"
	"github.com/khulnasoft/khulnasoft/sdk/v3/go/common/resource"
	"github.com/khulnasoft/khulnasoft/sdk/v3/go/common/util/cmdutil"
	"github.com/khulnasoft/khulnasoft/sdk/v3/go/khulnasoft"
	khulnasoftprovider "github.com/khulnasoft/khulnasoft/sdk/v3/go/khulnasoft/provider"
	khulnasoftrpc "github.com/khulnasoft/khulnasoft/sdk/v3/proto/go"

	"google.golang.org/protobuf/types/known/emptypb"
)

type Resource struct {
	khulnasoft.CustomResourceState
}

func NewResource(ctx *khulnasoft.Context, name string, opts ...khulnasoft.ResourceOption) (*Resource, error) {
	var resource Resource
	if err := ctx.RegisterResource("testcomponent:index:Resource", name, nil, &resource, opts...); err != nil {
		return nil, err
	}
	return &resource, nil
}

type Component struct {
	khulnasoft.ResourceState
}

type ComponentArgs struct {
	Children int `khulnasoft:"children"`
}

func NewComponent(ctx *khulnasoft.Context, name string, args *ComponentArgs,
	opts ...khulnasoft.ResourceOption,
) (*Component, error) {
	if args == nil {
		return nil, errors.New("args is required")
	}

	component := &Component{}
	err := ctx.RegisterComponentResource("testcomponent:index:Component", name, component, opts...)
	if err != nil {
		return nil, err
	}

	if args.Children > 0 {
		for i := 0; i < args.Children; i++ {
			_, err := NewResource(ctx, fmt.Sprintf("child-%s-%v", name, i+1), khulnasoft.Parent(component))
			if err != nil {
				return nil, err
			}
		}
	}

	if err := ctx.RegisterResourceOutputs(component, khulnasoft.Map{}); err != nil {
		return nil, err
	}

	return component, nil
}

const (
	providerName = "testcomponent"
	version      = "0.0.1"
)

var currentID int

func main() {
	err := provider.Main(providerName, func(host *provider.HostClient) (khulnasoftrpc.ResourceProviderServer, error) {
		return makeProvider(host, providerName, version)
	})
	if err != nil {
		cmdutil.ExitError(err.Error())
	}
}

type testcomponentProvider struct {
	khulnasoftrpc.UnimplementedResourceProviderServer

	host    *provider.HostClient
	name    string
	version string
}

func makeProvider(host *provider.HostClient, name, version string) (khulnasoftrpc.ResourceProviderServer, error) {
	return &testcomponentProvider{
		host:    host,
		name:    name,
		version: version,
	}, nil
}

func (p *testcomponentProvider) Create(ctx context.Context,
	req *khulnasoftrpc.CreateRequest,
) (*khulnasoftrpc.CreateResponse, error) {
	urn := resource.URN(req.GetUrn())
	typ := urn.Type()
	if typ != "testcomponent:index:Resource" {
		return nil, fmt.Errorf("Unknown resource type '%s'", typ)
	}

	id := currentID
	currentID++

	return &khulnasoftrpc.CreateResponse{
		Id: fmt.Sprintf("%v", id),
	}, nil
}

func (p *testcomponentProvider) Construct(ctx context.Context,
	req *khulnasoftrpc.ConstructRequest,
) (*khulnasoftrpc.ConstructResponse, error) {
	return khulnasoftprovider.Construct(ctx, req, p.host.EngineConn(), func(ctx *khulnasoft.Context, typ, name string,
		inputs khulnasoftprovider.ConstructInputs, options khulnasoft.ResourceOption,
	) (*khulnasoftprovider.ConstructResult, error) {
		if typ != "testcomponent:index:Component" {
			return nil, fmt.Errorf("unknown resource type %s", typ)
		}

		args := &ComponentArgs{}
		if err := inputs.CopyTo(args); err != nil {
			return nil, fmt.Errorf("setting args: %w", err)
		}

		component, err := NewComponent(ctx, name, args, options)
		if err != nil {
			return nil, fmt.Errorf("creating component: %w", err)
		}

		return khulnasoftprovider.NewConstructResult(component)
	})
}

func (p *testcomponentProvider) CheckConfig(ctx context.Context,
	req *khulnasoftrpc.CheckRequest,
) (*khulnasoftrpc.CheckResponse, error) {
	return &khulnasoftrpc.CheckResponse{Inputs: req.GetNews()}, nil
}

func (p *testcomponentProvider) DiffConfig(ctx context.Context,
	req *khulnasoftrpc.DiffRequest,
) (*khulnasoftrpc.DiffResponse, error) {
	return &khulnasoftrpc.DiffResponse{}, nil
}

func (p *testcomponentProvider) Configure(ctx context.Context,
	req *khulnasoftrpc.ConfigureRequest,
) (*khulnasoftrpc.ConfigureResponse, error) {
	return &khulnasoftrpc.ConfigureResponse{
		AcceptSecrets:   true,
		SupportsPreview: true,
		AcceptResources: true,
	}, nil
}

func (p *testcomponentProvider) Invoke(ctx context.Context,
	req *khulnasoftrpc.InvokeRequest,
) (*khulnasoftrpc.InvokeResponse, error) {
	return nil, fmt.Errorf("Unknown Invoke token '%s'", req.GetTok())
}

func (p *testcomponentProvider) StreamInvoke(req *khulnasoftrpc.InvokeRequest,
	server khulnasoftrpc.ResourceProvider_StreamInvokeServer,
) error {
	return fmt.Errorf("Unknown StreamInvoke token '%s'", req.GetTok())
}

func (p *testcomponentProvider) Call(ctx context.Context,
	req *khulnasoftrpc.CallRequest,
) (*khulnasoftrpc.CallResponse, error) {
	return nil, fmt.Errorf("Unknown Call token '%s'", req.GetTok())
}

func (p *testcomponentProvider) Check(ctx context.Context,
	req *khulnasoftrpc.CheckRequest,
) (*khulnasoftrpc.CheckResponse, error) {
	return &khulnasoftrpc.CheckResponse{Inputs: req.News, Failures: nil}, nil
}

func (p *testcomponentProvider) Diff(ctx context.Context, req *khulnasoftrpc.DiffRequest) (*khulnasoftrpc.DiffResponse, error) {
	return &khulnasoftrpc.DiffResponse{}, nil
}

func (p *testcomponentProvider) Read(ctx context.Context, req *khulnasoftrpc.ReadRequest) (*khulnasoftrpc.ReadResponse, error) {
	return &khulnasoftrpc.ReadResponse{
		Id:         req.GetId(),
		Properties: req.GetProperties(),
	}, nil
}

func (p *testcomponentProvider) Update(ctx context.Context,
	req *khulnasoftrpc.UpdateRequest,
) (*khulnasoftrpc.UpdateResponse, error) {
	return &khulnasoftrpc.UpdateResponse{
		Properties: req.GetNews(),
	}, nil
}

func (p *testcomponentProvider) Delete(ctx context.Context, req *khulnasoftrpc.DeleteRequest) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, nil
}

func (p *testcomponentProvider) GetPluginInfo(context.Context, *emptypb.Empty) (*khulnasoftrpc.PluginInfo, error) {
	return &khulnasoftrpc.PluginInfo{
		Version: p.version,
	}, nil
}

func (p *testcomponentProvider) Attach(ctx context.Context, req *khulnasoftrpc.PluginAttach) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, nil
}

func (p *testcomponentProvider) GetSchema(ctx context.Context,
	req *khulnasoftrpc.GetSchemaRequest,
) (*khulnasoftrpc.GetSchemaResponse, error) {
	return &khulnasoftrpc.GetSchemaResponse{}, nil
}

func (p *testcomponentProvider) Cancel(context.Context, *emptypb.Empty) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, nil
}

func (p *testcomponentProvider) GetMapping(
	context.Context, *khulnasoftrpc.GetMappingRequest,
) (*khulnasoftrpc.GetMappingResponse, error) {
	return &khulnasoftrpc.GetMappingResponse{}, nil
}
