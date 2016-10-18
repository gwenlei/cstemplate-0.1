#!/bin/bash
export GOPATH=`pwd`/../../
go get github.com/zieckey/goini
go get github.com/kesselborn/go-getopt
go get github.com/google/uuid
go build cstemplate.go
