#!/bin/sh
# Instalar kuaforia CLI
set -e
VERSION=${1:-latest}
OS=$(uname -s | tr '[:upper:]' '[:lower:]')
ARCH=$(uname -m)

case $ARCH in
    x86_64) ARCH="amd64" ;;
    aarch64|arm64) ARCH="arm64" ;;
esac

URL="https://github.com/kuaforia/cli/releases/download/$VERSION/kuaforia-${OS}-${ARCH}.tar.gz"
curl -sSL "$URL" | tar -xz -C /usr/local/bin
echo "kuaforia CLI installed to /usr/local/bin/kuaforia"
