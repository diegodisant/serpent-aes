#!/bin/sh

rm /bin/serpent-cli
go build serpent-cli.go
ln -s $(pwd)/serpent-cli /bin/serpent-cli
