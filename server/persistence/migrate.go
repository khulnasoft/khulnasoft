// Copyright 2020 - KhulnaSoft Authors <admin@khulnasoft.com>
// SPDX-License-Identifier: Apache-2.0

package persistence

// Migrate runs the defined database migrations in the given db or initializes it
// from the latest definition if it is still blank.
func (p *persistenceLayer) Migrate() error {
	return p.dal.ApplyMigrations()
}
