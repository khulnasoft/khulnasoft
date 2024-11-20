// Copyright 2016-2021, Pulumi Corporation.
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

package provider

import (
	"context"

	"github.com/khulnasoft/khulnasoft/sdk/v3/go/khulnasoft"
	khulnasoftrpc "github.com/khulnasoft/khulnasoft/sdk/v3/proto/go"

	"google.golang.org/grpc"
)

// This file relies on implementations in ../provider_linked.go that are made available in this
// package via go:linkname.

type ConstructFunc func(ctx *khulnasoft.Context, typ, name string, inputs ConstructInputs,
	options khulnasoft.ResourceOption) (*ConstructResult, error)

// Construct adapts the gRPC ConstructRequest/ConstructResponse to/from the Pulumi Go SDK programming model.
func Construct(ctx context.Context, req *khulnasoftrpc.ConstructRequest, engineConn *grpc.ClientConn,
	construct ConstructFunc,
) (*khulnasoftrpc.ConstructResponse, error) {
	return linkedConstruct(ctx, req, engineConn, func(khulnasoftCtx *khulnasoft.Context, typ, name string,
		inputs map[string]interface{}, options khulnasoft.ResourceOption,
	) (khulnasoft.URNInput, khulnasoft.Input, error) {
		ci := ConstructInputs{ctx: khulnasoftCtx, inputs: inputs}
		result, err := construct(khulnasoftCtx, typ, name, ci, options)
		if err != nil {
			return nil, nil, err
		}
		return result.URN, result.State, nil
	})
}

// ConstructInputs represents the inputs associated with a call to Construct.
type ConstructInputs struct {
	ctx    *khulnasoft.Context
	inputs map[string]interface{}
}

// Map returns the inputs as a Map.
func (inputs ConstructInputs) Map() (khulnasoft.Map, error) {
	return linkedConstructInputsMap(inputs.ctx, inputs.inputs)
}

// CopyTo sets the inputs on the given args struct.
func (inputs ConstructInputs) CopyTo(args interface{}) error {
	return linkedConstructInputsCopyTo(inputs.ctx, inputs.inputs, args)
}

// ConstructResult is the result of a call to Construct.
type ConstructResult struct {
	URN   khulnasoft.URNInput
	State khulnasoft.Input
}

// NewConstructResult creates a ConstructResult from the resource.
func NewConstructResult(resource khulnasoft.ComponentResource) (*ConstructResult, error) {
	urn, state, err := linkedNewConstructResult(resource)
	if err != nil {
		return nil, err
	}
	return &ConstructResult{
		URN:   urn,
		State: state,
	}, nil
}

type CallFunc func(ctx *khulnasoft.Context, tok string, args CallArgs) (*CallResult, error)

// Call adapts the gRPC CallRequest/CallResponse to/from the Pulumi Go SDK programming model.
func Call(ctx context.Context, req *khulnasoftrpc.CallRequest, engineConn *grpc.ClientConn,
	call CallFunc,
) (*khulnasoftrpc.CallResponse, error) {
	return linkedCall(ctx, req, engineConn, func(khulnasoftCtx *khulnasoft.Context, tok string,
		args map[string]interface{},
	) (khulnasoft.Input, []interface{}, error) {
		ca := CallArgs{ctx: khulnasoftCtx, args: args}
		result, err := call(khulnasoftCtx, tok, ca)
		if err != nil {
			return nil, nil, err
		}
		var failures []interface{}
		if len(result.Failures) > 0 {
			failures = make([]interface{}, len(result.Failures))
			for i, v := range result.Failures {
				failures[i] = linkedNewCallFailure(v.Property, v.Reason)
			}
		}
		return result.Return, failures, nil
	})
}

// CallArgs represents the Call's arguments.
type CallArgs struct {
	ctx  *khulnasoft.Context
	args map[string]interface{}
}

// Map returns the args as a Map.
func (a CallArgs) Map() (khulnasoft.Map, error) {
	// Use the same implementation as construct.
	return linkedConstructInputsMap(a.ctx, a.args)
}

