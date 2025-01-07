package streaming

import (
	"github.com/khulnasoft/khulnasoft/internal/compute"
	"github.com/khulnasoft/khulnasoft/internal/search/streaming"
)

type Event struct {
	Results []compute.Result // TODO(rvantonder): hydrate repo information in this Event type.
	Stats   streaming.Stats
}
