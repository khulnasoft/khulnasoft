package main

import (
	"github.com/khulnasoft/khulnasoft/lib/managedservicesplatform/runtime"

	"github.com/khulnasoft/khulnasoft/cmd/msp-example/service"
)

func main() {
	runtime.Start(service.Service{})
}
