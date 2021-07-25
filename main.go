package main

import (
	"daikincli/cmd"
)

// Values for version and commit are injected by the build.
var (
	version = "edge"
	commit  = "n/a"
)

func main() {
	cmd.Execute(version, commit)
}
