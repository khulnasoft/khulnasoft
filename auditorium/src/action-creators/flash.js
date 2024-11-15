/**
 * Copyright 2020 - KhulnaSoft Authors <admin@khulnasoft.com>
 * SPDX-License-Identifier: Apache-2.0
 */

exports.expire = (flashId) => ({
  type: 'EXPIRE_FLASH',
  payload: { flashId }
})
