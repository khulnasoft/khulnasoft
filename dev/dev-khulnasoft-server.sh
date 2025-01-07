#!/usr/bin/env bash

cd "$(dirname "${BASH_SOURCE[0]}")"/..
set -ex

# Build a Khulnasoft server docker image to run for development purposes. Note
# that this image is not exactly identical to the published sourcegraph/server
# images, as those include Khulnasoft Enterprise features.
time cmd/server/pre-build.sh
IMAGE=khulnasoft/server:0.0.0-DEVELOPMENT VERSION=0.0.0-DEVELOPMENT time cmd/server/build.sh

IMAGE=khulnasoft/server:0.0.0-DEVELOPMENT dev/run-server-image.sh
