package cmd

import (
	"github.com/glendc/go-gitignore/gitignore/cmd/list"
	"github.com/spf13/cobra"
)

// listCmd represents a root command for all resource list commands
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List a resource via its representative subcommand.",
}

func init() {
	listCmd.AddCommand(
		list.TemplatesCmd,
	)
}
