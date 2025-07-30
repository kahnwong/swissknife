#!/bin/bash

cd lib/system || exit
cross build --release --target aarch64-linux-musl-gcc
cd ../..

cp lib/system/target/aarch64-linux-musl-gcc/release/libsystem.a lib/
