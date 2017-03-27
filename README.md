# go-gitignore

A Go CLI to create and list gitignore files,
taken from a github repository or local directory.

## Configuration

Example YAML Configuration (default: `$HOME/.go-gitignore.yaml`):

```yaml
---
provider: github # default
log:
    path: <PATH> # empty by defauld
github:
    repository: github/gitignore # default
    token: <GITHUB_TOKEN> # empty by defualt
local:
    path: <PATH> # empty by default
```
