#!/bin/sh -e

set -x
export GOPROXY=direct
export GOINSECURE="cel.sergii/*"
export GONOSUMDB="cel.sergii/*"
export GOPRIVATE="cel.sergii/*"

# uncomment to clean
# go clean -modcache
go mod download -x
go build -v ./...
set +x

echo "Build success"
