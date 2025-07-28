#!/bin/bash

cd lib/system || exit
cargo build --release
cd ../..

if [[ "$(uname)" == "Linux" ]]; then
	cp lib/system/target/release/libsystem.a lib/
elif [[ "$(uname)" == "Darwin" ]]; then
	cp lib/system/target/release/libsystem.a lib/
fi

go build
