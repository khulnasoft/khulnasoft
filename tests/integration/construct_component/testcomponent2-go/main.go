// Copyright 2016-2022, Pulumi Corporation.  All rights reserved.
//go:build !all
// +build !all

package main

import (
	"context"
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/khulnasoft/khulnasoft/pkg/v3/resource/provider"
	"github.com/khulnasoft/khulnasoft/sdk/v3/go/common/resource"
	"github.com/khulnasoft/khulnasoft/sdk/v3/go/common/util/cmdutil"
	"github.com/khulnasoft/khulnasoft/sdk/v3/go/khulnasoft"
	"github.com/khulnasoft/khulnasoft/sdk/v3/go/khulnasoft/config"
	khulnasoftprovider "github.com/khulnasoft/khulnasoft/sdk/v3/go/khulnasoft/provider"
	khulnasoftrpc "github.com/khulnasoft/khulnasoft/sdk/v3/proto/go"

	"google.golang.org/protobuf/types/known/emptypb"
)

type Resource struct {
	khulnasoft.CustomResourceState
}

type resourceArgs struct {
	Echo interface{} `khulnasoft:"echo"`
}

type ResourceArgs struct {
	Echo khulnasoft.Input
}

func (ResourceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*resourceArgs)(nil)).Elem()
}

func NewResource(ctx *khulnasoft.Context, name string, echo khulnasoft.Input,
	opts ...khulnasoft.ResourceOption,
) (*Resource, error) {
	args := &ResourceArgs{Echo: echo}
	var resource Resource
	err := ctx.RegisterResource("testcomponent:index:Resource", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

type Component struct {
	khulnasoft.ResourceState

	Echo    khulnasoft.Input        `khulnasoft:"echo"`
	ChildID khulnasoft.IDOutput     `khulnasoft:"childId"`
	Secret  khulnasoft.StringOutput `khulnasoft:"secret"`
}

type ComponentArgs struct {
	Echo khulnasoft.Input `khulnasoft:"echo"`
}

func NewComponentComponent(ctx *khulnasoft.Context, name string, opts ...khulnasoft.ResourceOption) (*Component, error) {
	component := &Component{}
	err := ctx.RegisterComponentResource(providerName+":index:ComponentComponent", name, component, opts...)
	if err != nil {
		return nil, err
	}
	err = ctx.RegisterRemoteComponentResource("testcomponent:index:Component", name+"-child", khulnasoft.Map{
		"echo": khulnasoft.String("checkExpected"),
	}, &Component{}, khulnasoft.Parent(component))
	if err != nil {
		return nil, err
	}

	err = ctx.RegisterResourceOutputs(component, khulnasoft.Map{})
	return component, err
}

func NewComponent(ctx *khulnasoft.Context, name string, args *ComponentArgs,
	opts ...khulnasoft.ResourceOption,
) (*Component, error) {
	if args == nil {
		return nil, errors.New("args is required")
	}

	secretKey := "secret"
	fullSecretKey := fmt.Sprintf("%s:%s", ctx.Project(), secretKey)
	if !ctx.IsConfigSecret(fullSecretKey) {
		return nil, fmt.Errorf("expected configuration key to be secret: %s", fullSecretKey)
	}

	conf := config.New(ctx, "")
	secret := conf.RequireSecret(secretKey)

	component := &Component{}
	err := ctx.RegisterComponentResource(providerName+":index:Component", name, component, opts...)
	if err != nil {
		return nil, err
	}

	res, err := NewResource(ctx, fmt.Sprintf("child-%s", name), args.Echo, khulnasoft.Parent(component))
	if err != nil {
		return nil, err
	}

	component.Echo = args.Echo
	component.ChildID = res.ID()
	component.Secret = secret

	if err := ctx.RegisterResourceOutputs(component, khulnasoft.Map{
		"secret":  component.Secret,
		"echo":    component.Echo,
		"childId": component.ChildID,
	}); err != nil {
		return nil, err
	}

	return component, nil
}

const (
	providerName = "secondtestcomponent"
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

func (p *Provider) Create(ctx context.Context,
	req *khulnasoftrpc.CreateRequest,
) (*khulnasoftrpc.CreateResponse, error) {
	urn := resource.URN(req.GetUrn())
	typ := urn.Type()
	if typ != providerName+":index:Resource" {
		return nil, fmt.Errorf("Unknown resource type '%s'", typ)
	}

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
		switch strings.TrimPrefix(typ, providerName+":index:") {
		case "Component":
			args := &ComponentArgs{}
			if err := inputs.CopyTo(args); err != nil {
				return nil, fmt.Errorf("setting args: %w", err)
			}

			component, err := NewComponent(ctx, name, args, options)
			if err != nil {
				return nil, fmt.Errorf("creating component: %w", err)
			}

			return khulnasoftprovider.NewConstructResult(component)
		case "ComponentComponent":
			component, err := NewComponentComponent(ctx, name, options)
			if err != nil {
				return nil, fmt.Errorf("creating component: %w", err)
			}
			return khulnasoftprovider.NewConstructResult(component)
		default:
			return nil, fmt.Errorf("unknown resource type %s", typ)
		}
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
