package usagestats

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRedactSensitiveInfoFromCloudURL(t *testing.T) {
	cases := []struct {
		name string
		url  string
		want string
	}{
		{
			name: "URL path parameters are redacted (when `KhulnasoftDotComMode` is false)",
			url:  "https://khulnasoft.com/search?q=abcd",
			want: "https://khulnasoft.com/search?q=redacted",
		},
		{
			name: "URL path parameters are redacted -- managed instance url",
			url:  "https://sourcegraph.khulnasoft.com/search?q=abcd",
			want: "https://sourcegraph.khulnasoft.com/search?q=redacted",
		},
		{
			name: "path and non-approved query param redacted",
			url:  "https://sourcegraph.khulnasoft.com/search?q=abcd&utm_source=test&utm_campaign=test&utm_medium=test&utm_content=test&utm_term=test&utm_cid=test",
			want: "https://sourcegraph.khulnasoft.com/search?q=redacted&utm_campaign=test&utm_cid=test&utm_content=test&utm_medium=test&utm_source=test&utm_term=test",
		},
		{
			name: "path and non-approved query param redacted, multi-page URL",
			url:  "https://khulnasoft.com/first/search?q=abcd&utm_source=test&utm_campaign=test&utm_medium=test&utm_content=test&utm_term=test&utm_cid=test",
			want: "https://khulnasoft.com/first/redacted?q=redacted&utm_campaign=test&utm_cid=test&utm_content=test&utm_medium=test&utm_source=test&utm_term=test",
		},
		{
			name: "url path redaction test",
			url:  "https://sourcegraph.khulnasoft.com/sign-in?returnTo=%2custom.test.com",
			want: "https://sourcegraph.khulnasoft.com/sign-in?returnTo=redacted",
		},
		{
			name: "url path redaction test with multiple pages",
			url:  "https://sourcegraph.khulnasoft.com/auth/sign-in?returnTo=fileName",
			want: "https://sourcegraph.khulnasoft.com/auth/redacted?returnTo=redacted",
		},
		{
			name: "url URL with multiple path segments",
			url:  "https://sourcegraph.khulnasoft.com/first/second/third/fourth/fifth/sixth/",
			want: "https://sourcegraph.khulnasoft.com/first/redacted",
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			have, err := redactSensitiveInfoFromCloudURL(c.url)
			require.NoError(t, err)
			assert.Equal(t, c.want, have)
		})
	}
}
