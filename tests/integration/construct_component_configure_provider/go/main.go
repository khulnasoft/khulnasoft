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

package main

import (
	"github.com/khulnasoft/khulnasoft-tls/sdk/v4/go/tls"
	"github.com/khulnasoft/khulnasoft/sdk/v3/go/khulnasoft"
	"github.com/khulnasoft/khulnasoft/sdk/v3/go/khulnasoft/config"
	"github.com/khulnasoft/khulnasoft/tests/testdata/codegen/methods-return-plain-resource/go/metaprovider"
)

func main() {
	khulnasoft.Run(func(ctx *khulnasoft.Context) error {
		cfg := config.New(ctx, "")
		proxy := cfg.Require("proxy")

		configurer, err := metaprovider.NewConfigurer(ctx, "configurer", &metaprovider.ConfigurerArgs{
			TlsProxy: unknownIfDryRun(ctx, proxy),
		})
		if err != nil {
			return err
		}

		prov, err := configurer.TlsProvider(ctx)
		if err != nil {
			return err
		}

		key, err := tls.NewPrivateKey(ctx, "my-private-key", &tls.PrivateKeyArgs{
			Algorithm:  khulnasoft.String("ECDSA"),
			EcdsaCurve: khulnasoft.String("P384"),
		}, khulnasoft.Provider(prov))
		if err != nil {
			return err
		}

		var n int
		n, err = configurer.MeaningOfLife(ctx)
		if err != nil {
			return err
		}

		mix, err := configurer.ObjectMix(ctx)
		if err != nil {
			return err
		}

		ctx.Export("meaningOfLife", khulnasoft.Int(n))
		ctx.Export("keyAlgo", key.Algorithm)
		if mix.MeaningOfLife != nil {
			ctx.Export("meaningOfLife2", khulnasoft.Int(*mix.MeaningOfLife))
		}

		key2, err := tls.NewPrivateKey(ctx, "my-private-key-2", &tls.PrivateKeyArgs{
			Algorithm:  khulnasoft.String("ECDSA"),
			EcdsaCurve: khulnasoft.String("P384"),
		}, khulnasoft.Provider(mix.Provider))
		if err != nil {
			return err
		}

		ctx.Export("keyAlgo2", key2.Algorithm)
		return nil
	})
}

func unknownIfDryRun(ctx *khulnasoft.Context, value string) khulnasoft.StringOutput {
	if ctx.DryRun() {
		return khulnasoft.UnsafeUnknownOutput(nil).ApplyT(func(_ any) string {
			panic("impossible")
		}).(khulnasoft.StringOutput)
	}
	return khulnasoft.String(value).ToStringOutput()
}
