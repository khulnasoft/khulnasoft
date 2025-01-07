package main

import (
	"github.com/khulnasoft/khulnasoft/cmd/enterprise-portal/service"
	"github.com/khulnasoft/khulnasoft/lib/managedservicesplatform/runtime"
)

func main() {
	runtime.Start(&service.Service{})
}
