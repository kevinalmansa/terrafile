package terrafile

import (
	"github.com/kevinalmansa/terrafile/pkg/module"
)

//Config file representation
type Config struct {
	CacheDir string
	Branch   string
	Tag      string
	Modules  map[string]module.TerraformModule
}
