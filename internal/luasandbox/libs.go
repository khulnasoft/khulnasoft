package luasandbox

import (
	lua "github.com/yuin/gopher-lua"

	"github.com/khulnasoft/khulnasoft/internal/luasandbox/libs"
	"github.com/khulnasoft/khulnasoft/internal/luasandbox/util"
	"github.com/khulnasoft/khulnasoft/internal/memo"
)

type LuaLib interface {
	LuaAPI() map[string]lua.LGFunction
}

var defaultAPIs = map[string]LuaLib{
	"internal_path": libs.Path,
}

var DefaultGoModules = memo.NewMemoizedConstructor(func() (map[string]lua.LGFunction, error) {
	modules := map[string]lua.LGFunction{}
	for name, api := range defaultAPIs {
		modules[name] = util.CreateModule(api.LuaAPI())
	}

	return modules, nil
})
