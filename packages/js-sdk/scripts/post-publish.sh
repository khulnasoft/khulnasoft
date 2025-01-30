#!/usr/bin/env bash

npm pkg set 'name'='@khulnasoft/sdk'
npm publish --no-git-checks
npm pkg set 'name'='khulnasoft'
npm deprecate "@khulnasoft/sdk@$(npm pkg get version | tr -d \")" "The package @khulnasoft/sdk has been renamed to khulnasoft. Please uninstall the old one and install the new by running following command: npm uninstall @khulnasoft/sdk && npm install khulnasoft"
