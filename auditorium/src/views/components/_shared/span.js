/**
 * Copyright 2020 - KhulnaSoft Authors <admin@khulnasoft.com>
 * SPDX-License-Identifier: Apache-2.0
 */

/** @jsx h */
const { h } = require('preact')

const Span = (props) => {
  const { children, ...otherProps } = props
  return (
    <span
      {...otherProps}
      dangerouslySetInnerHTML={{ __html: children }}
    />
  )
}

module.exports = Span
