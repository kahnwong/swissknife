#!/bin/bash

cd lib/system || exit
cargo build --release
cd ../..

cp lib/system/target/release/libsystem.a lib/

go build -ldflags="-extldflags=-static"
