#!/usr/bin/env sh
# vim: ft=sh tw=80

BIN_NAME="nvdcve2json"

if [ -d build ]; then
  rm build/*
else
  mkdir build
fi
go build
mv $BIN_NAME build/"$BIN_NAME"_linux_x86_64

env GOOS=linux GOARCH=amd64 go build
mv $BIN_NAME build/"$BIN_NAME"_darwin_x86_64

cd build && shasum -a 512 "$BIN_NAME"_*_x86_64 > SHA512SUMS.txt && cd .. || exit 
gpg --clearsign build/SHA512SUMS.txt