// CopyTo sets the args on the given args struct. If there is a `__self__` argument, it will be
// returned, otherwise it will return nil.
func (a CallArgs) CopyTo(args interface{}) (khulnasoft.Resource, error) {
	return linkedCallArgsCopyTo(a.ctx, a.args, args)
}

// Self retrieves the `__self__` argument. If `__self__` is present the value is returned,
// otherwise the returned value will be nil.
func (a CallArgs) Self() (khulnasoft.Resource, error) {
	return linkedCallArgsSelf(a.ctx, a.args)
}

// CallFailure indicates that a call to Call failed; it contains the property and reason for the failure.
type CallFailure struct {
	Property string // the property that failed checking.
	Reason   string // the reason the property failed to check.
}

// CallResult is the result of the Call.
type CallResult struct {
	// The returned values, if the call was successful.
	Return khulnasoft.Input
	// The failures if any arguments didn't pass verification.
	Failures []CallFailure
}

// NewCallResult creates a CallResult from the given result.
func NewCallResult(result interface{}) (*CallResult, error) {
	ret, err := linkedNewCallResult(result)
	if err != nil {
		return nil, err
	}
	return &CallResult{
		Return: ret,
	}, nil
}

// NewCallResult expects a pointer to a struct to return multiple properties. When a function call
// returns only a single property, use NewSingletonCallResult instead.
func NewSingletonCallResult[T any](result T) (*CallResult, error) {
	type wrapper struct {
		// Wrapping the value in a "res" magic property by convention.
		Result T `khulnasoft:"res"`
	}
	return NewCallResult(&wrapper{result})
}

type constructFunc func(ctx *khulnasoft.Context, typ, name string, inputs map[string]interface{},
	options khulnasoft.ResourceOption) (khulnasoft.URNInput, khulnasoft.Input, error)

// linkedConstruct is made available here from ../provider_linked.go via go:linkname.
func linkedConstruct(ctx context.Context, req *khulnasoftrpc.ConstructRequest, engineConn *grpc.ClientConn,
	constructF constructFunc) (*khulnasoftrpc.ConstructResponse, error)

// linkedConstructInputsMap is made available here from ../provider_linked.go via go:linkname.
func linkedConstructInputsMap(ctx *khulnasoft.Context, inputs map[string]interface{}) (khulnasoft.Map, error)

// linkedConstructInputsCopyTo is made available here from ../provider_linked.go via go:linkname.
func linkedConstructInputsCopyTo(ctx *khulnasoft.Context, inputs map[string]interface{}, args interface{}) error

// linkedNewConstructResult is made available here from ../provider_linked.go via go:linkname.
func linkedNewConstructResult(resource khulnasoft.ComponentResource) (khulnasoft.URNInput, khulnasoft.Input, error)

type callFunc func(ctx *khulnasoft.Context, tok string, args map[string]interface{}) (khulnasoft.Input, []interface{}, error)

// linkedCall is made available here from ../provider_linked.go via go:linkname.
func linkedCall(ctx context.Context, req *khulnasoftrpc.CallRequest, engineConn *grpc.ClientConn,
	callF callFunc) (*khulnasoftrpc.CallResponse, error)

// linkedCallArgsCopyTo is made available here from ../provider_linked.go via go:linkname.
func linkedCallArgsCopyTo(ctx *khulnasoft.Context, source map[string]interface{},
	args interface{}) (khulnasoft.Resource, error)

// linkedCallArgsSelf is made available here from ../provider_linked.go via go:linkname.
func linkedCallArgsSelf(ctx *khulnasoft.Context, source map[string]interface{}) (khulnasoft.Resource, error)

// linkedNewCallResult is made available here from ../provider_linked.go via go:linkname.
func linkedNewCallResult(result interface{}) (khulnasoft.Input, error)

// linkedNewCallFailure is made available here from ../provider_linked.go via go:linkname.
func linkedNewCallFailure(property, reason string) interface{}
