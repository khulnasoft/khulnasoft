/**
 * Copyright 2020 - KhulnaSoft Authors <admin@khulnasoft.com>
 * SPDX-License-Identifier: Apache-2.0
 */

const { route } = require('preact-router')

module.exports = (store) => (next) => (action) => {
  switch (action.type) {
    case 'EXPRESS_CONSENT_SUCCESS':
      route(window.location.pathname)
      next(action)
      break
    default:
      next(action)
  }
}
