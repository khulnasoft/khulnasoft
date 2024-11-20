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
	"reflect"
	"testing"

	"github.com/khulnasoft/khulnasoft/sdk/v3/go/common/resource"
	"github.com/khulnasoft/khulnasoft/sdk/v3/go/khulnasoft"
	"github.com/khulnasoft/khulnasoft/sdk/v3/go/khulnasoftx"
	"github.com/stretchr/testify/require"
)

type testResourceInputs struct {
	PuxString khulnasoftx.Input[string]
	PuString  khulnasoft.StringInput

	PuxStringPtr khulnasoftx.Input[*string]
	PuStringPtr  khulnasoft.StringPtrInput

	PuxIntArray khulnasoftx.Input[[]int]
	PuIntArray  khulnasoft.IntArrayInput

	PuxIntMap khulnasoftx.Input[map[string]int]
	PuIntMap  khulnasoft.IntMapInput
}

func (*testResourceInputs) ElementType() reflect.Type {
	return reflect.TypeOf((*testResourceArgs)(nil))
}

type testResourceArgs struct {
	PuxString string `khulnasoft:"puxString"`
	PuString  string `khulnasoft:"puString"`

	PuxStringPtr *string `khulnasoft:"puxStringPtr"`
	PuStringPtr  *string `khulnasoft:"puStringPtr"`

	PuxIntArray []int `khulnasoft:"puxIntArray"`
	PuIntArray  []int `khulnasoft:"puIntArray"`

	PuxIntMap map[string]int `khulnasoft:"puxIntMap"`
	PuIntMap  map[string]int `khulnasoft:"puIntMap"`
}

