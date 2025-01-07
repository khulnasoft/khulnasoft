package userpasswd

import (
	"context"
	"net/url"
	"time"

	"github.com/khulnasoft/khulnasoft/cmd/frontend/internal/backend"
	"github.com/khulnasoft/khulnasoft/internal/database"
	"github.com/khulnasoft/khulnasoft/lib/errors"
)

func AttachEmailVerificationToPasswordReset(ctx context.Context, db database.UserEmailsStore, resetURL url.URL, userID int32, email string) (*url.URL, error) {
	code, err := backend.MakeEmailVerificationCode()
	if err != nil {
		return nil, errors.Wrap(err, "make password verification")
	}
	err = db.SetLastVerification(ctx, userID, email, code, time.Now())
	if err != nil {
		return nil, err
	}

	q := resetURL.Query()
	q.Set("emailVerifyCode", code)
	q.Set("email", email)
	resetURL.RawQuery = q.Encode()
	return &resetURL, nil
}
