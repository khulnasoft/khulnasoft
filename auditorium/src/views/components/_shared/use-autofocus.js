/**
 * Copyright 2020 - KhulnaSoft Authors <admin@khulnasoft.com>
 * SPDX-License-Identifier: Apache-2.0
 */

const { useRef, useEffect } = require('preact/hooks')

const useAutofocus = () => {
  const autofocus = useRef(null)
  useEffect(() => {
    if (autofocus.current) {
      autofocus.current.focus()
    }
  }, [autofocus])
  return autofocus
}

module.exports = useAutofocus
