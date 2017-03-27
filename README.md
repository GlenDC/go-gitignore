# go-gitignore

A Go CLI to manage gitignore files based on the github/gitignore repo

## Configuration

Example YAML Configuration (default: $HOME/.go-gitignore.yaml):

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
