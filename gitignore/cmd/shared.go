package cmd

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

// URL constants
const (
	repository = "https://raw.githubusercontent.com/github/gitignore/master"
)

// downloadAll downloads all gitignore files based on given templates,
// or none in case of an error
func downloadAll(templates ...string) ([]byte, error) {
	var content, current []byte
	var header string
	var err error

	for _, template := range templates {
		current, err = download(template)
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

// download a gitignore file based on a given template
func download(template string) ([]byte, error) {
	template = strings.TrimSuffix(template, ".gitignore")
	url := fmt.Sprintf("%s/%s.gitignore", repository, template)

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
