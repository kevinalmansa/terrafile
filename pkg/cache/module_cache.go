package cache

import (
	"terrafile/pkg/module"
)

type LocalModuleCache struct {
	modules []module.Module
}

func (c *LocalModuleCache) Clone() error {
	return nil
}

func (c *LocalModuleCache) Update() error {
	return nil
}

func (c *LocalModuleCache) Delete() error {
	return nil
}
