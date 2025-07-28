#!/bin/bash

cd lib/system || exit
cross build --release --target x86_64-pc-windows-gnu
tree # debug
cd ../..

cp lib/system/target/x86_64-pc-windows-gnu/release/libsystem.a lib/
