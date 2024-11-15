/**
 * Copyright 2020 - KhulnaSoft Authors <admin@khulnasoft.com>
 * SPDX-License-Identifier: Apache-2.0
 */

var assert = require('assert')
var checkSupport = require('./check-support')

describe('src/check-support.js', function () {
  describe('checkSupport(callback)', function () {
    it('passes when run in chromium', function (done) {
      checkSupport(function (err) {
        assert.strictEqual(err, null)
        done()
      })
    })

    it('calls the callback asynchronously', function (done) {
      var called = false
      function callback () {
        called = true
        done()
      }
      checkSupport(callback)
      assert.strictEqual(called, false)
    })
  })
})
