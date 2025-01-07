package main

import (
	"github.com/khulnasoft/khulnasoft/dev/linearhooks/internal/service"
	"github.com/khulnasoft/khulnasoft/lib/managedservicesplatform/runtime"
)

func main() {
	runtime.Start[service.Config](service.Service{})
}
