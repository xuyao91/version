#!/bin/sh
version="v1.0"
path="github.com/xuyao91/version"
flags="-X $path.Version=$version -X '$path.GoVersion=$(go version)' -X '$path.BuildTime=`date +"%Y-%m-%d %H:%m:%S"`' -X $path.GitCommit=`git rev-parse HEAD` -X $path.GitBranch=`git rev-parse --abbrev-ref HEAD`"
go build -ldflags "$flags" -o example example.go
