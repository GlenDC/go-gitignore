package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// listTemplatesCmd represents the command used to create gitignore files
var listTemplatesCmd = &cobra.Command{
	Use:   "templates",
	Short: "Lists all available templates in the master branch.",
	RunE:  listTemplates,
}

// templates is the function for the TemplatesCmd
func listTemplates(*cobra.Command, []string) error {
	provider, err := newProvider()
	if err != nil {
		return err
	}

	templates, err := provider.List()
	if err != nil {
		return err
	}

	for _, template := range templates {
		fmt.Println(template)
	}

	return nil
}
