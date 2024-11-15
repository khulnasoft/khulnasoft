/**
 * Copyright 2020 - KhulnaSoft Authors <admin@khulnasoft.com>
 * SPDX-License-Identifier: Apache-2.0
 */

var enqueueCryptoTests = require('./web-crypto.test')
var forgeCrypto = require('./forge-crypto')

enqueueCryptoTests(forgeCrypto, 'src/forge-crypto.js')
