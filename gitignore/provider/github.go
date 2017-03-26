package provider

import (
	"context"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"

	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

// GithubProvider returns a gitignoreProvider,
// which uses the official gitignore GitHub repo,
// as a source for all the gitignore info/content requests.
func GithubProvider(owner, repo, token string) GitignoreProvider {
	return &githubProvider{
		owner: owner,
		repo:  repo,
		token: token,
	}
}

type githubProvider struct {
	owner string
	repo  string
	token string
}

// Get implements GitignoreProvider.Get
func (p *githubProvider) Get(template string) (content []byte, err error) {
	template = strings.TrimSuffix(template, ".gitignore")
	url := fmt.Sprintf("%s/%s.gitignore", p.getRawURL(), template)

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

	branch, _, err := client.Repositories.GetBranch(
		ctx, p.owner, p.repo, "master")
	if err != nil {
		err = fmt.Errorf("couldn't get repo master branch: %s", err)
		return
	}

	masterSHA := branch.Commit.Commit.Tree.GetSHA()
	tree, _, err := client.Git.GetTree(
		ctx, p.owner, p.repo, masterSHA, false)
	if err != nil {
		err = fmt.Errorf("couldn't get repo master tree: %s", err)
		return
	}

	for _, entry := range tree.Entries {
		if match := ghrex.FindStringSubmatch(entry.GetPath()); len(match) == 2 {
			templates = append(templates, match[1])
		}
	}

	if len(templates) == 0 {
		err = errors.New("no templates could be found")
		return
	}

	return
}

func (p *githubProvider) getRawURL() string {
	return fmt.Sprintf(
		"https://raw.githubusercontent.com/%s/%s/master",
		p.owner, p.repo)
}

// Regex
const ghre = `^([A-Z][A-Za-z+_\-0-9]*)\.gitignore$`

var ghrex = regexp.MustCompile(ghre)
