package cmd

import (
	"fmt"

	"github.com/glendc/go-gitignore/gitignore/version"

	"github.com/spf13/cobra"
)

const (
	cliVersion = "1.0"
)

// versionCmd represents the command used to get the version number
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Get the version number of the cli.",
	Run:   printVersion,
}

func printVersion(*cobra.Command, []string) {
	fmt.Printf("go-gitignore\n%s\n", version.GogitignoreVersion)
}
