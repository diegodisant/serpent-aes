#!/bin/bash

echo "Building docker environment ....."

rm serpent-cli
docker stop serpent-aes
docker rm serpent-aes
docker rmi serpent-aes
docker build -t serpent-aes .
docker run -it --name serpent-aes serpent-aes:latest
