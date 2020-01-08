#!/bin/sh

MAIN_PACKAGE_NAME=$(basename $(pwd))
MAIN_PACKAGE_PATH=$(pwd)
GO_SOURCES=/root/go/src/

ln -s $MAIN_PACKAGE_PATH $GO_SOURCES$MAIN_PACKAGE_NAME

./build-cli.sh

exit 0
