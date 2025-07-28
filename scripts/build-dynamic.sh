#!/bin/bash

cd lib/system || exit
cargo build --release
cd ../..

if [[ "$(uname)" == "Linux" ]]; then
	cp lib/system/target/release/libsystem.so lib/
elif [[ "$(uname)" == "Darwin" ]]; then
	cp lib/system/target/release/libsystem.dylib lib/
elif [[ "$(uname)" == "Windows" ]]; then
	cp lib/system/target/release/libsystem.dll lib/
fi

go build -ldflags="-r lib"
