<!--
Copyright 2020 - KhulnaSoft Authors <admin@khulnasoft.com>
SPDX-FileCopyrightText: 2020 KhulnaSoft Authors

SPDX-License-Identifier: Apache-2.0
-->

# server

`server` exposes a HTTP API that can be used to record and query event data as well as static assets.

`server` is a Go application built on top of the [gin-framework][] that uses Go modules for dependency management. The expected Go version is 1.17.

[gin-framework]: https://github.com/gin-gonic/gin

---

The app builds into a single executable that can be used to run the application.
