package main

import (
	"example.com/khulnasoft-pkg/sdk/go/pkg"
	"github.com/khulnasoft/khulnasoft/sdk/v3/go/khulnasoft"
)

func main() {
	khulnasoft.Run(func(ctx *khulnasoft.Context) error {
		_, err := pkg.NewRandom(ctx, "random", &pkg.RandomArgs{
			Length: khulnasoft.Int(8),
		})
		if err != nil {
			return err
		}

		hello := "hello"
		_, err = pkg.DoEcho(ctx, &pkg.DoEchoArgs{
			Echo: &hello,
		})
		if err != nil {
			return err
		}

		_ = pkg.DoEchoOutput(ctx, pkg.DoEchoOutputArgs{
			Echo: khulnasoft.String("hello"),
		})

		p, err := pkg.NewEcho(ctx, "echo", &pkg.EchoArgs{})
		if err != nil {
			return err
		}

		_, err = p.DoEchoMethod(ctx, &pkg.EchoDoEchoMethodArgs{
			Echo: khulnasoft.String("hello"),
		})
		return err
	})
}
