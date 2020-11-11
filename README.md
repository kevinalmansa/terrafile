# terrafile

[Terrafile]() implementation in Golang for use with modular trunk-based git workflows.


## Example File

```yaml
tf-aws-vpc:
 source:  "git@github.com:kevinalmansa/terraform-modules"
 branch: "main"
 path: "vpc/"

tf-aws-iam:
 source:  "git@github.com:kevinalmansa/tf-aws-iam"
 version: "v1.1.2"
```
