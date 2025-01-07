package bitbucketcloud

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/khulnasoft/khulnasoft/internal/extsvc/auth"
	bbtest "github.com/khulnasoft/khulnasoft/internal/extsvc/bitbucketcloud/testing"
)

func TestClient_CurrentUser(t *testing.T) {
	// WHEN UPDATING: as long as you provide a valid token, this should work
	// fine.

	ctx := context.Background()
	c := newTestClient(t)

	t.Run("valid token", func(t *testing.T) {
		user, err := c.CurrentUser(ctx)
		assert.Nil(t, err)
		assert.NotNil(t, user)
		assert.NotEmpty(t, user.Username)
		assertGolden(t, user)
	})

	t.Run("invalid token", func(t *testing.T) {
		user, err := c.WithAuthenticator(&auth.BasicAuth{
			Username: bbtest.GetenvTestBitbucketCloudUsername(),
			Password: "this is not a valid password",
		}).CurrentUser(ctx)
		assert.Nil(t, user)
		assert.NotNil(t, err)
	})
}
