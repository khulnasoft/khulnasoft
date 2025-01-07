package parser

import (
	"testing"

	"github.com/khulnasoft/khulnasoft/internal/observation"
)

func TestNewOperation(t *testing.T) {
	// tiny test that check for the side-effects of registering. EG if we have
	// duplicate prometheus metrics.
	_ = newOperations(observation.TestContextTB(t))
}
