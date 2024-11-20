package main

import (
	"git.example.org/khulnasoft-synthetic/resourceProperties"
	"github.com/khulnasoft/khulnasoft/sdk/v3/go/khulnasoft"
)

func main() {
	khulnasoft.Run(func(ctx *khulnasoft.Context) error {
		rt, err := resourceProperties.NewRoot(ctx, "rt", nil)
		if err != nil {
			return err
		}
		ctx.Export("trivial", rt)
		ctx.Export("simple", rt.Res1)
		ctx.Export("foo", rt.Res1.ApplyT(func(res1 *resourceproperties.Res1) (resourceproperties.Obj2, error) {
			return res1.Obj1.Res2.Obj2, nil
		}).(resourceproperties.Obj2Output))
		ctx.Export("complex", rt.Res1.ApplyT(func(res1 *resourceproperties.Res1) (*float64, error) {
			return &res1.Obj1.Res2.Obj2.Answer, nil
		}).(khulnasoft.Float64PtrOutput))
		return nil
	})
}
