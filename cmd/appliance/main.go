package main

import (
	"github.com/khulnasoft/khulnasoft/cmd/appliance/shared"
	"github.com/khulnasoft/khulnasoft/internal/sanitycheck"
	"github.com/khulnasoft/khulnasoft/internal/service/svcmain"
)

func main() {
	sanitycheck.Pass()
	svcmain.SingleServiceMainWithoutConf(shared.Service, nil, svcmain.OutOfBandConfiguration{})
}
