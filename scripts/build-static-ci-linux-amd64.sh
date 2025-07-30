#!/bin/bash

cd lib/system || exit
cross build --release --target x86_64-linux-musl-gcc
cd ../..

cp lib/system/target/x86_64-linux-musl-gcc/release/libsystem.a lib/
