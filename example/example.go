package main

import (
	"flag"

	"github.com/xuyao91/version"
)

//通过flag包设置-version参数
var printVersion bool

func init() {
	flag.BoolVar(&printVersion, "version", false, "print program build version")
	flag.Parse()
}

func main() {
	if printVersion {
		version.Print()
	}
}
