#!/usr/bin/env bash

set -euo pipefail

khulnasoft version

time khulnasoft destroy --yes

khulnasoft config set mode new
time khulnasoft up --yes --skip-preview

khulnasoft config set mode alias
time khulnasoft up --yes --skip-preview


export PATH=~/.khulnasoft-dev/bin:$PATH

khulnasoft version

time khulnasoft destroy --yes

khulnasoft config set mode new
time khulnasoft up --yes --skip-preview

khulnasoft config set mode alias
time khulnasoft up --yes --skip-preview
