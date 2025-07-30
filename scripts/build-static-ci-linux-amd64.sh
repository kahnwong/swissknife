#!/bin/bash

cd lib/system || exit
cross build --release --target x86_64-unknown-linux-musl
cd ../..

cp lib/system/target/x86_64-unknown-linux-musl/release/libsystem.a lib/
