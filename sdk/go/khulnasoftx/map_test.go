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

func TestMap(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	arr := khulnasoftx.Map[string]{
		"1": khulnasoftx.Val("foo"),
		"2": khulnasoft.String("bar"),
		"3": khulnasoftx.Ptr("baz").Elem(),
	}.ToOutput(ctx)

	val, known, secret, deps, err := internal.AwaitOutput(ctx, arr)
	require.NoError(t, err)
	assert.True(t, known)
	assert.False(t, secret)
	assert.Empty(t, deps)

	assert.Equal(t, map[string]string{
		"1": "foo",
		"2": "bar",
		"3": "baz",
	}, val)
}

func TestGMapOutput(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	o := khulnasoftx.Cast[khulnasoftx.GMapOutput[int, khulnasoft.IntOutput], map[string]int](
		khulnasoftx.Map[int]{
			"foo": khulnasoft.Int(1),
			"bar": khulnasoftx.Val(2),
			"baz": khulnasoftx.Ptr(3).Elem(),
		},
	)

	t.Run("MapIndex/match", func(t *testing.T) {
		t.Parallel()

		el := o.MapIndex(khulnasoftx.Val("foo"))
		assert.IsType(t, khulnasoft.IntOutput{}, el)

		val, _, _, _, err := internal.AwaitOutput(ctx, el)
		require.NoError(t, err)
		assert.Equal(t, 1, val)
	})

	t.Run("index/unknown", func(t *testing.T) {
		t.Parallel()

		el := o.MapIndex(khulnasoftx.Val("not a key"))
		val, _, _, _, err := internal.AwaitOutput(ctx, el)
		require.NoError(t, err)
		assert.Zero(t, val)
	})

	t.Run("value", func(t *testing.T) {
		t.Parallel()

		v, known, secret, deps, err := internal.AwaitOutput(ctx, o)
		require.NoError(t, err)
		assert.True(t, known)
		assert.False(t, secret)
		assert.Empty(t, deps)
		assert.Equal(t, map[string]int{
			"foo": 1,
			"bar": 2,
			"baz": 3,
		}, v)
	})
}
