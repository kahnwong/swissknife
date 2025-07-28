#!/bin/bash

cd lib/system
cargo build --release
cd ../..

cp lib/system/target/release/libsystem.a lib/
