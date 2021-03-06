package provider

import (
	"context"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/glendc/go-gitignore/gitignore/logger"
	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

// GithubProvider returns a gitignoreProvider,
// which uses the official gitignore GitHub repo,
// as a source for all the gitignore info/content requests.
func GithubProvider(owner, repo, token string, logger logger.Logger) GitignoreProvider {
	return &githubProvider{
		owner:  owner,
		repo:   repo,
		token:  token,
		logger: logger,
	}
}

type githubProvider struct {
	owner  string
	repo   string
	token  string
	logger logger.Logger
}

// Get implements GitignoreProvider.Get
func (p *githubProvider) Get(template string) (content []byte, err error) {
	template = strings.TrimSuffix(template, ".gitignore")
	url := fmt.Sprintf("%s/%s.gitignore", p.getRawURL(), template)

	p.logger.Infof("downloading template from %q", url)

	resp, err := http.DefaultClient.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New(resp.Status)
	}

	return ioutil.ReadAll(resp.Body)
}

// List implements GitignoreProvider.List
func (p *githubProvider) List() (templates []string, err error) {
	var tc *http.Client

	ctx := context.Background()

	// specify tc in case token is given by user
	if p.token != "" {
		ts := oauth2.StaticTokenSource(
			&oauth2.Token{AccessToken: p.token},
		)
		tc = oauth2.NewClient(ctx, ts)
	}

	// create GitHub client
	client := github.NewClient(tc)

	p.logger.Infof(`listing templates from "github.com/%s/%s"`, p.owner, p.repo)

	branch, _, err := client.Repositories.GetBranch(
		ctx, p.owner, p.repo, "master")
	if err != nil {
		err = fmt.Errorf("couldn't get repo master branch: %s", err)
		return
	}

	masterSHA := branch.Commit.Commit.Tree.GetSHA()
	templates, err = p.getTreeEntries(ctx, client, masterSHA, "")

	if len(templates) == 0 {
		err = errors.New("no templates could be found")
		return
	}

	return
}

func (p *githubProvider) getTreeEntries(ctx context.Context, client *github.Client, sha, subdir string) (templates []string, err error) {
	tree, _, err := client.Git.GetTree(ctx, p.owner, p.repo, sha, false)
	if err != nil {
		err = fmt.Errorf("couldn't get repo master tree: %s", err)
		return
	}

	var subtrees []github.TreeEntry

	for _, entry := range tree.Entries {
		if entry.GetMode() == "040000" {
			// subdirectory, we'll parse it later
			subtrees = append(subtrees, entry)
			continue
		}

		if template, ok := extractTemplateName(entry.GetPath()); ok {
			if subdir != "" {
				template = subdir + template
			}
			templates = append(templates, template)
		}
	}

	var subTemplates []string
	// we go through the subdirs as a last step,
	// so we can add them in order of being found,
	// at the front of the templates list
	for _, subtree := range subtrees {
		subTemplates, err = p.getTreeEntries(
			ctx,
			client,
			subtree.GetSHA(),
			subdir+subtree.GetPath()+"/",
		)
		if err != nil {
			templates = nil
			return
		}

		templates = append(subTemplates, templates...)
	}

	return
}

func (p *githubProvider) getRawURL() string {
	return fmt.Sprintf(
		"https://raw.githubusercontent.com/%s/%s/master",
		p.owner, p.repo)
}
