#!/bin/bash

set -ex

export GOPATH=$PWD/../../../../
export GOBIN=$PWD/bin
go install -race ./...
