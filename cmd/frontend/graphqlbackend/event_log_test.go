package graphqlbackend

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/khulnasoft/khulnasoft/internal/conf"
	"github.com/khulnasoft/khulnasoft/internal/database"
	"github.com/khulnasoft/khulnasoft/schema"
)

func TestUserEventLogResolver_URL(t *testing.T) {
	conf.Mock(
		&conf.Unified{
			SiteConfiguration: schema.SiteConfiguration{
				ExternalURL: "https://sourcegraph.test:3443",
			},
		},
	)
	defer conf.Mock(nil)

	tests := []struct {
		name string
		url  string
		want string
	}{
		{
			name: "valid URL",
			url:  "https://sourcegraph.test:3443/search",
			want: "https://sourcegraph.test:3443/search",
		},
		{
			name: "invalid URL",
			url:  "https://localhost:3080/search",
			want: "",
		},
		{
			name: "not a URL",
			url:  `javascript:alert("HIJACKED")`,
			want: "",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := (&userEventLogResolver{
				event: &database.Event{
					URL: test.url,
				},
			}).URL()
			assert.Equal(t, test.want, got)
		})
	}
}
