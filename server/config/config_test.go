// Copyright 2020 - KhulnaSoft Authors <admin@khulnasoft.com>
// SPDX-License-Identifier: Apache-2.0

package config

import (
	"os"
	"testing"
)

func TestNew(t *testing.T) {
	defer os.Setenv("KHULNASOFT_APP_DEPLOYTARGET", os.Getenv("KHULNASOFT_APP_DEPLOYTARGET"))
	os.Setenv("KHULNASOFT_APP_DEPLOYTARGET", "heroku")
	defer os.Setenv("PORT", os.Getenv("PORT"))
	os.Setenv("PORT", "9876")

	c, err := New(false, "./testdata/khulnasoft.env")
	if err != nil {
		t.Errorf("Unexpected error %v", err)
	}

	if c.Server.Port != 9876 {
		t.Errorf("Unexpected port value %v", c.Server.Port)
	}

	if c.Server.AutoTLS[0] != "analytics.khulnasoft.com" {
		t.Errorf("Unexpected AutoTLS config %v", c.Server.AutoTLS)
	}

	if c.Secret == nil {
		t.Error("Expected app secret to be populated")
	}
}
