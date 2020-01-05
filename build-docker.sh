#!/bin/bash

echo "Building docker environment ....."

rm serpent-cli
docker stop bmex-go-serpent-aes
docker rm bmex-serpent
docker rmi bmex-go-serpent-aes
docker build -t bmex-go-serpent-aes .
docker run -it --name bmex-serpent bmex-go-serpent-aes:latest
