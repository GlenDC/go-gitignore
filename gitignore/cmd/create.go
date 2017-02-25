package cmd

import (
	"errors"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// local flags
var createCfg struct {
	Target string
	Force  bool
	Append bool
}

// createCmd represents the command used to create gitignore files
var createCmd = &cobra.Command{
	Use:   "create template",
	Short: "Create a new gitignore file based on a template.",
	Long: `Create a new gitignore file based on a template
which by default is fetched from github.com/github/gitignore`,
	RunE: create,
}

func create(cmd *cobra.Command, args []string) error {
	if len(args) != 1 {
		return errors.New("template argument is required")
	}

	if createCfg.Target == "" {
		createCfg.Target = ".gitignore"
	}

	// download the gitignore file based on a template
	template := args[0]
	content, err := download(template)
	if err != nil {
		return fmt.Errorf(
			"couldn't downloaded gitignore template %q: %s", template, err)
	}

	flags, mode := os.O_WRONLY|os.O_CREATE, os.FileMode(0644)

	if createCfg.Force {
		if createCfg.Append {
			flags |= os.O_APPEND
		} else {
			flags |= os.O_TRUNC
		}
	} else {
		flags |= os.O_EXCL | os.O_TRUNC
	}

	// try open/create/append file
	target, err := os.OpenFile(createCfg.Target, flags, mode)
	if err != nil {
		return fmt.Errorf("couldn't open/create %q: %s", createCfg.Target, err)
	}
	defer target.Close()

	// write gitignore file content
	_, err = target.Write(content)
	if err != nil {
		return fmt.Errorf("couldn't write gitignore template content: %s", err)
	}

	return nil
}

func init() {
	// local flags

	createCmd.Flags().StringVarP(
		&createCfg.Target, "target", "t", ".gitignore",
		"defines the target path of the gitignore file")
	createCmd.Flags().BoolVarP(
		&createCfg.Force, "force", "f", false,
		"overwrites the gitignore file, in case it already exists")
	createCmd.Flags().BoolVarP(
		&createCfg.Append, "append", "a", false,
		"appends template content to existing gitignore file (requires force flag)")
}
