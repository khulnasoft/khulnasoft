package main

import (
	"github.com/khulnasoft/khulnasoft/cmd/telemetry-gateway/service"
	"github.com/khulnasoft/khulnasoft/lib/managedservicesplatform/runtime"
)

func main() {
	runtime.Start(&service.Service{})
}
