package cache

import (
	"github.com/kevinalmansa/terrafile/pkg/module"
)

//LocalModuleCache is an implimentation of the Module interface for creating a
//local cache to store modules.
type LocalModuleCache struct {
	modules []module.Module
}

//Clone all modules, storing them in the cache
func (c *LocalModuleCache) Clone() error {
	return nil
}

//Update cache for all modules, cloning first if they don't exist
func (c *LocalModuleCache) Update() error {
	return nil
}

//Delete cache
func (c *LocalModuleCache) Delete() error {
	return nil
}
