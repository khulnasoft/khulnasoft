// Copyright 2020 - KhulnaSoft Authors <admin@khulnasoft.com>
// SPDX-License-Identifier: Apache-2.0

package persistence

// CheckHealth returns an error when the database connection is not working.
func (p *persistenceLayer) CheckHealth() error {
	return p.dal.Ping()
}
