package main

import (
	"os"
	"strconv"

	"github.com/khulnasoft/khulnasoft/cmd/server/shared"
	"github.com/khulnasoft/khulnasoft/internal/sanitycheck"

	"github.com/khulnasoft/khulnasoft/ui/assets"
)

func main() {
	assets.Init()
	sanitycheck.Pass()

	enableEmbeddings, _ := strconv.ParseBool(os.Getenv("SRC_ENABLE_EMBEDDINGS"))
	if enableEmbeddings {
		shared.ProcfileAdditions = append(
			shared.ProcfileAdditions,
			`embeddings: embeddings`,
		)
		shared.SrcProfServices = append(
			shared.SrcProfServices,
			map[string]string{"Name": "embeddings", "Host": "127.0.0.1:6099"},
		)
	}

	shared.Main()
}
