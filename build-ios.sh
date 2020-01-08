#!/bin/bash

ROOT_DIR=../

echo "Building serpent_aes framework for iOS ....."
rm Serpent_aes.framework
mv serpent-cli.go serpent-cli.go-rr
gomobile bind -target=ios serpent_aes
mv serpent-cli.go-rr serpent-cli.go
ls -l Serpent_aes.framework
 
