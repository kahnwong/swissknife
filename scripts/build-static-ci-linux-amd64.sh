#!/bin/bash

cd lib/system || exit
cross build --release --target x86_64-unknown-linux-gnu
cd ../..

cp lib/system/target/x86_64-unknown-linux-gnu/release/libsystem.a lib/
