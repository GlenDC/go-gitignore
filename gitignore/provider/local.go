package provider

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"

	"github.com/glendc/go-gitignore/gitignore/logger"
)

// LocalProvider returns a gitignoreProvider,
// which uses a local directory,
// as a source for all the gitignore info/content requests.
func LocalProvider(path string, logger logger.Logger) GitignoreProvider {
	return &localProvider{
		path:   path,
		logger: logger,
	}
}

type localProvider struct {
	path   string
	logger logger.Logger
}

// Get implements GitignoreProvider.Get
func (p *localProvider) Get(template string) (content []byte, err error) {
	tpath := p.templatePath(template)
	p.logger.Infof("opening template from %q", tpath)
	file, err := os.Open(tpath)
	if err != nil {
		err = fmt.Errorf("template %q could be found at %q", template, p.path)
		return
	}
	defer file.Close()

	content, err = ioutil.ReadAll(file)
	return
}

// List implements GitignoreProvider.List
func (p *localProvider) List() (templates []string, err error) {
	p.logger.Infof("listening templates from %q", p.path)
	return p.listDir(p.path, "", 0)
}

func (p *localProvider) listDir(root, dir string, level int) (templates []string, err error) {
	dirPath := path.Join(root, dir)
	files, err := ioutil.ReadDir(dirPath)
	if err != nil {
		err = fmt.Errorf("couldn't list templates at %q: %s", dirPath, err)
		return
	}

	var subdirs []os.FileInfo

	for _, file := range files {
		if file.IsDir() {
			// subdirectory, we'll parse it later
			subdirs = append(subdirs, file)
			continue
		}

		if template, ok := extractTemplateName(file.Name()); ok {
			if dir != "" {
				template = path.Join(dir, template)
			}
			templates = append(templates, template)
		}
	}

	// we limit the levels so this doesn't get suck in an infinite loop
	if level <= 3 {
		// we go through the subdirs as a last step,
		// so we can add them in order of being found,
		// at the front of the templates list
		nextLevel := level + 1
		var subTemplates []string
		for _, subdir := range subdirs {
			subTemplates, err = p.listDir(
				root,
				path.Join(dir, subdir.Name()),
				nextLevel,
			)

			if err != nil {
				templates = nil
				return
			}

			templates = append(subTemplates, templates...)
		}
	}

	return
}

func (p *localProvider) templatePath(template string) string {
	return path.Join(p.path, template+".gitignore")
}
