package cmd

import (
	"errors"
	"fmt"

	"github.com/spf13/cobra"
)

// printCmd represents the command used to print
// the content of a gitignore template
var printCmd = &cobra.Command{
	Use:   "print template...",
	Short: "Print the content of a template.",
	RunE:  print,
}

func print(cmd *cobra.Command, args []string) error {
	if len(args) < 1 {
		return errors.New("at least 1 template argument is required")
	}
	args = unique(args)

	// download the gitignore file based on a template
	content, err := downloadAll(args...)
	if err != nil {
		return err
	}

	fmt.Println(string(content))
	return nil
}
