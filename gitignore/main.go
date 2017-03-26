package main

import (
	"github.com/glendc/go-gitignore/gitignore/cmd"
	"github.com/glendc/go-gitignore/gitignore/version"
)

// Build is the git sha of this binaries build.
var Build string

func main() {
	version.GogitignoreVersion.Build = Build
	cmd.Execute()
}
