package licensing

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProductNameWithBrand(t *testing.T) {
	tests := []struct {
		licenseTags []string
		want        string
	}{
		{licenseTags: GetFreeLicenseInfo().Tags, want: "Khulnasoft Free"},
		{licenseTags: nil, want: "Khulnasoft Enterprise"},
		{licenseTags: []string{}, want: "Khulnasoft Enterprise"},
		{licenseTags: []string{"x"}, want: "Khulnasoft Enterprise"}, // unrecognized tag "x" is ignored
		{licenseTags: []string{"trial"}, want: "Khulnasoft Enterprise (trial)"},
		{licenseTags: []string{"dev"}, want: "Khulnasoft Enterprise (dev use only)"},
		{licenseTags: []string{"trial", "dev"}, want: "Khulnasoft Enterprise (trial, dev use only)"},
		{licenseTags: []string{"internal"}, want: "Khulnasoft Enterprise (internal use only)"},

		{licenseTags: []string{"plan:team-0"}, want: "Khulnasoft Team"},
		{licenseTags: []string{"plan:team-0", "trial"}, want: "Khulnasoft Team (trial)"},
		{licenseTags: []string{"plan:team-0", "dev"}, want: "Khulnasoft Team (dev use only)"},
		{licenseTags: []string{"plan:team-0", "dev", "trial"}, want: "Khulnasoft Team (trial, dev use only)"},
		{licenseTags: []string{"plan:team-0", "internal"}, want: "Khulnasoft Team (internal use only)"},

		{licenseTags: []string{"plan:enterprise-0"}, want: "Khulnasoft Enterprise"},
		{licenseTags: []string{"plan:enterprise-0", "trial"}, want: "Khulnasoft Enterprise (trial)"},
		{licenseTags: []string{"plan:enterprise-0", "dev"}, want: "Khulnasoft Enterprise (dev use only)"},
		{licenseTags: []string{"plan:enterprise-0", "dev", "trial"}, want: "Khulnasoft Enterprise (trial, dev use only)"},
		{licenseTags: []string{"plan:enterprise-0", "internal"}, want: "Khulnasoft Enterprise (internal use only)"},

		{licenseTags: []string{"plan:enterprise-1"}, want: "Code Search Enterprise"},
		{licenseTags: []string{"plan:enterprise-1", "trial"}, want: "Code Search Enterprise (trial)"},
		{licenseTags: []string{"plan:enterprise-1", "dev"}, want: "Code Search Enterprise (dev use only)"},
		{licenseTags: []string{"plan:enterprise-1", "dev", "trial"}, want: "Code Search Enterprise (trial, dev use only)"},
		{licenseTags: []string{"plan:enterprise-1", "internal"}, want: "Code Search Enterprise (internal use only)"},
	}
	for _, test := range tests {
		t.Run(fmt.Sprintf("licenseTags=%v", test.licenseTags), func(t *testing.T) {
			assert.Equal(t, test.want, ProductNameWithBrand(test.licenseTags))
		})
	}
}
