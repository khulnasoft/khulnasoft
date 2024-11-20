#!/usr/bin/env bash

set -euo pipefail

# Populates ./bin directories for use by goreleaser.
rm -rf ./bin

LOCAL="${1:-"false"}"

COMMIT_TIME=$(git log -n1 --pretty='format:%cd' --date=format:'%Y%m%d%H%M')

install_file () {
    src="$1"
    shift

    for OS in "$@"; do # for each argument after the first:
        # if LOCAL == "true" and `go env goos` is not equal to the OS, skip it
        if [ "${LOCAL}" = "local" ] && [ "$(go env GOOS)" != "${OS}" ]; then
            continue
        fi
        DESTDIR="bin/${OS}"
        mkdir -p "${DESTDIR}"
        dest=$(basename "${src}")
        cp "$src" "${DESTDIR}/${dest}"
        touch -t "${COMMIT_TIME}" "$dest"
    done
}

install_file sdk/nodejs/dist/khulnasoft-analyzer-policy                         linux   darwin
install_file sdk/nodejs/dist/khulnasoft-analyzer-policy.cmd                     windows

install_file sdk/nodejs/dist/khulnasoft-resource-khulnasoft-nodejs                  linux   darwin
install_file sdk/nodejs/dist/khulnasoft-resource-khulnasoft-nodejs.cmd              windows

install_file sdk/python/dist/khulnasoft-analyzer-policy-python                  linux   darwin
install_file sdk/python/dist/khulnasoft-analyzer-policy-python.cmd              windows

install_file sdk/python/dist/khulnasoft-resource-khulnasoft-python                  linux   darwin
install_file sdk/python/dist/khulnasoft-resource-khulnasoft-python.cmd              windows

install_file sdk/python/dist/khulnasoft-python-shim.cmd                         windows
install_file sdk/python/dist/khulnasoft-python3-shim.cmd                        windows

install_file sdk/python/cmd/khulnasoft-language-python-exec          linux darwin windows

# Get khulnasoft-watch binaries
./scripts/get-khulnasoft-watch.sh "${LOCAL}"
./scripts/get-language-providers.sh "${LOCAL}"
