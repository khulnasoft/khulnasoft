#!/bin/bash
# Copyright 2020 - KhulnaSoft Authors <admin@khulnasoft.com>
# SPDX-License-Identifier: Apache-2.0

set -eo pipefail

for pofile in *.po; do
  msguniq $pofile --no-wrap --color=no --output $pofile
done

msgcat *.po --no-wrap --output messages.po

for locale in $(cat ./locales/LINGUAS); do
  touch "./locales/${locale}.po"
  msgmerge "./locales/${locale}.po" messages.po --no-wrap --output "./locales/${locale}.po"
done
