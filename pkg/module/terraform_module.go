package module

import (
	"errors"
	"log"
	"os"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/config"
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
		ReferenceName: plumbing.NewBranchReferenceName(m.Branch),
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
	targetBranch := plumbing.NewBranchReferenceName(m.Branch)

	repo, err := git.PlainOpenWithOptions(cachePath, &git.PlainOpenOptions{DetectDotGit: true})
	if err != nil {
		log.Printf("Error opening repo - %s", err)
		return err
	}

	//Fetch to synchornize with remote
	remote, err := repo.Remote("origin")
	if err != nil {
		log.Printf("Error getting remote \"origin\" - %s", err)
		return err
	}
	err = remote.Fetch(&git.FetchOptions{
		Progress: os.Stdout,
		// RefSpecs: []config.RefSpec{"refs/*:refs/*", "HEAD:refs/heads/HEAD"},
		RefSpecs: []config.RefSpec{"refs/*:refs/*"},
	})
	if err != nil && !errors.Is(err, git.NoErrAlreadyUpToDate) {
		log.Printf("Error fetching repo - %s", err)
		return err
	}

	//Checkout
	w, err := repo.Worktree()
	if err != nil {
		return err
	}
	currentBranch, err := repo.Head()
	if err != nil {
		log.Printf("Error retrieving current HEAD - %s", err)
		return err
	}
	if currentBranch.Name() != targetBranch {
		err = w.Checkout(&git.CheckoutOptions{
			Branch: targetBranch,
		})
		if err != nil {
			log.Printf("Error checking out branch %s - %s", m.Branch, err)
		}
	}

	//Pull latest updates
	err = w.Pull(&git.PullOptions{
		RemoteName:    "origin",
		ReferenceName: targetBranch,
	})
	if err != nil && errors.Is(err, git.NoErrAlreadyUpToDate) {
		log.Printf("Module %s already up-to-date", cachePath)
		err = nil //ignore already up to date "error"
	} else if err != nil {
		log.Printf("Error pulling repo - %s", err)
		return err
	}

	//Checkout tag
	if err == nil && m.Tag != "" {
		err = m.checkoutTag(repo)
	}
	return err
}
