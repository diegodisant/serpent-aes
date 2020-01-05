#!/bin/bash

echo "Configuring project (YOU MUST NEED TO HAVE GO INSTALLED) ...."
./setup.sh

echo
echo "-----"
echo

echo "Installing dependiencies ....."
./install-deps.sh

echo
echo "-----"
echo

echo "Building serpent_aes framework for iOS ....."
rm Serpent_aes.framework
mv serpent-cli.go serpent-cli.go-rr
gomobile bind -target=ios serpent_aes
mv serpent-cli.go-rr serpent-cli.go
ls -l Serpent_aes.framework

echo "Please drag and drop serpent framework on Frameworks folder for telegram application"

echo
echo "-----"
echo

echo "Building docker environment ....."
rm serpent-cli
docker stop bmex-go-serpent-aes
docker rm bmex-serpent
docker rmi bmex-go-serpent-aes
docker build -t bmex-go-serpent-aes .
docker run -it --name bmex-serpent bmex-go-serpent-aes:latest

echo
echo "-----"
echo
