package version

import (
	"fmt"
	"os"
)

var (
	//Project Version Information
	Version = ""
	//Golang Version Information
	GoVersion = ""
	//Git Branch Information
	GitBranch = ""
	//Git Commit Information
	GitCommit = ""
	//Build Time
	BuildTime = ""
)

//PrintVersion 输出版本信息
func Print() {
	fmt.Printf("Version: %s\n", Version)
	fmt.Printf("Go Version: %s\n", GoVersion)
	fmt.Printf("Git Commit: %s\n", GitCommit)
	fmt.Printf("Git Branch: %s\n", GitBranch)
	fmt.Printf("Build Time: %s\n", BuildTime)
	os.Exit(0)
}
