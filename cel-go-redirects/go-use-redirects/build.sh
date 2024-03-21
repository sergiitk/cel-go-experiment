#!/bin/sh -e

set -x
export GOPROXY=direct
export GOINSECURE="cel.sergii.org/*"
export GONOSUMDB="cel.sergii.org/*"
export GOPRIVATE="cel.sergii.org/*"

# uncomment to clean
# go clean -modcache
go mod download -x
go build -v ./...
set +x

echo "Build success"
