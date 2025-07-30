#!/bin/bash

cd lib/system || exit
cross build --release --target aarch64-unknown-linux-musl
cd ../..

cp lib/system/target/aarch64-unknown-linux-musl/release/libsystem.a lib/
