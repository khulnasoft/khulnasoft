package main

import (
	"example.com/khulnasoft-splat/sdk/go/splat"
	"github.com/khulnasoft/khulnasoft/sdk/v3/go/khulnasoft"
)

func main() {
	khulnasoft.Run(func(ctx *khulnasoft.Context) error {
		allKeys, err := splat.GetSshKeys(ctx, map[string]interface{}{}, nil)
		if err != nil {
			return err
		}
		var splat0 []string
		for _, val0 := range allKeys.SshKeys {
			splat0 = append(splat0, val0.Name)
		}
		_, err = splat.NewServer(ctx, "main", &splat.ServerArgs{
			SshKeys: splat0,
		})
		if err != nil {
			return err
		}
		return nil
	})
}
