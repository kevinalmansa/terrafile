# terrafile

[Terrafile](https://bensnape.com/2016/01/14/terraform-design-patterns-the-terrafile/) implementation in Golang for use with modular trunk-based git workflows.

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

## Dependencies

https://github.com/go-git/go-git for git operations

https://github.com/spf13/viper for configuration
