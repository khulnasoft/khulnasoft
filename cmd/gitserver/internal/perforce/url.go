package perforce

import (
	"github.com/khulnasoft/khulnasoft/internal/vcs"
	"github.com/khulnasoft/khulnasoft/lib/errors"
)

// DecomposePerforceRemoteURL decomposes information back from a clone URL for a
// Perforce depot.
func DecomposePerforceRemoteURL(remoteURL *vcs.URL) (username, password, host, depot string, err error) {
	if remoteURL.Scheme != "perforce" {
		return "", "", "", "", errors.New(`scheme is not "perforce"`)
	}
	password, _ = remoteURL.User.Password()
	return remoteURL.User.Username(), password, remoteURL.Host, remoteURL.Path, nil
}
