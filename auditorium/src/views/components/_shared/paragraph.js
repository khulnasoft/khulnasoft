/**
 * Copyright 2020 - KhulnaSoft Authors <admin@khulnasoft.com>
 * SPDX-License-Identifier: Apache-2.0
 */

/** @jsx h */
const { h } = require('preact')
const _ = require('underscore')

const Paragraph = (props) => {
  const { children, ...otherProps } = props
  if (!_.isString(children)) {
    return (
      <p
        {...otherProps}
      >
        {children}
      </p>
    )
  }
  return (
    <p
      {...otherProps}
      dangerouslySetInnerHTML={{ __html: children }}
    />
  )
}

module.exports = Paragraph
