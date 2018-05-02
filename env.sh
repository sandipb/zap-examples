#!/bin/bash

THIS_DIR=$(cd $(dirname $BASH_SOURCE) && pwd)
export GOPATH=$THIS_DIR GOBIN=$THIS_DIR/bin

if [[ ! $PATH == *$THIS_DIR* ]];then
	export PATH=$PATH:$GOBIN
fi
