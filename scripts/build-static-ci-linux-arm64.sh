#!/bin/bash

cd lib/system || exit
cross build --release --target aarch64-unknown-linux-gnu
cd ../..

cp lib/system/target/aarch64-unknown-linux-gnu/release/libsystem.a lib/
