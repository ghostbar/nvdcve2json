#!/usr/bin/env sh
# vim: ft=sh tw=80

BIN_NAME="nvdcve2json"
TARGETS="linux;amd64 linux;arm darwin;amd64"

if [ -d build ]; then
  rm build/*
else
  mkdir build
fi

for target in $TARGETS; do
  targetos=$(echo $target | awk -F ";" '{ print $1 }')
  targetarch=$(echo $target | awk -F ";" '{ print $2 }')
  env GOOS=$targetos GOARCH=$targetarch go build
  mv $BIN_NAME build/"$BIN_NAME"_"$targetos"_"$targetarch"
done

cd build && shasum -a 512 "$BIN_NAME"_* > SHA512SUMS.txt && cd .. || exit 
gpg --clearsign build/SHA512SUMS.txt
