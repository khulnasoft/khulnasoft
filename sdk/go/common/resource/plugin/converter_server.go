// Copyright 2016-2023, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package plugin

import (
	"context"

	"github.com/khulnasoft/khulnasoft/sdk/v3/go/common/slice"
	khulnasoftrpc "github.com/khulnasoft/khulnasoft/sdk/v3/proto/go"
	codegenrpc "github.com/khulnasoft/khulnasoft/sdk/v3/proto/go/codegen"
)

type converterServer struct {
	khulnasoftrpc.UnsafeConverterServer // opt out of forward compat

	converter Converter
}

func NewConverterServer(converter Converter) khulnasoftrpc.ConverterServer {
	return &converterServer{converter: converter}
}

func (c *converterServer) ConvertState(ctx context.Context,
	req *khulnasoftrpc.ConvertStateRequest,
) (*khulnasoftrpc.ConvertStateResponse, error) {
	resp, err := c.converter.ConvertState(ctx, &ConvertStateRequest{
		MapperTarget: req.MapperTarget,
		Args:         req.Args,
	})
	if err != nil {
		return nil, err
	}

	resources := make([]*khulnasoftrpc.ResourceImport, len(resp.Resources))
	for i, resource := range resp.Resources {
		resources[i] = &khulnasoftrpc.ResourceImport{
			Type:              resource.Type,
			Name:              resource.Name,
			Id:                resource.ID,
			Version:           resource.Version,
			PluginDownloadURL: resource.PluginDownloadURL,
			LogicalName:       resource.LogicalName,
			IsRemote:          resource.IsRemote,
			IsComponent:       resource.IsComponent,
		}
	}

	// Translate the hcl.Diagnostics into rpc diagnostics.
	diags := slice.Prealloc[*codegenrpc.Diagnostic](len(resp.Diagnostics))
	for _, diag := range resp.Diagnostics {
		diags = append(diags, HclDiagnosticToRPCDiagnostic(diag))
	}

	rpcResp := &khulnasoftrpc.ConvertStateResponse{
		Resources:   resources,
		Diagnostics: diags,
	}
	return rpcResp, nil
}

func (c *converterServer) ConvertProgram(ctx context.Context,
	req *khulnasoftrpc.ConvertProgramRequest,
) (*khulnasoftrpc.ConvertProgramResponse, error) {
	resp, err := c.converter.ConvertProgram(ctx, &ConvertProgramRequest{
		SourceDirectory: req.SourceDirectory,
		TargetDirectory: req.TargetDirectory,
		MapperTarget:    req.MapperTarget,
		LoaderTarget:    req.LoaderTarget,
		Args:            req.Args,
	})
	if err != nil {
		return nil, err
	}

	// Translate the hcl.Diagnostics into rpc diagnostics.
	diags := slice.Prealloc[*codegenrpc.Diagnostic](len(resp.Diagnostics))
	for _, diag := range resp.Diagnostics {
		diags = append(diags, HclDiagnosticToRPCDiagnostic(diag))
	}

	return &khulnasoftrpc.ConvertProgramResponse{
		Diagnostics: diags,
	}, nil
}
