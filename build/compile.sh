#!/bin/sh
# Copyright 2021 - KhulnaSoft Authors <admin@khulnasoft.com>
# SPDX-License-Identifier: Apache-2.0

if [ -z "$LDFLAGS" ]; then
  xgo \
    --targets=$TARGETS \
    --tags 'osusergo netgo static_build sqlite_omit_load_extension' \
    --ldflags="-s -w -X github.com/khulnasoft/khulnasoft/server/config.Revision=$GIT_REVISION" \
    github.com/khulnasoft/khulnasoft/server/cmd/khulnasoft
else
  xgo \
    --targets=$TARGETS \
    --tags 'osusergo netgo static_build sqlite_omit_load_extension' \
    --ldflags="-linkmode external -extldflags '$LDFLAGS' -s -w -X github.com/khulnasoft/khulnasoft/server/config.Revision=$GIT_REVISION" \
    github.com/khulnasoft/khulnasoft/server/cmd/khulnasoft
fi
