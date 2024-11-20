// Copyright 2016-2021, Pulumi Corporation.  All rights reserved.
//go:build !all
// +build !all

package main

import (
	"context"
	"fmt"

	"github.com/khulnasoft/khulnasoft/pkg/v3/resource/provider"
	"github.com/khulnasoft/khulnasoft/sdk/v3/go/common/util/cmdutil"
	"github.com/khulnasoft/khulnasoft/sdk/v3/go/khulnasoft"
	khulnasoftprovider "github.com/khulnasoft/khulnasoft/sdk/v3/go/khulnasoft/provider"
	khulnasoftrpc "github.com/khulnasoft/khulnasoft/sdk/v3/proto/go"

	"google.golang.org/protobuf/types/known/emptypb"
)

type Bucket struct {
	khulnasoft.CustomResourceState
}

func NewBucket(ctx *khulnasoft.Context, name string, opts ...khulnasoft.ResourceOption) (*Bucket, error) {
	resource := &Bucket{}
	err := ctx.RegisterResource(typeToken("Bucket"), name, nil, resource, opts...)
	if err != nil {
		return nil, err
	}
	return resource, nil
}

type BucketComponent struct {
	khulnasoft.ResourceState
}

func NewBucketComponent(ctx *khulnasoft.Context, name string, opts ...khulnasoft.ResourceOption) (*BucketComponent, error) {
	component := &BucketComponent{}
	err := ctx.RegisterComponentResource(typeToken("BucketComponent"), name, component, opts...)
	if err != nil {
		return nil, err
	}

	_, err = NewBucket(ctx, name+"child", khulnasoft.Parent(component))
	if err != nil {
		return nil, err
	}

	return component, nil
}

type BucketV2 struct {
	khulnasoft.CustomResourceState
}

func NewBucketV2(ctx *khulnasoft.Context, name string, opts ...khulnasoft.ResourceOption) (*BucketV2, error) {
	resource := &BucketV2{}
	aliases := khulnasoft.Aliases([]khulnasoft.Alias{
		{
			Type: khulnasoft.String(typeToken("Bucket")),
		},
	})
	opts = append(opts, aliases)
	err := ctx.RegisterResource(typeToken("BucketV2"), name, nil, resource, opts...)
	if err != nil {
		return nil, err
	}
	return resource, nil
}

type BucketComponentV2 struct {
	khulnasoft.ResourceState
}

func NewBucketComponentV2(ctx *khulnasoft.Context, name string, opts ...khulnasoft.ResourceOption) (*BucketComponentV2, error) {
	component := &BucketComponentV2{}
	aliases := khulnasoft.Aliases([]khulnasoft.Alias{
		{
			Type: khulnasoft.String(typeToken("BucketComponent")),
		},
	})
	opts = append(opts, aliases)
	err := ctx.RegisterComponentResource(typeToken("BucketComponentV2"), name, component, opts...)
	if err != nil {
		return nil, err
	}

	_, err = NewBucketV2(ctx, name+"child", khulnasoft.Parent(component))
	if err != nil {
		return nil, err
	}

	return component, nil
}

const (
	providerName = "wibble"
	version      = "0.0.1"
)

func typeToken(t string) string {
	return fmt.Sprintf("%s:index:%s", providerName, t)
}

var currentID int

func main() {
	err := provider.Main(providerName, func(host *provider.HostClient) (khulnasoftrpc.ResourceProviderServer, error) {
		return makeProvider(host, providerName, version)
	})
	if err != nil {
		cmdutil.ExitError(err.Error())
	}
}

type Provider struct {
	khulnasoftrpc.UnimplementedResourceProviderServer

	host    *provider.HostClient
	name    string
	version string
}

func makeProvider(host *provider.HostClient, name, version string) (khulnasoftrpc.ResourceProviderServer, error) {
	return &Provider{
		host:    host,
		name:    name,
		version: version,
	}, nil
}

func (p *Provider) Create(ctx context.Context, req *khulnasoftrpc.CreateRequest) (*khulnasoftrpc.CreateResponse, error) {
	id := currentID
	currentID++

	return &khulnasoftrpc.CreateResponse{
		Id: fmt.Sprintf("%v", id),
	}, nil
}

func (p *Provider) Construct(ctx context.Context,
	req *khulnasoftrpc.ConstructRequest,
) (*khulnasoftrpc.ConstructResponse, error) {
	return khulnasoftprovider.Construct(ctx, req, p.host.EngineConn(), func(ctx *khulnasoft.Context, typ, name string,
		inputs khulnasoftprovider.ConstructInputs, options khulnasoft.ResourceOption,
	) (*khulnasoftprovider.ConstructResult, error) {
		var component khulnasoft.ComponentResource
		var err error
		switch typ {
		case typeToken("BucketComponent"):
			component, err = NewBucketComponent(ctx, name, options)
		case typeToken("BucketComponentV2"):
			component, err = NewBucketComponentV2(ctx, name, options)
		default:
			err = fmt.Errorf("unknown resource type %s", req.GetType())
		}
		if err != nil {
			return nil, fmt.Errorf("creating component: %w", err)
		}

		return khulnasoftprovider.NewConstructResult(component)
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
	return &khulnasoftrpc.ConfigureResponse{
		AcceptSecrets:   true,
		SupportsPreview: true,
		AcceptResources: true,
	}, nil
}

func (p *Provider) Invoke(ctx context.Context,
	req *khulnasoftrpc.InvokeRequest,
) (*khulnasoftrpc.InvokeResponse, error) {
	return nil, fmt.Errorf("Unknown Invoke token '%s'", req.GetTok())
}

func (p *Provider) StreamInvoke(req *khulnasoftrpc.InvokeRequest,
	server khulnasoftrpc.ResourceProvider_StreamInvokeServer,
) error {
	return fmt.Errorf("Unknown StreamInvoke token '%s'", req.GetTok())
}

func (p *Provider) Call(ctx context.Context,
	req *khulnasoftrpc.CallRequest,
) (*khulnasoftrpc.CallResponse, error) {
	return nil, fmt.Errorf("Unknown Call token '%s'", req.GetTok())
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
