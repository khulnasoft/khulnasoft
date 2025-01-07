package main

import (
	"github.com/khulnasoft/khulnasoft/cmd/pings/service"
	"github.com/khulnasoft/khulnasoft/lib/managedservicesplatform/runtime"
)

func main() {
	runtime.Start[service.Config](&service.Service{})
}
