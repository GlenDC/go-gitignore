package cmd

import (
	"errors"
	"fmt"

	"github.com/spf13/cobra"
)

// printCmd represents the command used to print
// the content of a gitignore template
var printCmd = &cobra.Command{
	Use:   "print template",
	Short: "Print the content of a template.",
	RunE:  print,
}

func print(cmd *cobra.Command, args []string) error {
	if len(args) != 1 {
		return errors.New("template argument is required")
	}

	// download the gitignore file based on a template
	template := args[0]
	content, err := download(template)
	if err != nil {
		return fmt.Errorf(
			"couldn't downloaded gitignore template %q: %s", template, err)
	}

	fmt.Println(string(content))
	return nil
}
