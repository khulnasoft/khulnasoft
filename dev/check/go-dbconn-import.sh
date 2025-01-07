#!/usr/bin/env bash

# This script ensures pkg/database/dbconn is only imported by services allowed to
# directly speak with the database.

echo "--- go dbconn import"

trap "echo ^^^ +++" ERR

set -euf -o pipefail
cd "$(dirname "${BASH_SOURCE[0]}")"/../..

allowed_prefix=(
  github.com/khulnasoft/khulnasoft/cmd/embeddings
  github.com/khulnasoft/khulnasoft/cmd/frontend
  github.com/khulnasoft/khulnasoft/cmd/gitserver
  github.com/khulnasoft/khulnasoft/cmd/migrator
  # Transitively depends on updatecheck package which imports but does not use DB
  github.com/khulnasoft/khulnasoft/cmd/pings
  github.com/khulnasoft/khulnasoft/cmd/precise-code-intel-worker
  github.com/khulnasoft/khulnasoft/cmd/syntactic-code-intel-worker
  github.com/khulnasoft/khulnasoft/cmd/repo-updater
  # Doesn't connect but uses db internals for use with sqlite
  github.com/khulnasoft/khulnasoft/cmd/symbols
  github.com/khulnasoft/khulnasoft/cmd/worker
)

# Create regex ^(a|b|c)
allowed=$(printf "|%s" "${allowed_prefix[@]}")
allowed=$(printf "^(%s)" "${allowed:1}")

# shellcheck disable=SC2016
template='{{with $pkg := .}}{{ range $pkg.Deps }}{{ printf "%s imports %s\n" $pkg.ImportPath .}}{{end}}{{end}}'

if go list ./cmd/... |
  grep -Ev "$allowed" |
  xargs go list -f "$template" |
  grep "github.com/khulnasoft/khulnasoft/internal/database/dbconn"; then
  echo "Error: the above service(s) are not allowed to import internal/database/dbconn"
  echo "^^^ +++"
  exit 1
fi
