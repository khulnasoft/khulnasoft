// Copyright 2016-2022, Pulumi Corporation.
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

// Code generated by "generate.go"; DO NOT EDIT.

package config

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cast"

	"github.com/khulnasoft/khulnasoft/sdk/v3/go/khulnasoft"
)

func failf(format string, a ...interface{}) {
	panic(fmt.Errorf(format, a...))
}

func require(ctx *khulnasoft.Context, key string, secret bool, use, insteadOf string) string {
	v, ok := get(ctx, key, use, insteadOf)

	secretText := " "
	if secret {
		secretText = " --secret "
	}

	if !ok {
		failf("Missing required configuration variable '%s'\n"+
			"\tplease set a value using the command `khulnasoft config set%s%s <value>`",
			key, secretText, key)
	}
	return v
}

// Require loads a configuration value by its key, or panics if it doesn't exist.
func Require(ctx *khulnasoft.Context, key string) string {
	return require(ctx, key, false, "RequireSecret", "Require")
}

func requireObject(ctx *khulnasoft.Context, key string, secret bool, output interface{}, use, insteadOf string) {
	v := require(ctx, key, secret, use, insteadOf)
	if err := json.Unmarshal([]byte(v), output); err != nil {
		failf("unable to unmarshall required configuration variable '%s'; %s", key, err)
	}
}

// RequireObject loads an optional configuration value by its key into the output variable,
// or panics if unable to do so.
func RequireObject(ctx *khulnasoft.Context, key string, output interface{}) {
	requireObject(ctx, key, false, output, "RequireSecretObject", "RequireObject")
}

func requireBool(ctx *khulnasoft.Context, key string, secret bool, use, insteadOf string) bool {
	v := require(ctx, key, secret, use, insteadOf)
	o, err := cast.ToBoolE(v)
	if err != nil {
		failf("unable to parse required configuration variable '%s'; %s", key, err)
	}
	return o
}

// RequireBool loads an optional configuration value by its key, as a bool, or panics if it doesn't exist.
func RequireBool(ctx *khulnasoft.Context, key string) bool {
	return requireBool(ctx, key, false, "RequireSecretBool", "RequireBool")
}

func requireFloat64(ctx *khulnasoft.Context, key string, secret bool, use, insteadOf string) float64 {
	v := require(ctx, key, secret, use, insteadOf)
	o, err := cast.ToFloat64E(v)
	if err != nil {
		failf("unable to parse required configuration variable '%s'; %s", key, err)
	}
	return o
}

// RequireFloat64 loads an optional configuration value by its key, as a float64, or panics if it doesn't exist.
func RequireFloat64(ctx *khulnasoft.Context, key string) float64 {
	return requireFloat64(ctx, key, false, "RequireSecretFloat64", "RequireFloat64")
}

func requireInt(ctx *khulnasoft.Context, key string, secret bool, use, insteadOf string) int {
	v := require(ctx, key, secret, use, insteadOf)
	o, err := cast.ToIntE(v)
	if err != nil {
		failf("unable to parse required configuration variable '%s'; %s", key, err)
	}
	return o
}

// RequireInt loads an optional configuration value by its key, as a int, or panics if it doesn't exist.
func RequireInt(ctx *khulnasoft.Context, key string) int {
	return requireInt(ctx, key, false, "RequireSecretInt", "RequireInt")
}

// RequireSecret loads a configuration value by its key returning it wrapped in a secret Output,
// or panics if it doesn't exist.
func RequireSecret(ctx *khulnasoft.Context, key string) khulnasoft.StringOutput {
	return khulnasoft.ToSecret(require(ctx, key, true, "", "")).(khulnasoft.StringOutput)
}

// RequireSecretObject loads an optional configuration value by its key into the output variable,
// returning it wrapped in a secret Output, or panics if unable to do so.
func RequireSecretObject(ctx *khulnasoft.Context, key string, output interface{}) khulnasoft.Output {
	requireObject(ctx, key, true, output, "", "")
	return khulnasoft.ToSecret(output)
}

// RequireSecretBool loads an optional configuration value by its key,
// as a bool wrapped in a secret Output, or panics if it doesn't exist.
func RequireSecretBool(ctx *khulnasoft.Context, key string) khulnasoft.BoolOutput {
	return khulnasoft.ToSecret(requireBool(ctx, key, true, "", "")).(khulnasoft.BoolOutput)
}

// RequireSecretFloat64 loads an optional configuration value by its key,
// as a float64 wrapped in a secret Output, or panics if it doesn't exist.
func RequireSecretFloat64(ctx *khulnasoft.Context, key string) khulnasoft.Float64Output {
	return khulnasoft.ToSecret(requireFloat64(ctx, key, true, "", "")).(khulnasoft.Float64Output)
}

// RequireSecretInt loads an optional configuration value by its key,
// as a int wrapped in a secret Output, or panics if it doesn't exist.
func RequireSecretInt(ctx *khulnasoft.Context, key string) khulnasoft.IntOutput {
	return khulnasoft.ToSecret(requireInt(ctx, key, true, "", "")).(khulnasoft.IntOutput)
}
