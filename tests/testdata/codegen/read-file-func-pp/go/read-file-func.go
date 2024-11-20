package main

import (
	"os"

	"github.com/khulnasoft/khulnasoft/sdk/v3/go/khulnasoft"
)

func readFileOrPanic(path string) khulnasoft.StringPtrInput {
	data, err := os.ReadFile(path)
	if err != nil {
		panic(err.Error())
	}
	return khulnasoft.String(string(data))
}

func main() {
	khulnasoft.Run(func(ctx *khulnasoft.Context) error {
		key := readFileOrPanic("key.pub")
		ctx.Export("result", key)
		return nil
	})
}
