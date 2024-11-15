/**
 * Copyright 2020 - KhulnaSoft Authors <admin@khulnasoft.com>
 * SPDX-License-Identifier: Apache-2.0
 */

module.exports = (state = null, action) => {
  switch (action.type) {
    case 'UNRECOVERABLE_ERROR':
      return action.payload
    case 'NAVIGATE':
      return null
    default:
      return state
  }
}
