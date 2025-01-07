package global

import (
	"testing"

	bt "github.com/khulnasoft/khulnasoft/internal/batches/testing"
	btypes "github.com/khulnasoft/khulnasoft/internal/batches/types"
	"github.com/khulnasoft/khulnasoft/internal/conf"
	"github.com/khulnasoft/khulnasoft/schema"
)

func TestDefaultReconcilerEnqueueState(t *testing.T) {
	t.Run("no windows", func(t *testing.T) {
		bt.MockConfig(t, &conf.Unified{})

		have := DefaultReconcilerEnqueueState()
		want := btypes.ReconcilerStateQueued
		if have != want {
			t.Errorf("unexpected default state: have=%v want=%v", have, want)
		}
	})

	t.Run("windows", func(t *testing.T) {
		bt.MockConfig(t, &conf.Unified{
			SiteConfiguration: schema.SiteConfiguration{
				BatchChangesRolloutWindows: &[]*schema.BatchChangeRolloutWindow{
					{Rate: "unlimited"},
				},
			},
		})

		have := DefaultReconcilerEnqueueState()
		want := btypes.ReconcilerStateScheduled
		if have != want {
			t.Errorf("unexpected default state: have=%v want=%v", have, want)
		}
	})
}
