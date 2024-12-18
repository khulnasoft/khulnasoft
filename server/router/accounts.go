// Copyright 2020 - KhulnaSoft Authors <admin@khulnasoft.com>
// SPDX-License-Identifier: Apache-2.0

package router

import (
	"errors"
	"fmt"
	"html"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/khulnasoft/khulnasoft/server/persistence"
)

func (rt *router) getAccount(c *gin.Context) {
	accountID := c.Param("accountID")
	if l := <-rt.getLimiter().LinearThrottle(time.Second, fmt.Sprintf("getAccount-%s", accountID)); l.Error != nil {
		newJSONError(
			fmt.Errorf("router: error rate limiting request: %w", l.Error),
			http.StatusTooManyRequests,
		).Pipe(c)
		return
	}
	accountUser, ok := c.Value(contextKeyAuth).(persistence.LoginResult)
	if !ok {
		newJSONError(
			errors.New("router: could not find account user object in request context"),
			http.StatusNotFound,
		).Pipe(c)
		return
	}

	if ok := accountUser.CanAccessAccount(accountID); !ok {
		newJSONError(
			fmt.Errorf("router: account user does not have permissions to access account %s", accountID),
			http.StatusForbidden,
		).Pipe(c)
		return
	}

	result, err := rt.db.GetAccount(accountID, true, true, c.Query("since"))
	if err != nil {
		var errUnknown persistence.ErrUnknownAccount
		if errors.As(err, &errUnknown) {
			newJSONError(
				fmt.Errorf("router: account %s not found", accountID),
				http.StatusNotFound,
			).Pipe(c)
			return
		}
		newJSONError(
			fmt.Errorf("router: error looking up account: %w", err),
			http.StatusInternalServerError,
		).Pipe(c)
		return
	}
	result.RetentionPeriod = rt.config.App.Retention.String()
	c.JSON(http.StatusOK, result)
}

func (rt *router) deleteAccount(c *gin.Context) {
	accountID := c.Param("accountID")

	if l := <-rt.getLimiter().LinearThrottle(time.Second, fmt.Sprintf("deleteAccount-%s", accountID)); l.Error != nil {
		newJSONError(
			fmt.Errorf("router: error rate limiting request: %w", l.Error),
			http.StatusTooManyRequests,
		).Pipe(c)
		return
	}

	accountUser, ok := c.Value(contextKeyAuth).(persistence.LoginResult)
	if !ok {
		newJSONError(
			errors.New("router: could not find account user object in request context"),
			http.StatusNotFound,
		).Pipe(c)
		return
	}

	if ok := accountUser.CanAccessAccount(accountID) && accountUser.IsSuperAdmin(); !ok {
		newJSONError(
			fmt.Errorf("router: account user does not have permissions to delete account %s", accountID),
			http.StatusForbidden,
		).Pipe(c)
		return
	}

	err := rt.db.RetireAccount(accountID)
	if err != nil {
		var errUnknown persistence.ErrUnknownAccount
		if errors.As(err, &errUnknown) {
			newJSONError(
				fmt.Errorf("router: account %s not found", accountID),
				http.StatusNotFound,
			).Pipe(c)
			return
		}
		newJSONError(
			fmt.Errorf("router: error deleting account: %w", err),
			http.StatusInternalServerError,
		).Pipe(c)
		return
	}
	c.Status(http.StatusNoContent)
}

type createAccountRequest struct {
	AccountName  string `json:"accountName"`
	EmailAddress string `json:"emailAddress"`
	Password     string `json:"password"`
}

func (rt *router) postAccount(c *gin.Context) {
	accountUser, ok := c.Value(contextKeyAuth).(persistence.LoginResult)
	if !ok {
		newJSONError(
			errors.New("router: could not find account user object in request context"),
			http.StatusUnauthorized,
		).Pipe(c)
		return
	}

	var req createAccountRequest
	if err := c.BindJSON(&req); err != nil {
		newJSONError(
			fmt.Errorf("router: error decoding response body: %w", err),
			http.StatusBadRequest,
		).Pipe(c)
		return
	}

	if l := <-rt.getLimiter().LinearThrottle(time.Second*5, fmt.Sprintf("postAccount-%s", req.EmailAddress)); l.Error != nil {
		newJSONError(
			fmt.Errorf("router: error rate limiting request: %w", l.Error),
			http.StatusTooManyRequests,
		).Pipe(c)
		return
	}

	accountInRequest, err := rt.db.Login(req.EmailAddress, req.Password)
	if err != nil {
		newJSONError(
			fmt.Errorf("router: error validating given credentials: %w", err),
			http.StatusUnauthorized,
		).Pipe(c)
		return
	}
	// the given credentials might be valid, but belong to a different user
	// than the one who is calling this
	if accountInRequest.AccountUserID != accountUser.AccountUserID {
		newJSONError(
			fmt.Errorf("router: given credentials belong to user other than requester with id %s", accountUser.AccountUserID),
			http.StatusBadRequest,
		).Pipe(c)
		return
	}

	if ok := accountUser.IsSuperAdmin(); !ok {
		newJSONError(
			errors.New("router: account user does not have permissions to create account"),
			http.StatusForbidden,
		).Pipe(c)
		return
	}

	if err := rt.db.CreateAccount(html.UnescapeString(rt.sanitizer.Sanitize(req.AccountName)), req.EmailAddress, req.Password); err != nil {
		newJSONError(
			fmt.Errorf("router: error creating account %s: %w", req.AccountName, err),
			http.StatusInternalServerError,
		).Pipe(c)
		return
	}
	c.JSON(http.StatusCreated, nil)
}
