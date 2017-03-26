package version

import "fmt"

// Version represents the current version of go-gitignore.
type Version struct {
	Major string
	Minor string
	Build string
}

var (
	// GogitignoreVersion is the current version of go-gitignore.
	GogitignoreVersion = Version{Major: "0", Minor: "1"}
)

func (v Version) String() string {
	ver := fmt.Sprintf("Version: %s.%s", v.Major, v.Minor)
	return fmt.Sprintf("%s\nBuild: %s", ver, v.Build)
}
