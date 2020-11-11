package module

//TerraformModule is an implimentation of Module for Terraform Modules hosted in
//git repos
type TerraformModule struct {
	Repo   string `yaml:"repo,omitempty"`
	Branch string `yaml:"branch,omitempty"`
	Tag    string `yaml:"tag,omitempty"`
	Path   string `yaml:"path,omitempty"`
}

//Clone will clone the module and checkout the specified branch/tag
func (m *TerraformModule) Clone(cachePath string) error {
	return nil
}

//Update will update the local cache with the latest changes for the module branch/tag
func (m *TerraformModule) Update(cachePath string) error {
	return nil
}

//Delete removes the module from the cachePath if found
func (m *TerraformModule) Delete(cachePath string) error {
	return nil
}
