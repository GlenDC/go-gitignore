package cmd

import (
	"fmt"
	"runtime"

	"github.com/spf13/cobra"
)

const (
	cliVersion = "1.0"
)

// versionCmd represents the command used to get the version number
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Get the version number of the cli.",
	Run:   version,
}

func version(*cobra.Command, []string) {
	fmt.Printf("gitignore v%s %s/%s\n",
		cliVersion, runtime.GOOS, runtime.GOARCH)
}
