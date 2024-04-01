package maps

import (
	"encoding/json"
	"time"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/object"
	"github.com/luca-patrignani/maps/regions"
	"github.com/spf13/afero"
)

type RegionRepository struct {
	Fs       afero.Fs
	Filename string
}

func (rr RegionRepository) Save(regions []regions.Region) error {
	jsonR, err := json.MarshalIndent(regions, "", "  ")
	if err != nil {
		return err
	}
	file, err := rr.Fs.Create(rr.Filename)
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = file.Write(jsonR)
	return err
}

func (rr RegionRepository) Load() ([]regions.Region, error) {
	data, err := afero.ReadFile(rr.Fs, rr.Filename)
	if err != nil {
		return []regions.Region{}, err
	}
	rs := []regions.Region{}
	err = json.Unmarshal(data, &rs)
	if err != nil {
		return []regions.Region{}, err
	}
	return rs, nil
}

func (rr RegionRepository) Commit(message string) error {
	gitRepo, err := git.PlainOpen("./")
	if err != nil {
		return err
	}
	worktree, err := gitRepo.Worktree()
	if err != nil {
		return err
	}
	if _, err := worktree.Add("."); err != nil {
		return err
	}
	_, err = worktree.Commit(message, &git.CommitOptions{
		Committer: &object.Signature{
			Name: "maps-repo-sys",
			When: time.Now(),
		},
	})
	return err
}
