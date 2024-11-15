/**
 * Copyright 2020 - KhulnaSoft Authors <admin@khulnasoft.com>
 * SPDX-License-Identifier: Apache-2.0
 */

module.exports = (state = null, action) => {
  switch (action.type) {
    case 'QUERY_SUCCESS':
    case 'PURGE_SUCCESS':
      return action.payload
    case 'NAVIGATE':
      return null
    default:
      return state
  }
}
