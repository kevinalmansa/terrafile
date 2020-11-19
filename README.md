# Terrafile

[Terrafile](https://bensnape.com/2016/01/14/terraform-design-patterns-the-terrafile/) implementation in Golang for use with modular trunk-based git workflows.

```
$ terrafile
terrafile is a CLI command to enable dynamic versioning of terraform
modules stored in git without modifying the terraform code. The aim is to
simplify development of terraform modules and integration within CI/CD
solutions.

Usage:
  terrafile [command]

Available Commands:
  delete      Delete modules in cache
  help        Help about any command
  install     install/update modules based on config file
  show-config show loaded configuration

Flags:
      --branch string   branch to checkout for modules (default is main)
      --cache string    cache directory (default is ./terrafile/modules)
      --config string   config file (default is ./terrafile/config.yaml)
  -h, --help            help for terrafile
      --tag string      tag to checkout for modules (default is unset - overrides branch)

Use "terrafile [command] --help" for more information about a command.
```

## Example File

```yaml
CacheDir: "./modules"
Branch: "main"
modules:
  vpc:
    repo:  "git@github.com:kevinalmansa/terrafile"
    tag: "v0.0.1"

  iam_role:
    repo:  "git@github.com:kevinalmansa/vagrant-kubernetes"
    branch: "k8s-tooling"
```

Overrides can be specified for each module as seen above, such as branch, tag or repo. This can
be used to "fix" specific modules to specific versions.

## Install

From within this directory, run:
```
go install ./
```

The binary will be installed to your GOPATH/bin directory.

## Build

```
go build -o terrafile ./main.go
```

## Dependencies

- https://github.com/go-git/go-git for git operations
- https://github.com/spf13/viper for configuration
- https://github.com/spf13/cobra for the CLI

## FAQ

### Flag Order

In order of priority (highest to lowest):
1. Module specific flag specified in config file (branch or tag)
2. Command paramaters (--branch, --tag, etc)
3. Global value set in config file ("Branch", "Tag")
4. Default values

### How to use with private repos?

Authentication is not managed by this application - if you require authentication
for accessing private repos, please use SSH authentication with your git repo.

The SSH Agent running on your machine will then handle the authentication.

```sh
$ eval "$(ssh-agent -s)" # start the ssh agent if not already running
$ ssh-add ~/.ssh/id_rsa # example key name
$ terrafile install # use terrafile
```
