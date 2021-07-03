package main

import (
	"flag"
	"fmt"
)

var flagVersion = flag.Bool("v", false, "`dg -v` show devgo version")

// these information will be collected when build, by `-ldflags "-X appversion.appVersion=0.1"`
// They have to be global variables in package to accept the data by `-ldflags`.
var (
	appVersion string
	buildTime  string
	gitCommit  string
)

func showVersion() bool {
	if !*flagVersion {
		return false
	}
	fmt.Printf("version: %s build_time: %s git_commit: %s", appVersion, buildTime, gitCommit)
	return true
}
