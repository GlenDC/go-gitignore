package cmd

import (
	"fmt"
	"strings"

	"github.com/glendc/go-gitignore/gitignore/provider"
)

type providerKind string

// Set is only called when flag is defined,
// therefore we'll default providerKind to "github"
func (pk *providerKind) Set(val string) error {
	if val != "github" {
		return fmt.Errorf("%q is not a valid providerKind", val)
	}

	*pk = providerKind(val)
	return nil
}

func (pk *providerKind) Type() string {
	return "providerKind"
}

func (pk *providerKind) String() string { return string(*pk) }

func newProvider() (provider.GitignoreProvider, error) {
	switch pkind {
	default:
		parts := strings.Split(ghrepo, "/")
		if len(parts) != 2 {
			return nil, fmt.Errorf("%q is not a valid github repo", ghrepo)
		}

		return provider.GithubProvider(parts[0], parts[1], ghtoken), nil
	}
}

// URL constants
const (
	repository = "https://raw.githubusercontent.com/github/gitignore/master"
)

// downloadAll downloads all gitignore files based on given templates,
// or none in case of an error
func downloadAll(provider provider.GitignoreProvider, templates ...string) ([]byte, error) {
	if provider == nil {
		return nil, fmt.Errorf("no gitignore provider given")
	}

	var content, current []byte
	var header string
	var err error

	for _, template := range templates {
		current, err = provider.Get(template)
		if err != nil {
			return nil, fmt.Errorf("failed to get %q: %s", template, err)
		}

		current = append(current, '\n')
		header = fmt.Sprintf("# %s\n\n", template)
		content = append(content, []byte(header)...)
		content = append(content, current...)
	}

	return content, nil
}

// unique returns all unique elements in an array (n^2)
func unique(input []string) (output []string) {
	for _, k := range input {
		var found bool
		for _, p := range output {
			if k == p {
				found = true
				break
			}
		}

		if !found {
			output = append(output, k)
		}
	}

	return
}
