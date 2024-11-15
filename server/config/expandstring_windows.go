// Copyright 2020 - KhulnaSoft Authors <admin@khulnasoft.com>
// SPDX-License-Identifier: Apache-2.0

// +build windows

package config

import (
	"golang.org/x/sys/windows/registry"
)

// ExpandString expands all environment variables in the given string
func ExpandString(s string) string {
	r, _ := registry.ExpandString(s)
	return r
}
