package maps

import (
	"encoding/json"

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
