package provider

// GitignoreProvider describes the interface
// which provides gitignore content and information.
// The details are left to the implementations.
type GitignoreProvider interface {
	Get(template string) (content []byte, err error)
	List() (templates []string, err error)
}
