# terrafile

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
cacheDir: "./modules"
branch: "main"
modules:
  vpc:
    repo:  "git@github.com:kevinalmansa/aws-vpc"
    tag: "v1.0.1"

  iam_role:
    repo:  "git@github.com:kevinalmansa/modules"
    branch: "dev"
    path: "data/security_groups"

```

Overrides can be specified for each module as seen above, such as branch, tag, repo, or path. This can
be used to "fix" specific modules to specific versions.

## Build binary

```
go build -o terrafile
```

## Dependencies

https://github.com/go-git/go-git for git operations

https://github.com/spf13/viper for configuration
