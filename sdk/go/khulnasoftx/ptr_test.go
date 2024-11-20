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

package khulnasoftx_test

import (
	"context"
	"testing"

	"github.com/khulnasoft/khulnasoft/sdk/v3/go/internal"
	"github.com/khulnasoft/khulnasoft/sdk/v3/go/khulnasoft"
	"github.com/khulnasoft/khulnasoft/sdk/v3/go/khulnasoftx"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestPtr(t *testing.T) {
	t.Parallel()

	o := khulnasoftx.Ptr(42)
	v, known, secret, deps, err := internal.AwaitOutput(context.Background(), o)
	require.NoError(t, err)
	assert.True(t, known)
	assert.False(t, secret)
	assert.Empty(t, deps)

	assert.Equal(t, 42, *v.(*int))
}

func TestPtrOf(t *testing.T) {
	t.Parallel()

	vo := khulnasoftx.Val("foo")
	o := khulnasoftx.PtrOf[string](vo)

	v, known, secret, deps, err := internal.AwaitOutput(context.Background(), o)
	require.NoError(t, err)
	assert.True(t, known)
	assert.False(t, secret)
	assert.Empty(t, deps)

	assert.Equal(t, "foo", *v.(*string))
}

func TestPtrOutput_Elem(t *testing.T) {
	t.Parallel()

	o := khulnasoftx.Ptr(42)
	v, _, _, _, err := internal.AwaitOutput(context.Background(), o.Elem())
	require.NoError(t, err)
	assert.Equal(t, 42, v)
}

func TestPtrOutput_Elem_nil(t *testing.T) {
	t.Parallel()

	o := khulnasoftx.Cast[khulnasoftx.PtrOutput[string], *string](khulnasoftx.Val[*string]((*string)(nil)))

	v, _, _, _, err := internal.AwaitOutput(context.Background(), o.Elem())
	require.NoError(t, err)
	assert.Equal(t, "", v)
}

func TestGPtrOutput(t *testing.T) {
	t.Parallel()

	o := khulnasoftx.GPtrOutput[string, khulnasoftx.Output[string]](khulnasoftx.Ptr("foo"))
	v, known, secret, deps, err := internal.AwaitOutput(context.Background(), o)
	require.NoError(t, err)
	assert.True(t, known)
	assert.False(t, secret)
	assert.Empty(t, deps)

	assert.Equal(t, "foo", *v.(*string))
}

func TestGPtrOutput_Elem(t *testing.T) {
	t.Parallel()

	// Given an Output[*string], we want a pux.GPtrOutput
	// that will produce a khulnasoft.StringOutput when Elem is called.
	o := khulnasoftx.Cast[
		khulnasoftx.GPtrOutput[string, khulnasoft.StringOutput],
		*string,
	](khulnasoftx.Ptr("foo"))

	so := o.Elem()
	assert.IsType(t, khulnasoft.StringOutput{}, so)

	v, _, _, _, err := internal.AwaitOutput(context.Background(), so)
	require.NoError(t, err)
	assert.Equal(t, "foo", v)
}
