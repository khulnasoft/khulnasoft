package auth

import (
	"context"
	"net/url"

	"github.com/khulnasoft/khulnasoft/internal/encryption/keyring"
	"github.com/khulnasoft/khulnasoft/internal/extsvc/auth"
	"github.com/khulnasoft/khulnasoft/internal/extsvc/github"
	ghstore "github.com/khulnasoft/khulnasoft/internal/github_apps/store"
	"github.com/khulnasoft/khulnasoft/internal/types"
	"github.com/khulnasoft/khulnasoft/lib/errors"
)

type CreateAuthenticatorForCredentialOpts struct {
	Repo           *types.Repo
	GitHubAppStore ghstore.GitHubAppsStore
}

func CreateAuthenticatorForCredential(ctx context.Context, ghAppID int, opts CreateAuthenticatorForCredentialOpts) (auth.Authenticator, error) {
	var authenticator auth.Authenticator

	ghApp, err := opts.GitHubAppStore.GetByID(ctx, ghAppID)
	if err != nil {
		return nil, err
	}

	authenticator, err = NewGitHubAppAuthenticator(ghApp.AppID, []byte(ghApp.PrivateKey))
	if err != nil {
		return nil, err
	}

	if opts.Repo != nil {
		baseURL, err := url.Parse(opts.Repo.ExternalRepo.ServiceID)
		if err != nil {
			return nil, err
		}

		md, ok := opts.Repo.Metadata.(*github.Repository)
		if !ok {
			return nil, errors.Newf("expected repo metadata to be a github.Repository, but got %T", opts.Repo.Metadata)
		}

		owner, _, err := github.SplitRepositoryNameWithOwner(md.NameWithOwner)
		if err != nil {
			return nil, err
		}
		installID, err := opts.GitHubAppStore.GetInstallID(ctx, ghApp.AppID, owner)
		if err != nil {
			return nil, err
		}
		authenticator = NewInstallationAccessToken(baseURL, installID, authenticator, keyring.Default().GitHubAppKey)
	}
	return authenticator, nil
}
