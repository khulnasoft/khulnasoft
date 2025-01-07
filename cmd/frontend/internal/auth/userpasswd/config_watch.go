package userpasswd

import (
	"github.com/khulnasoft/khulnasoft/cmd/frontend/internal/auth/providers"
	"github.com/khulnasoft/khulnasoft/internal/conf"
)

// Watch for configuration changes related to the builtin auth provider.
func Init() {
	go func() {
		conf.Watch(func() {
			newPC, _ := GetProviderConfig()
			if newPC == nil {
				providers.Update("builtin", nil)
				return
			}
			providers.Update("builtin", []providers.Provider{&provider{c: newPC}})
		})
	}()
}
