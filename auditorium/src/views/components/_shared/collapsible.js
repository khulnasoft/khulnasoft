/**
 * Copyright 2020 - KhulnaSoft Authors <admin@khulnasoft.com>
 * SPDX-License-Identifier: Apache-2.0
 */

/** @jsx h */
const { h, Fragment } = require('preact')
const { useState } = require('preact/hooks')

const Collapsible = (props) => {
  const { header, body, initAsCollapsed = true } = props
  const [isCollapsed, setIsCollapsed] = useState(initAsCollapsed)

  const headerContent = header({
    handleToggle: () => setIsCollapsed(!isCollapsed),
    isCollapsed: isCollapsed
  })

  return (
    <Fragment>
      {headerContent}
      {isCollapsed ? null : body({ isCollapsed })}
    </Fragment>
  )
}

module.exports = Collapsible
