#!/bin/bash

cd lib/system || exit
cross build --release --target "$TARGET"
cd ../..

cp "lib/system/target/$TARGET/release/libsystem.a" lib/
