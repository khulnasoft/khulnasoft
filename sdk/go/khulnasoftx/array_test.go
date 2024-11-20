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

func TestArray(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	arr := khulnasoftx.Array[string]{
		khulnasoftx.Val("foo"),
		khulnasoft.String("bar"),
		khulnasoftx.Ptr("baz").Elem(),
	}.ToOutput(ctx)

	val, known, secret, deps, err := internal.AwaitOutput(ctx, arr)
	require.NoError(t, err)
	assert.True(t, known)
	assert.False(t, secret)
	assert.Empty(t, deps)

	assert.Equal(t, []string{"foo", "bar", "baz"}, val)
}

func TestGArrayOutput(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	o := khulnasoftx.Cast[khulnasoftx.GArrayOutput[string, khulnasoft.StringOutput], []string](
		khulnasoftx.Array[string]{
			khulnasoft.String("foo"),
			khulnasoftx.Val("bar"),
			khulnasoftx.Ptr("baz").Elem(),
		},
	)

	t.Run("index/match", func(t *testing.T) {
		t.Parallel()

		el := o.Index(khulnasoftx.Val(0))
		assert.IsType(t, khulnasoft.StringOutput{}, el)

		val, _, _, _, err := internal.AwaitOutput(ctx, el)
		require.NoError(t, err)
		assert.Equal(t, "foo", val)
	})

	t.Run("index/out of bounds", func(t *testing.T) {
		t.Parallel()

		el := o.Index(khulnasoftx.Val(42))
		val, _, _, _, err := internal.AwaitOutput(ctx, el)
		require.NoError(t, err)
		assert.Empty(t, val)
	})

	t.Run("value", func(t *testing.T) {
		t.Parallel()

		v, known, secret, deps, err := internal.AwaitOutput(ctx, o)
		require.NoError(t, err)
		assert.True(t, known)
		assert.False(t, secret)
		assert.Empty(t, deps)
		assert.Equal(t, []string{"foo", "bar", "baz"}, v)
	})
}
