/**
 * Copyright 2020 - KhulnaSoft Authors <admin@khulnasoft.com>
 * SPDX-License-Identifier: Apache-2.0
 */

exports.unrecoverable = (err) => ({
  type: 'UNRECOVERABLE_ERROR',
  payload: {
    message: err.message,
    stack: err.originalStack || err.stack
  }
})

exports.formValidation = (err) => ({
  type: 'FORM_VALIDATION_ERROR',
  payload: {
    flash: err.message
  }
})
