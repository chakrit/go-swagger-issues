#!/usr/bin/env bash

set -euo pipefail

export GOPATH=$PWD

go get -v github.com/go-swagger/go-swagger/...

cd src/output
../../bin/swagger generate server -f=../../swagger.yml
cd ../..

go build -v output/models



