#!/bin/sh -e

set -x
export GOPROXY=direct
export GOINSECURE="cel.wtf/*"
export GONOSUMDB="cel.wtf/*"
export GOPRIVATE="cel.wtf/*"

# uncomment to clean
# go clean -modcache
go mod download -x
go build -v ./...
set +x

echo "Build success"
