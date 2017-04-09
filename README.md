# go-gitignore [![license](https://img.shields.io/github/license/glendc/go-gitignore.svg)](https://github.com/GlenDC/go-gitignore/blob/master/LICENSE)

A Go CLI to create and list gitignore files,
taken from a github repository or local directory.

## Install

Install from source requires Go 1.6 or above, and can be done as follows:

```
$ go install github.com/glendc/go-gitignore
```

Or you can [download the latest release](https://github.com/GlenDC/go-gitignore/releases) for your supported platform of your choice.

## Usage

```
$ gitignore --help
Create and list gitignore files.

Usage:
  gitignore [command]

Available Commands:
  create      Create a new gitignore file based on given templates.
  list        List a resource via its representative subcommand.
  print       Print the content of a template.
  version     Get the version number of the cli.

Flags:
      --config string           config file used (default: $HOME/.go-gitignore.yaml)
      --github-repo string      github repo used in case github provider is used (default: github/gitignore)
      --github-token string     github token used for some commands in case github provider is used
      --local-path string       local dir used in case local provider is used
      --log-path string         log file used, logs to STDERR if no file is specified
      --provider providerKind   defines the provider to use for getting gitignore content, options: github, local (default: github)
  -v, --verbose                 log info logs in case verbose is enabled

Use "gitignore [command] --help" for more information about a command.
```

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

All configuration parameters also have a CLI flag alternative,
which has precedence over the YAML configuration.

## Contributions

Contributions are welcome. Please always file an issue first,
whether it is about a bug or a feature. This way a discussion can take place
before any work is done.
