/**
 * Copyright 2021 - KhulnaSoft Authors <admin@khulnasoft.com>
 * SPDX-License-Identifier: Apache-2.0
 */

module.exports = (state = {}, action) => {
  switch (action.type) {
    case 'UDPATE_QUERY_PARAMS':
      return action.payload
    default:
      return state
  }
}
