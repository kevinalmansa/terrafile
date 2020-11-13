package module

import (
	"fmt"
	"os"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
)

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
	_, err := git.PlainClone(cachePath, false, &git.CloneOptions{
		URL:           m.Repo,
		Progress:      os.Stdout,
		ReferenceName: plumbing.ReferenceName(fmt.Sprintf("refs/heads/%s", m.Branch)),
	})
	return err
}

//Update will update the local cache with the latest changes for the module branch/tag
func (m *TerraformModule) Update(cachePath string) error {
	repo, err := git.PlainOpen(cachePath)
	if err != nil {
		return err
	}
	w, err := repo.Worktree()
	if err != nil {
		return err
	}
	err = w.Pull(&git.PullOptions{RemoteName: "origin"})
	return err
}
