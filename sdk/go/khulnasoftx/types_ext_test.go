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

package khulnasoftx_test // to avoid import cycles

import (
	"context"
	"testing"

	"github.com/khulnasoft/khulnasoft/sdk/v3/go/internal"
	"github.com/khulnasoft/khulnasoft/sdk/v3/go/khulnasoft"
	"github.com/khulnasoft/khulnasoft/sdk/v3/go/khulnasoftx"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Tests that use the full machinery of inputs, outputs, etc.
// must be in a separate package than khulnasoftx itself.
//
// This is necessary because if we don't import "khulnasoft",
// the internal machinery of inputs and outputs is not
// initialized, and we get a panic when we try to use it.
//
// See internal.AnyOutputType for details.

func TestOutput_Untyped(t *testing.T) {
	t.Parallel()

	o := khulnasoftx.Val(42).Untyped().(khulnasoft.IntOutput)
	v, known, secret, _, err := internal.AwaitOutput(context.Background(), o)
	require.NoError(t, err)
	assert.True(t, known)
	assert.False(t, secret)
	assert.Equal(t, 42, v)
}

func TestOutput_AsAny(t *testing.T) {
	t.Parallel()

	o := khulnasoftx.Val("foo").AsAny()

	v, known, secret, _, err := internal.AwaitOutput(context.Background(), o)
	require.NoError(t, err)
	assert.True(t, known)
	assert.False(t, secret)
	assert.Equal(t, "foo", v)
}

func TestOutput_ConvertTyped(t *testing.T) {
	t.Parallel()

	stringOut := khulnasoft.String("bar").ToStringOutput()
	out, err := khulnasoftx.ConvertTyped[string](stringOut)
	require.NoError(t, err)

	v, _, _, _, err := internal.AwaitOutput(context.Background(), out)
	require.NoError(t, err)
	assert.Equal(t, "bar", v)
}

func TestOutput_ConvertTyped_error(t *testing.T) {
	t.Parallel()

	stringOut := khulnasoft.String("bar").ToStringOutput()

	_, err := khulnasoftx.ConvertTyped[int](stringOut)
	assert.ErrorContains(t, err, "cannot convert string to int")
}

func TestOutput_MustConvertTyped(t *testing.T) {
	t.Parallel()

	stringOut := khulnasoft.String("bar").ToStringOutput()
	out := khulnasoftx.MustConvertTyped[string](stringOut)

	v, _, _, _, err := internal.AwaitOutput(context.Background(), out)
	require.NoError(t, err)
	assert.Equal(t, "bar", v)
}

func TestOutput_MustConvertTyped_error(t *testing.T) {
	t.Parallel()

	stringOut := khulnasoft.String("bar").ToStringOutput()

	assert.PanicsWithError(t, "cannot convert string to int", func() {
		khulnasoftx.MustConvertTyped[int](stringOut)
	})
}
