#!/bin/bash

cd lib/system || exit
cargo build --release
cd ../..

if [[ "$(uname)" == "Linux" ]]; then
	cp lib/system/target/release/libsystem.so lib/
elif [[ "$(uname)" == "Darwin" ]]; then
	cp lib/system/target/release/libsystem.dylib lib/
fi

go build -ldflags="-r lib"
