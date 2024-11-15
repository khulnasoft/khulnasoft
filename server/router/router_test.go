// Copyright 2020 - KhulnaSoft Authors <admin@khulnasoft.com>
// SPDX-License-Identifier: Apache-2.0

package router

import (
	"html/template"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/khulnasoft/khulnasoft/server/config"
	"github.com/khulnasoft/khulnasoft/server/persistence"
)

type mockDatabase struct {
	persistence.Service
}

func TestMain(m *testing.M) {
	gin.SetMode(gin.ReleaseMode)
	os.Exit(m.Run())
}

func TestNew(t *testing.T) {
	New(
		WithDatabase(&mockDatabase{}),
		WithConfig(&config.Config{}),
		WithTemplate(template.New("a test")),
	)
}
