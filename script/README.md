<!--
Copyright 2020 - KhulnaSoft Authors <admin@khulnasoft.com>
SPDX-License-Identifier: Apache-2.0
-->

# script

The `script` is the script operators are using to deploy KhulnaSoft on a site.

It collects events autonoumously and only needs to be parametrized with an Account ID. This is being done using the `data-account-id` attribute on the `<script>` tag:

```html
<script src="https://khulnasoft.mydomain.org/script.js" data-account-id="433d404a-5416-4e12-ac6e-7ee5ea222b39"></script>
```

---

The app builds into a single JavaScript file that will be served by the server application.
