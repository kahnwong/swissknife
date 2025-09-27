#!/bin/bash

# build static lib
if [[ $(uname -s) == 'Linux' ]]; then
	TARGET="x86_64-unknown-linux-gnu"
elif [[ $(uname -s) == 'Darwin' ]]; then
	TARGET="aarch64-apple-darwin"
fi

cd lib/system || exit
cross build --release --target "$TARGET"

cd ../..
cp "lib/system/target/$TARGET/release/libsystem.a" lib/

# build go binary
if [[ $(uname -s) == 'Linux' ]]; then
	go build -ldflags="-extldflags=-static"
elif [[ $(uname -s) == 'Darwin' ]]; then
	CGO_ENABLED=1 go build -ldflags="-s -w -linkmode 'external'" .
fi
