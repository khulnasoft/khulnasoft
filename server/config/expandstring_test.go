// Copyright 2020 - KhulnaSoft Authors <admin@khulnasoft.com>
// SPDX-License-Identifier: Apache-2.0

// +build !windows

package config

import (
	"os"
	"testing"
)

func TestExpandString(t *testing.T) {
	defer os.Setenv("FOOBAR", os.Getenv("FOOBAR"))
	os.Setenv("FOOBAR", "test-value")

	result := ExpandString("my-$FOOBAR")
	if result != "my-test-value" {
		t.Errorf("Unexpected value %v", result)
	}
}