func TestRegisterResource_inputSerialization(t *testing.T) {
	t.Parallel()

	j := "j"

	tests := []struct {
		desc string
		give khulnasoft.Input
		want resource.PropertyMap
	}{
		// --- khulnasoftx.Input[string] ---
		{
			desc: "pux.String/pu.String",
			give: &testResourceInputs{PuxString: khulnasoft.String("a")},
			want: resource.PropertyMap{"puxString": resource.NewStringProperty("a")},
		},
		{
			desc: "pux.String/pu.StringOutput",
			give: &testResourceInputs{PuxString: khulnasoft.String("b").ToStringOutput()},
			want: resource.PropertyMap{"puxString": resource.NewStringProperty("b")},
		},
		{
			desc: "pux.String/pux.Output[string]",
			give: &testResourceInputs{PuxString: khulnasoftx.Val("c")},
			want: resource.PropertyMap{"puxString": resource.NewStringProperty("c")},
		},

		// --- khulnasoft.StringInput ---
		{
			desc: "pu.String/pu.String",
			give: &testResourceInputs{PuString: khulnasoft.String("d")},
			want: resource.PropertyMap{"puString": resource.NewStringProperty("d")},
		},
		{
			desc: "pu.String/pu.StringOutput",
			give: &testResourceInputs{PuString: khulnasoft.String("e").ToStringOutput()},
			want: resource.PropertyMap{"puString": resource.NewStringProperty("e")},
		},
		{
			desc: "pu.String/pux.Output[string] untyped",
			give: &testResourceInputs{PuString: khulnasoftx.Val("f").Untyped().(khulnasoft.StringOutput)},
			want: resource.PropertyMap{"puString": resource.NewStringProperty("f")},
		},

		// --- khulnasoftx.Input[*string] ---
		{
			desc: "pux.StringPtr/pu.String PtrOf",
			give: &testResourceInputs{PuxStringPtr: khulnasoftx.PtrOf[string](khulnasoft.String("g"))},
			want: resource.PropertyMap{"puxStringPtr": resource.NewStringProperty("g")},
		},
		{
			desc: "pux.StringPtr/pu.StringPtrOutput",
			give: &testResourceInputs{PuxStringPtr: khulnasoft.String("h").ToStringPtrOutput()},
			want: resource.PropertyMap{"puxStringPtr": resource.NewStringProperty("h")},
		},
		{
			desc: "pux.StringPtr/pux.PtrOutput[string]",
			give: &testResourceInputs{PuxStringPtr: khulnasoftx.Ptr("i")},
			want: resource.PropertyMap{"puxStringPtr": resource.NewStringProperty("i")},
		},
		{
			desc: "pux.StringPtr/pux.Output[*string]",
			give: &testResourceInputs{PuxStringPtr: khulnasoftx.Val[*string](&j)},
			want: resource.PropertyMap{"puxStringPtr": resource.NewStringProperty("j")},
		},

		// --- khulnasoft.StringPtrInput ---
		{
			desc: "pu.StringPtr/pu.StringPtr",
			give: &testResourceInputs{PuStringPtr: khulnasoft.StringPtr("j")},
			want: resource.PropertyMap{"puStringPtr": resource.NewStringProperty("j")},
		},
		{
			desc: "pu.StringPtr/pu.String",
			give: &testResourceInputs{PuStringPtr: khulnasoft.String("k")},
			want: resource.PropertyMap{"puStringPtr": resource.NewStringProperty("k")},
		},
		{
			desc: "pu.StringPtr/pu.StringPtrOutput",
			give: &testResourceInputs{PuStringPtr: khulnasoft.String("l").ToStringPtrOutput()},
			want: resource.PropertyMap{"puStringPtr": resource.NewStringProperty("l")},
		},
		{
			desc: "pu.StringPtr/pux.PtrOutput[string] untyped",
			give: &testResourceInputs{PuStringPtr: khulnasoftx.Ptr("m").Untyped().(khulnasoft.StringPtrOutput)},
			want: resource.PropertyMap{"puStringPtr": resource.NewStringProperty("m")},
		},
		{
			desc: "pu.StringPtr/pux.GPtrOutput[string] untyped",
			give: &testResourceInputs{
				PuStringPtr: khulnasoftx.Cast[khulnasoftx.GPtrOutput[string, khulnasoft.StringOutput], *string](
					khulnasoftx.Ptr("n"),
				).Untyped().(khulnasoft.StringPtrOutput),
			},
			want: resource.PropertyMap{"puStringPtr": resource.NewStringProperty("n")},
		},

		// --- khulnasoftx.Input[[]int] ---
		{
			desc: "pux.IntArray/pu.IntArray",
			give: &testResourceInputs{
				PuxIntArray: khulnasoft.IntArray{
					khulnasoft.Int(1),
					khulnasoft.Int(2),
					khulnasoft.Int(3),
				},
			},
			want: resource.PropertyMap{
				"puxIntArray": resource.NewArrayProperty(
					[]resource.PropertyValue{
						resource.NewNumberProperty(1),
						resource.NewNumberProperty(2),
						resource.NewNumberProperty(3),
					},
				),
			},
		},
		{
			desc: "pux.IntArray/pu.IntArrayOutput",
			give: &testResourceInputs{
				PuxIntArray: khulnasoft.IntArray{
					khulnasoft.Int(4),
					khulnasoft.Int(5),
					khulnasoft.Int(6),
				}.ToIntArrayOutput(),
			},
			want: resource.PropertyMap{
				"puxIntArray": resource.NewArrayProperty(
					[]resource.PropertyValue{
						resource.NewNumberProperty(4),
						resource.NewNumberProperty(5),
						resource.NewNumberProperty(6),
					},
				),
			},
		},
		{
			desc: "pux.IntArray/pux.Output[[]int]",
			give: &testResourceInputs{
				PuxIntArray: khulnasoftx.Val([]int{7, 8, 9}),
			},
			want: resource.PropertyMap{
				"puxIntArray": resource.NewArrayProperty(
					[]resource.PropertyValue{
						resource.NewNumberProperty(7),
						resource.NewNumberProperty(8),
						resource.NewNumberProperty(9),
					},
				),
			},
		},
		{
			desc: "pux.IntArray/pux.ArrayOutput[int]",
			give: &testResourceInputs{
				PuxIntArray: khulnasoftx.Cast[khulnasoftx.ArrayOutput[int], []int](
					khulnasoftx.Val([]int{10, 11, 12}),
				),
			},
			want: resource.PropertyMap{
				"puxIntArray": resource.NewArrayProperty(
					[]resource.PropertyValue{
						resource.NewNumberProperty(10),
						resource.NewNumberProperty(11),
						resource.NewNumberProperty(12),
					},
				),
			},
		},
		{
			desc: "pux.IntArray/pux.GArrayOutput",
			give: &testResourceInputs{
				PuxIntArray: khulnasoftx.Cast[khulnasoftx.GArrayOutput[int, khulnasoft.IntOutput], []int](
					khulnasoftx.Val([]int{13, 14, 15}),
				),
			},
			want: resource.PropertyMap{
				"puxIntArray": resource.NewArrayProperty(
					[]resource.PropertyValue{
						resource.NewNumberProperty(13),
						resource.NewNumberProperty(14),
						resource.NewNumberProperty(15),
					},
				),
			},
		},

		// --- khulnasoft.IntArrayInput ---
		{
			desc: "pu.IntArray/pux.Output[[]int] untyped",
			give: &testResourceInputs{
				PuIntArray: khulnasoftx.Val([]int{1, 2, 3}).Untyped().(khulnasoft.IntArrayOutput),
			},
			want: resource.PropertyMap{
				"puIntArray": resource.NewArrayProperty(
					[]resource.PropertyValue{
						resource.NewNumberProperty(1),
						resource.NewNumberProperty(2),
						resource.NewNumberProperty(3),
					},
				),
			},
		},
		{
			desc: "pu.IntArray/pux.ArrayOutput[int] untyped",
			give: &testResourceInputs{
				PuIntArray: khulnasoftx.Cast[khulnasoftx.ArrayOutput[int], []int](
					khulnasoftx.Val([]int{4, 5, 6}),
				).Untyped().(khulnasoft.IntArrayOutput),
			},
			want: resource.PropertyMap{
				"puIntArray": resource.NewArrayProperty(
					[]resource.PropertyValue{
						resource.NewNumberProperty(4),
						resource.NewNumberProperty(5),
						resource.NewNumberProperty(6),
					},
				),
			},
		},
		{
			desc: "pu.IntArray/pux.GArrayOutput untyped",
			give: &testResourceInputs{
				PuIntArray: khulnasoftx.Cast[khulnasoftx.GArrayOutput[int, khulnasoft.IntOutput], []int](
					khulnasoftx.Val([]int{7, 8, 9}),
				).Untyped().(khulnasoft.IntArrayOutput),
			},
			want: resource.PropertyMap{
				"puIntArray": resource.NewArrayProperty(
					[]resource.PropertyValue{
						resource.NewNumberProperty(7),
						resource.NewNumberProperty(8),
						resource.NewNumberProperty(9),
					},
				),
			},
		},

		// --- khulnasoftx.Input[map[string]int] ---
		{
			desc: "pux.IntMap/pu.IntMap",
			give: &testResourceInputs{
				PuxIntMap: khulnasoft.IntMap{"a": khulnasoft.Int(1), "b": khulnasoft.Int(2)},
			},
			want: resource.PropertyMap{
				"puxIntMap": resource.NewObjectProperty(
					resource.PropertyMap{
						"a": resource.NewNumberProperty(1),
						"b": resource.NewNumberProperty(2),
					},
				),
			},
		},
		{
			desc: "pux.IntMap/pu.IntMapOutput",
			give: &testResourceInputs{
				PuxIntMap: khulnasoft.IntMap{"c": khulnasoft.Int(3), "d": khulnasoft.Int(4)}.ToIntMapOutput(),
			},
			want: resource.PropertyMap{
				"puxIntMap": resource.NewObjectProperty(
					resource.PropertyMap{
						"c": resource.NewNumberProperty(3),
						"d": resource.NewNumberProperty(4),
					},
				),
			},
		},
		{
			desc: "pux.IntMap/pux.Output[map[string]int]",
			give: &testResourceInputs{
				PuxIntMap: khulnasoftx.Val(map[string]int{"e": 5, "f": 6}),
			},
			want: resource.PropertyMap{
				"puxIntMap": resource.NewObjectProperty(
					resource.PropertyMap{
						"e": resource.NewNumberProperty(5),
						"f": resource.NewNumberProperty(6),
					},
				),
			},
		},
		{
			desc: "pux.IntMap/pux.MapOutput[int]",
			give: &testResourceInputs{
				PuxIntMap: khulnasoftx.Cast[khulnasoftx.MapOutput[int], map[string]int](
					khulnasoftx.Val(map[string]int{"g": 7, "h": 8}),
				),
			},
			want: resource.PropertyMap{
				"puxIntMap": resource.NewObjectProperty(
					resource.PropertyMap{
						"g": resource.NewNumberProperty(7),
						"h": resource.NewNumberProperty(8),
					},
				),
			},
		},
		{
			desc: "pux.IntMap/pux.GMapOutput",
			give: &testResourceInputs{
				PuxIntMap: khulnasoftx.Cast[khulnasoftx.GMapOutput[int, khulnasoft.IntOutput], map[string]int](
					khulnasoftx.Val(map[string]int{"i": 9, "j": 10}),
				),
			},
			want: resource.PropertyMap{
				"puxIntMap": resource.NewObjectProperty(
					resource.PropertyMap{
						"i": resource.NewNumberProperty(9),
						"j": resource.NewNumberProperty(10),
					},
				),
			},
		},

		// --- khulnasoft.IntMapInput ---
		{
			desc: "pu.IntMap/pu.IntMap",
			give: &testResourceInputs{
				PuIntMap: khulnasoft.IntMap{"a": khulnasoft.Int(1), "b": khulnasoft.Int(2)},
			},
			want: resource.PropertyMap{
				"puIntMap": resource.NewObjectProperty(
					resource.PropertyMap{
						"a": resource.NewNumberProperty(1),
						"b": resource.NewNumberProperty(2),
					},
				),
			},
		},
		{
			desc: "pu.IntMap/pux.Output[map[string]int] untyped",
			give: &testResourceInputs{
				PuIntMap: khulnasoftx.Val(map[string]int{"c": 3, "d": 4}).Untyped().(khulnasoft.IntMapOutput),
			},
			want: resource.PropertyMap{
				"puIntMap": resource.NewObjectProperty(
					resource.PropertyMap{
						"c": resource.NewNumberProperty(3),
						"d": resource.NewNumberProperty(4),
					},
				),
			},
		},
		{
			desc: "pu.IntMap/pux.MapOutput[int] untyped",
			give: &testResourceInputs{
				PuIntMap: khulnasoftx.Cast[khulnasoftx.MapOutput[int], map[string]int](
					khulnasoftx.Val(map[string]int{"e": 5, "f": 6}),
				).Untyped().(khulnasoft.IntMapOutput),
			},
			want: resource.PropertyMap{
				"puIntMap": resource.NewObjectProperty(
					resource.PropertyMap{
						"e": resource.NewNumberProperty(5),
						"f": resource.NewNumberProperty(6),
					},
				),
			},
		},
		{
			desc: "pu.IntMap/pux.GMapOutput untyped",
			give: &testResourceInputs{
				PuIntMap: khulnasoftx.Cast[khulnasoftx.GMapOutput[int, khulnasoft.IntOutput], map[string]int](
					khulnasoftx.Val(map[string]int{"g": 7, "h": 8}),
				).Untyped().(khulnasoft.IntMapOutput),
			},
			want: resource.PropertyMap{
				"puIntMap": resource.NewObjectProperty(
					resource.PropertyMap{
						"g": resource.NewNumberProperty(7),
						"h": resource.NewNumberProperty(8),
					},
				),
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.desc, func(t *testing.T) {
			t.Parallel()

			var (
				got resource.PropertyMap
				ok  bool
			)
			err := khulnasoft.RunErr(func(ctx *khulnasoft.Context) error {
				var res khulnasoft.CustomResourceState
				require.NoError(t,
					ctx.RegisterResource("test:index:MyResource", "testResource", tt.give, &res))
				return nil
			}, khulnasoft.WithMocks("project", "stack", &mockResourceMonitor{
				CallF: func(args khulnasoft.MockCallArgs) (resource.PropertyMap, error) {
					t.Fatalf("unexpected Call(%#v)", args)
					return nil, nil
				},
				NewResourceF: func(args khulnasoft.MockResourceArgs) (string, resource.PropertyMap, error) {
					if args.TypeToken == "test:index:MyResource" {
						got = args.Inputs
						ok = true
					} else {
						t.Fatalf("unexpected NewResource(%#v)", args)
					}
					return args.Name + "_id", nil, nil
				},
			}))
			require.NoError(t, err)
			require.True(t, ok)

			require.Equal(t, tt.want, got)
		})
	}
}

func TestRegisterResourceOutputs(t *testing.T) {
	t.Parallel()

	tests := []struct {
		desc string
		give khulnasoft.Input
		// TODO: find a way to intercept RegisterResourceOutput calls.
	}{
		{"pu.String", khulnasoft.String("a")},
		{"pu.StringOutput", khulnasoft.String("b").ToStringOutput()},
		{"pux.Output[string]", khulnasoftx.Val("c")},
		{"pux.Output[string]/untyped", khulnasoftx.Val("c").Untyped()},
		{"pux.PtrOf", khulnasoftx.PtrOf[string](khulnasoft.String("d"))},
		{"pu.StringPtr", khulnasoft.StringPtr("e")},
		{"pu.StringPtrOutput", khulnasoft.String("f").ToStringPtrOutput()},
		{"pux.Output[*string]", khulnasoftx.Ptr("g")},
		{"pux.Output[*string]/untyped", khulnasoftx.Ptr("h").Untyped()},
		{"pux.Output[[]int]", khulnasoftx.Val([]int{1, 2, 3})},
		{
			"pux.ArrayOutput[int]",
			khulnasoftx.Cast[khulnasoftx.ArrayOutput[int], []int](
				khulnasoftx.Val([]int{4, 5, 6}),
			),
		},
		{
			"pux.GArrayOutput",
			khulnasoftx.Cast[khulnasoftx.GArrayOutput[int, khulnasoft.IntOutput], []int](
				khulnasoftx.Val([]int{7, 8, 9}),
			),
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.desc, func(t *testing.T) {
			t.Parallel()

			err := khulnasoft.RunErr(func(ctx *khulnasoft.Context) error {
				ctx.Export("x", tt.give)
				return nil
			}, khulnasoft.WithMocks("project", "stack", &mockResourceMonitor{}))
			require.NoError(t, err)
		})
	}
}

type mockResourceMonitor struct {
	CallF        func(khulnasoft.MockCallArgs) (resource.PropertyMap, error)
	NewResourceF func(khulnasoft.MockResourceArgs) (string, resource.PropertyMap, error)
}

var _ khulnasoft.MockResourceMonitor = (*mockResourceMonitor)(nil)

func (m *mockResourceMonitor) Call(args khulnasoft.MockCallArgs) (resource.PropertyMap, error) {
	return m.CallF(args)
}

func (m *mockResourceMonitor) NewResource(args khulnasoft.MockResourceArgs) (string, resource.PropertyMap, error) {
	return m.NewResourceF(args)
}
