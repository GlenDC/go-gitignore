package provider

import (
	"path"
	"regexp"
)

func extractTemplateName(pathName string) (name string, ok bool) {
	parts := templateFileRex.FindStringSubmatch(path.Base(pathName))
	if len(parts) != 2 {
		ok = false
		return
	}

	name, ok = parts[1], true
	return
}

// Regex
const templateFileRe = `(.+)\.gitignore$`

var templateFileRex = regexp.MustCompile(templateFileRe)
