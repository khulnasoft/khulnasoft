package config

import (
	"gopkg.in/yaml.v2"

	"github.com/khulnasoft/khulnasoft/lib/errors"
)

func UnmarshalYAML(data []byte) (AutoIndexJobSpecList, error) {
	configuration := AutoIndexJobSpecList{}
	if err := yaml.Unmarshal(data, &configuration); err != nil {
		return AutoIndexJobSpecList{}, errors.Errorf("invalid YAML: %v", err)
	}

	return configuration, nil
}
