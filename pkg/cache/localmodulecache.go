package cache

import (
	"io/ioutil"
	"log"
	"os"
	"path"

	"github.com/kevinalmansa/terrafile/cmd/terrafile"
	"github.com/kevinalmansa/terrafile/pkg/module"
)

//LocalModuleCache is an implimentation of the Module interface for creating a
//local cache to store modules.
type LocalModuleCache struct {
	configuration *terrafile.Config
}

func configureModule(configuration *terrafile.Config, module *module.TerraformModule) {
	if module.Tag == "" && configuration.Tag != "" {
		module.Tag = configuration.Tag
	}
	if module.Branch == "" {
		module.Branch = configuration.Branch
	}
}

//Create all modules, storing them in the cache
func (c *LocalModuleCache) Create() error {
	if c.configuration == nil {
		return &InitError{}
	}
	for name, module := range c.configuration.Modules {
		modLocation := path.Join(c.configuration.CacheDir, name)

		configureModule(c.configuration, &module)
		if _, err := os.Stat(modLocation); os.IsNotExist(err) {
			if err = module.Clone(modLocation); err != nil {
				log.Printf("Error cloning module: %s", err)
			}
		} else {
			if err = module.Update(modLocation); err != nil {
				log.Printf("Error updating module: %s", err)
			}
		}
	}
	return nil
}

//Delete cache
func (c *LocalModuleCache) Delete(deleteAll bool) error {
	files, err := ioutil.ReadDir(c.configuration.CacheDir)
	if err != nil {
		return err
	}

	for _, f := range files {
		if f.IsDir() {
			if _, ok := c.configuration.Modules[f.Name()]; deleteAll || ok {
				err := os.RemoveAll(f.Name())
				if err != nil {
					log.Printf("Error removing cached directory: %s", err)
				}
			}
		}
	}
	return nil
}
