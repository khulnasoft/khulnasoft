#!/bin/bash
# Copyright 2020-2021 - KhulnaSoft Authors <admin@khulnasoft.com>
# SPDX-License-Identifier: Apache-2.0

set -eo pipefail

go run cmd/extract-strings/main.go public/static/*.go.html \
  | xgettext --omit-header --color=never -o - -c~ --no-location --language=python -
