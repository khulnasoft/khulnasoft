/**
 * Copyright 2020-2021 - KhulnaSoft Authors <admin@khulnasoft.com>
 * SPDX-License-Identifier: Apache-2.0
 */

exports.serialize = serialize

function serialize (obj) {
  return Object.keys(obj)
    .map(function (key) {
      if (obj[key] === true) {
        return key
      }
      if (obj[key] === false) {
        return null
      }
      return key + '=' + obj[key]
    })
    .filter(Boolean)
    .join('; ')
}

exports.parse = parse

function parse (cookieString) {
  return cookieString.split(';')
    .reduce(function (acc, pair) {
      var chunks = pair.trim().split('=')
      acc[chunks[0]] = chunks[1]
      return acc
    }, {})
}

exports.defaultCookie = defaultCookie

function defaultCookie (name, value, opts) {
  opts = opts || {}
  var isLocalhost = window.location.hostname === 'localhost'
  var cookie = {}
  cookie[name] = value
  cookie.expires = opts.expires ? opts.expires.toUTCString() : false
  return Object.assign(cookie, {
    // it is important not to lock this cookie down to `/vault` as the
    // server checks for it before accepting events
    Path: opts.path || '/',
    SameSite: isLocalhost ? 'Lax' : 'None',
    Secure: !isLocalhost
  })
}
