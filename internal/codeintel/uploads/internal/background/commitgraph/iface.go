package commitgraph

import (
	"context"

	"github.com/khulnasoft/khulnasoft/internal/database/locker"
)

type Locker interface {
	Lock(ctx context.Context, key int32, blocking bool) (bool, locker.UnlockFunc, error)
}
