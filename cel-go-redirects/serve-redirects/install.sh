#!/bin/sh -e

if [[ ! -d ./venv ]]; then
  python3 -m venv venv
  ./venv/bin/python3 -c 'import sys; assert sys.version_info >= (3, 11), f"Python 3.11 required, got {sys.version_info}"'
fi

if [[ ! -d ./certs || ! -f ./certs/server.pem ]]; then
  mkdir -p certs
  pushd certs > /dev/null
  set -x
  openssl req -new -x509 -keyout key.pem -out server.pem -days 365 -nodes \
    -subj "/C=US/ST=CA/O=sergiitk/CN=cel.sergii"
  set +x
  popd > /dev/null
fi

echo "Done``"
