<!--
Copyright 2020 - KhulnaSoft Authors <admin@khulnasoft.com>
SPDX-License-Identifier: Apache-2.0
-->

# packages

`packages` contains JavaScript modules shared across applications. Consumer should use the `file:` scheme to install the package like so:

```json
{
  "dependencies": {
    "khulnasoft": "file:./../packages"
  }
}
```
