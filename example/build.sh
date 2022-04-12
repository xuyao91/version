#!/bin/sh
version="v0.1"
path="github.com/xuyao91/version"
flags="-X $path.Version=$version -X '$path.GoVersion=$(go version)' -X '$path.BuildTime=`date +"%Y-%m-%d %H:%m:%S"`' -X $path.GitCommit=`git rev-parse HEAD`"
go build -ldflags "$flags" -o example example.go
