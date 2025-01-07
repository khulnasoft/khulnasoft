package graphqlbackend

import (
	"github.com/khulnasoft/khulnasoft/cmd/frontend/internal/backend"
)

func resetMocks() {
	backend.Mocks = backend.MockServices{}
}
