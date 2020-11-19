package cache

import (
	"io/ioutil"
	"log"
	"os"
	"path"

	terrafile "github.com/kevinalmansa/terrafile/pkg/config"
	"github.com/kevinalmansa/terrafile/pkg/module"
)

//LocalModuleCache is an implimentation of the Module interface for creating a
//local cache to store modules.
type LocalModuleCache struct {
	//Configuration to use by the cache
	Configuration *terrafile.Config
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
	if c.Configuration == nil {
		return &InitError{}
	}
	for name, module := range c.Configuration.Modules {
		modLocation := path.Join(c.Configuration.CacheDir, name)

		configureModule(c.Configuration, &module)
		if _, err := os.Stat(modLocation); os.IsNotExist(err) {
			if err = module.Clone(modLocation); err != nil {
				log.Printf("Error cloning module: %s", err)
			}
		} else {
			log.Printf("Updating module at %s", modLocation)
			if err = module.Update(modLocation); err != nil {
				log.Printf("Error updating module: %s", err)
			}
		}
	}
	return nil
}

//Delete cache
func (c *LocalModuleCache) Delete(deleteAll bool) error {
	files, err := ioutil.ReadDir(c.Configuration.CacheDir)
	if err != nil {
		return err
	}

	for _, f := range files {
		if f.IsDir() {
			if _, ok := c.Configuration.Modules[f.Name()]; deleteAll || ok {
				modulePath := path.Join(c.Configuration.CacheDir, f.Name())
				log.Printf("Removing module %s ...", modulePath)
				err := os.RemoveAll(modulePath)
				if err != nil {
					log.Printf("Error removing cached directory: %s", err)
				}
			}
		}
	}
	return nil
}
