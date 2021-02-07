#!/bin/bash

set -ex

cd $(dirname $0)

go build -o .tmpbin
VERSION_OUTPUT=$(./.tmpbin --version)
VERSION=${VERSION_OUTPUT:9}
rm ./.tmpbin

rm -rf dist/
mkdir dist/

for combo in linux/amd64 linux/arm64 darwin/amd64 windows/amd64
do IFS="/"; set -- $combo;
  OS="$1"
  ARCH="$2"
  BASE_NAME="dist/http-testing-server-$VERSION-$OS-$ARCH"
  BIN="$BASE_NAME"
  if [ "$OS" = "windows" ]
  then
    BIN="$BIN.exe"
  fi
  GOOS=$OS GOARCH=$ARCH go build -o "$BIN"
  tar -czvf "$BASE_NAME.tar.gz" "$BIN"
done