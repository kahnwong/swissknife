#!/bin/bash

cd lib/system || exit
cross build --release --target aarch64-apple-darwin
cd ../..

cp lib/system/target/aarch64-apple-darwin/release/libsystem.dylib lib/
