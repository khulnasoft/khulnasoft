/**
 * Copyright 2020 - KhulnaSoft Authors <admin@khulnasoft.com>
 * SPDX-License-Identifier: Apache-2.0
 */

/** @jsx h */
const { h } = require('preact')
const { forwardRef } = require('preact/compat')
const { useState } = require('preact/hooks')

const LabeledInput = require('./../_shared/labeled-input')
const SubmitButton = require('./../_shared/submit-button')
const ForwardingAnchor = require('./../_shared/forwarding-anchor')

const Form = forwardRef((props, ref) => {
  const [isDisabled, setIsDisabled] = useState(false)
  function handleSubmit (e) {
    e.preventDefault()
    var formData = new window.FormData(e.currentTarget)
    setIsDisabled(true)
    props.onLogin(
      formData.get('username'),
      formData.get('password'),
      __('Could not log in using the given credentials. Try again.')
    )
      .then(() => setIsDisabled(false))
  }

  return (
    <div class='pa3 bg-black-05'>
      <h4 class='f4 normal mt0 mb3'>
        {__('Log in as operator')}
      </h4>
      <form data-testid='login/form' class='mw6 center' onsubmit={handleSubmit}>
        <LabeledInput
          type='email'
          name='username'
          ref={ref}
          required
          disabled={isDisabled}
          data-testid='login/username-input'
        >
          {__('Email address')}
        </LabeledInput>
        <LabeledInput
          type='password'
          name='password'
          required
          disabled={isDisabled}
          data-testid='login/password-input'
        >
          {__('Password')}
        </LabeledInput>
        <SubmitButton disabled={isDisabled}>
          {__('Log in')}
        </SubmitButton>
        <div class='mb3'>
          <ForwardingAnchor pick={['locale']} class='b link dim dark-green' href='/forgot-password/'>
            {__('Forgot password?')}
          </ForwardingAnchor>
        </div>
      </form>
    </div>
  )
})

module.exports = Form
