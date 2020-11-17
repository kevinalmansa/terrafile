package module

import (
	"errors"
	"fmt"
	"log"
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

func (m *TerraformModule) checkoutTag(repo *git.Repository) error {
	ref, err := repo.Tag(m.Tag)
	if err != nil {
		return err
	}
	w, err := repo.Worktree()
	if err != nil {
		return err
	}
	err = w.Checkout(&git.CheckoutOptions{
		Hash: ref.Hash(),
	})
	return err
}

//Clone will clone the module and checkout the specified branch/tag
func (m *TerraformModule) Clone(cachePath string) error {
	log.Printf("Downloading %s to %s...", m.Repo, cachePath)
	repo, err := git.PlainClone(cachePath, false, &git.CloneOptions{
		URL:           m.Repo,
		Progress:      os.Stdout,
		ReferenceName: plumbing.ReferenceName(fmt.Sprintf("refs/heads/%s", m.Branch)),
	})
	if err != nil {
		return err
	}
	if m.Tag != "" {
		err = m.checkoutTag(repo)
	}
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
	err = w.Checkout(&git.CheckoutOptions{
		Branch: plumbing.ReferenceName(fmt.Sprintf("refs/heads/%s", m.Branch)),
	})
	if err != nil {
		log.Printf("Error checking out branch: %s", err)
	}
	err = w.Pull(&git.PullOptions{RemoteName: "origin"})
	if err != nil && errors.Is(err, git.NoErrAlreadyUpToDate) {
		err = nil //ignore already up to date "error"
	}
	if err == nil && m.Tag != "" {
		err = m.checkoutTag(repo)
	}
	return err
}
