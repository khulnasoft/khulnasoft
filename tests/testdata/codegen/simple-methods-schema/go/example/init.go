// Code generated by test DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package example

import (
	"fmt"

	"github.com/blang/semver"
	"github.com/khulnasoft/khulnasoft/pkg/v3/codegen/testing/test/testdata/simple-methods-schema/go/example/internal"
	"github.com/khulnasoft/khulnasoft/sdk/v3/go/khulnasoft"
)

type module struct {
	version semver.Version
}

func (m *module) Version() semver.Version {
	return m.version
}

func (m *module) Construct(ctx *khulnasoft.Context, name, typ, urn string) (r khulnasoft.Resource, err error) {
	switch typ {
	case "example::Foo":
		r = &Foo{}
	default:
		return nil, fmt.Errorf("unknown resource type: %s", typ)
	}

	err = ctx.RegisterResource(typ, name, nil, r, khulnasoft.URN_(urn))
	return
}

type pkg struct {
	version semver.Version
}

func (p *pkg) Version() semver.Version {
	return p.version
}

func (p *pkg) ConstructProvider(ctx *khulnasoft.Context, name, typ, urn string) (khulnasoft.ProviderResource, error) {
	if typ != "khulnasoft:providers:example" {
		return nil, fmt.Errorf("unknown provider type: %s", typ)
	}

	r := &Provider{}
	err := ctx.RegisterResource(typ, name, nil, r, khulnasoft.URN_(urn))
	return r, err
}

func init() {
	version, err := internal.PkgVersion()
	if err != nil {
		version = semver.Version{Major: 1}
	}
	khulnasoft.RegisterResourceModule(
		"example",
		"",
		&module{version},
	)
	khulnasoft.RegisterResourcePackage(
		"example",
		&pkg{version},
	)
}
