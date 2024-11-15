// Copyright 2020 - KhulnaSoft Authors <admin@khulnasoft.com>
// SPDX-License-Identifier: Apache-2.0

package mailer

// Mailer is used to send transactional emails
type Mailer interface {
	Send(from, to, subject, body string) error
}
