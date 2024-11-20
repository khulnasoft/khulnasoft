// Copyright 2016-2023, Pulumi Corporation.  All rights reserved.
//go:build !all
// +build !all

package main

import (
	"context"
	"fmt"
	"strconv"
	"sync/atomic"

	"github.com/khulnasoft/khulnasoft/pkg/v3/resource/provider"
	"github.com/khulnasoft/khulnasoft/sdk/v3/go/common/resource"
	"github.com/khulnasoft/khulnasoft/sdk/v3/go/common/util/cmdutil"
	"github.com/khulnasoft/khulnasoft/sdk/v3/go/khulnasoft"
	khulnasoftprovider "github.com/khulnasoft/khulnasoft/sdk/v3/go/khulnasoft/provider"
	khulnasoftrpc "github.com/khulnasoft/khulnasoft/sdk/v3/proto/go"
	"google.golang.org/protobuf/types/known/emptypb"
)

type Component struct {
	khulnasoft.ResourceState
}

func NewComponent(ctx *khulnasoft.Context, name string, opts ...khulnasoft.ResourceOption) (*Component, error) {
	var comp Component
	if err := ctx.RegisterComponentResource("testcomponent:index:Component", name, &comp, opts...); err != nil {
		return nil, err
	}

	return &comp, nil
}

func main() {
	err := provider.Main("testcomponent", func(host *provider.HostClient) (khulnasoftrpc.ResourceProviderServer, error) {
		return NewProvider(host, "testcomponent", "0.0.1"), nil
	})
	if err != nil {
		cmdutil.Exit(err)
	}
}

type Provider struct {
	khulnasoftrpc.UnimplementedResourceProviderServer

	host    *provider.HostClient
	name    string
	version string

	// The next ID to use for a resource.
	currentID atomic.Int64
}

func NewProvider(host *provider.HostClient, name, version string) khulnasoftrpc.ResourceProviderServer {
	return &Provider{
		host:    host,
		name:    name,
		version: version,
	}
}

func (p *Provider) Create(ctx context.Context,
	req *khulnasoftrpc.CreateRequest,
) (*khulnasoftrpc.CreateResponse, error) {
	urn := resource.URN(req.GetUrn())
	typ := urn.Type()
	if typ != "testcomponent:index:Resource" {
		return nil, fmt.Errorf("unknown resource type %q", typ)
	}

	id := p.currentID.Add(1)
	return &khulnasoftrpc.CreateResponse{
		Id: strconv.FormatInt(id, 10),
	}, nil
}

func (p *Provider) Construct(ctx context.Context,
	req *khulnasoftrpc.ConstructRequest,
) (*khulnasoftrpc.ConstructResponse, error) {
	return khulnasoftprovider.Construct(ctx, req, p.host.EngineConn(), func(
		ctx *khulnasoft.Context,
		typ, name string,
		inputs khulnasoftprovider.ConstructInputs,
		options khulnasoft.ResourceOption,
	) (*khulnasoftprovider.ConstructResult, error) {
		if typ != "testcomponent:index:Component" {
			return nil, fmt.Errorf("unknown resource type %q", typ)
		}

		comp, err := NewComponent(ctx, name, options)
		if err != nil {
			return nil, err
		}

		return khulnasoftprovider.NewConstructResult(comp)
	})
}

func (p *Provider) CheckConfig(ctx context.Context,
	req *khulnasoftrpc.CheckRequest,
) (*khulnasoftrpc.CheckResponse, error) {
	return &khulnasoftrpc.CheckResponse{Inputs: req.GetNews()}, nil
}

func (p *Provider) DiffConfig(ctx context.Context,
	req *khulnasoftrpc.DiffRequest,
) (*khulnasoftrpc.DiffResponse, error) {
	return &khulnasoftrpc.DiffResponse{}, nil
}

func (p *Provider) Configure(ctx context.Context,
	req *khulnasoftrpc.ConfigureRequest,
) (*khulnasoftrpc.ConfigureResponse, error) {
	panic("great sadness")
}

func (p *Provider) Check(ctx context.Context,
	req *khulnasoftrpc.CheckRequest,
) (*khulnasoftrpc.CheckResponse, error) {
	return &khulnasoftrpc.CheckResponse{Inputs: req.News, Failures: nil}, nil
}

func (p *Provider) Diff(ctx context.Context, req *khulnasoftrpc.DiffRequest) (*khulnasoftrpc.DiffResponse, error) {
	return &khulnasoftrpc.DiffResponse{}, nil
}

func (p *Provider) Read(ctx context.Context, req *khulnasoftrpc.ReadRequest) (*khulnasoftrpc.ReadResponse, error) {
	return &khulnasoftrpc.ReadResponse{
		Id:         req.GetId(),
		Properties: req.GetProperties(),
	}, nil
}

func (p *Provider) Update(ctx context.Context,
	req *khulnasoftrpc.UpdateRequest,
) (*khulnasoftrpc.UpdateResponse, error) {
	return &khulnasoftrpc.UpdateResponse{
		Properties: req.GetNews(),
	}, nil
}

func (p *Provider) Delete(ctx context.Context, req *khulnasoftrpc.DeleteRequest) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, nil
}

func (p *Provider) GetPluginInfo(context.Context, *emptypb.Empty) (*khulnasoftrpc.PluginInfo, error) {
	return &khulnasoftrpc.PluginInfo{
		Version: p.version,
	}, nil
}

func (p *Provider) Attach(ctx context.Context, req *khulnasoftrpc.PluginAttach) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, nil
}

func (p *Provider) GetSchema(ctx context.Context,
	req *khulnasoftrpc.GetSchemaRequest,
) (*khulnasoftrpc.GetSchemaResponse, error) {
	return &khulnasoftrpc.GetSchemaResponse{}, nil
}

func (p *Provider) Cancel(context.Context, *emptypb.Empty) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, nil
}

func (p *Provider) GetMapping(context.Context, *khulnasoftrpc.GetMappingRequest) (*khulnasoftrpc.GetMappingResponse, error) {
	return &khulnasoftrpc.GetMappingResponse{}, nil
}

func (p *Provider) Invoke(ctx context.Context,
	req *khulnasoftrpc.InvokeRequest,
) (*khulnasoftrpc.InvokeResponse, error) {
	return nil, fmt.Errorf("unknown Invoke %q", req.GetTok())
}

func (p *Provider) StreamInvoke(req *khulnasoftrpc.InvokeRequest,
	server khulnasoftrpc.ResourceProvider_StreamInvokeServer,
) error {
	return fmt.Errorf("unknown StreamInvoke %q", req.GetTok())
}

func (p *Provider) Call(ctx context.Context,
	req *khulnasoftrpc.CallRequest,
) (*khulnasoftrpc.CallResponse, error) {
	return nil, fmt.Errorf("unknown Call %q", req.GetTok())
}
