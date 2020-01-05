#!/bin/bash

if [ -z $GOPATH ]
  then
  echo "Please configure GOPATH to run this script"
fi

MAIN_PACKAGE_NAME=$(basename $(pwd))
MAIN_PACKAGE_PATH=$(pwd)
GO_SOURCES="${GOPATH}/src/"

echo "[+] Creating symlink in GOPATH to be able to compile project"

sudo ln -s $MAIN_PACKAGE_PATH $GO_SOURCES$MAIN_PACKAGE_NAME
