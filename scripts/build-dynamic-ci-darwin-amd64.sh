#!/bin/bash

cd lib/system || exit
cross build --release --target x86_64-apple-darwin
cd ../..

cp lib/system/target/x86_64-apple-darwin/release/libsystem.dylib lib/
