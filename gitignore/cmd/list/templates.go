package list

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"regexp"

	"github.com/google/go-github/github"
	"github.com/spf13/cobra"
	"golang.org/x/oauth2"
)

// URL constants
const (
	repoOwner  = "github"
	repoName   = "gitignore"
	repoBranch = "master"
)

// Regex
const re = `^([A-Z][A-Za-z+_\-0-9]*)\.gitignore$`

var rex = regexp.MustCompile(re)

// local flags
var templatesCfg struct {
	// OAuth2 Access Token for GitHub, otherwise public API is used
	Token string
}

// TemplatesCmd represents the command used to create gitignore files
var TemplatesCmd = &cobra.Command{
	Use:   "templates",
	Short: "Lists all available templates in the master branch.",
	RunE:  templates,
}

// templates is the function for the TemplatesCmd
func templates(*cobra.Command, []string) error {
	var tc *http.Client

	ctx := context.Background()

	// specify tc in case token is given by user
	if templatesCfg.Token != "" {
		ts := oauth2.StaticTokenSource(
			&oauth2.Token{AccessToken: templatesCfg.Token},
		)
		tc = oauth2.NewClient(ctx, ts)
	}

	// create GitHub client
	client := github.NewClient(tc)

	branch, _, err := client.Repositories.GetBranch(
		ctx, repoOwner, repoName, repoBranch)
	if err != nil {
		return fmt.Errorf("couldn't get repo master branch: %s", err)
	}

	masterSHA := branch.Commit.Commit.Tree.GetSHA()
	tree, _, err := client.Git.GetTree(
		ctx, repoOwner, repoName, masterSHA, false)
	if err != nil {
		return fmt.Errorf("couldn't get repo master tree: %s", err)
	}

	var found bool
	for _, entry := range tree.Entries {
		if match := rex.FindStringSubmatch(entry.GetPath()); len(match) == 2 {
			fmt.Println(match[1])
			found = true
		}
	}

	if !found {
		return errors.New("no templates could be found")
	}

	return nil
}

func init() {
	// local flags

	TemplatesCmd.Flags().StringVarP(
		&templatesCfg.Token, "token", "k", "",
		"OAuth2 Access Token for GitHub")
}
