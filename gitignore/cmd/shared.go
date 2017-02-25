package cmd

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

// URL constants
const (
	repository = "https://raw.githubusercontent.com/github/gitignore/master"
)

// template alias
var templateAliasMap = map[string]string{
	"golang": "Go",
	"cpp":    "C++",
}

// download a gitignore file based on a given template
func download(template string) ([]byte, error) {
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
