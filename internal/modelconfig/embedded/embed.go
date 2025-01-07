package embedded

import (
	_ "embed"
	"encoding/json"

	"github.com/khulnasoft/khulnasoft/internal/modelconfig/types"
	"github.com/khulnasoft/khulnasoft/lib/errors"
)

// Embed the latest Cody Gateway config data directly into the
// current Go binary. This allows us to version the configuration
// along with the source code and "ship" it without needing to
// have consumers have access to block storage, etc.
//
// See cmd/cody-gateway-config for how this file gets generated.
//
//go:embed models.json
var modelConfigJSON string

func getRawCodyGatewayModelConfig() string {
	return modelConfigJSON
}

// GetCodyGatewayModelConfig returns the latest embedded models document,
// which is dependent on the specific commit the code was built from.
func GetCodyGatewayModelConfig() (*types.ModelConfiguration, error) {
	rawJSON := getRawCodyGatewayModelConfig()

	var modelConfig types.ModelConfiguration
	if err := json.Unmarshal([]byte(rawJSON), &modelConfig); err != nil {
		return nil, errors.Errorf("unmarshalling %d bytes of modelConfig: %w", len(rawJSON), err)
	}

	return &modelConfig, nil
}
