#!/bin/bash

#
# go PAckage Dependency Graph analizer (PADG)
#

if [ -z $1 ]
  then
  echo "Usage: ./$0 <go-package>"
  exit 1
fi

# check dependencies

if [ ! -f $GODEGRAPH_BIN ];
  then
  echo "Error: Godegraph package doesn't exists"
  echo "[+] Installing"

  go get github.com/kisielk/godepgraph
fi

GO_PACKAGE=$1
GODEGRAPH_BIN="${GOPATH}/bin/godepgraph"
GO_PACKAGE_DIR="${GOPATH}/src/${GO_PACKAGE}"
GRAPH_IMAGE_NAME=$(echo $GO_PACKAGE | awk '{subPackagesSize=split($0, packages, "/"); print packages[subPackagesSize]}')
GRAP_IMAGE="${GRAPH_IMAGE_NAME}-package.png"

# NEEDS TO HAVE INSTALLED dot (graphviz), godepgraph

cd $GO_PACKAGE_DIR

$GODEGRAPH_BIN $GO_PACKAGE | dot -Tpng -o $GRAP_IMAGE
