/**
 * Copyright 2020 - KhulnaSoft Authors <admin@khulnasoft.com>
 * SPDX-License-Identifier: Apache-2.0
 */

/** @jsx h */
const { h } = require('preact')

const AccountPicker = require('./../_shared/account-picker')

module.exports = (props) => {
  return (
    <AccountPicker
      headline={__('Open account')}
      {...props}
    />
  )
}
