package island

import (
	"encoding/json"
	"errors"
	"os"

	"github.com/spf13/afero"
)

type islandsJson struct {
	Islands  map[uint]Island
	Sequence uint
}

type IslandRepository struct {
	Fs       afero.Fs
	Filename string
	data     islandsJson
}

func (ir *IslandRepository) Save(island Island) (id uint, err error) {
	ir.data.Islands[ir.data.Sequence] = island
	id = ir.data.Sequence
	ir.data.Sequence++
	defer func() {
		if err != nil {
			delete(ir.data.Islands, ir.data.Sequence)
			ir.data.Sequence--
		}
	}()
	jsonR, err := json.MarshalIndent(ir.data, "", "  ")
	if err != nil {
		return 0, err
	}
	file, err := ir.Fs.Create(ir.Filename)
	if err != nil {
		return 0, err
	}
	defer file.Close()
	_, err = file.Write(jsonR)
	return 0, err
}

func NewIslandRepository(fs afero.Fs, filename string) (IslandRepository, error) {
	ir := IslandRepository{fs, filename, islandsJson{Islands: map[uint]Island{}}}
	data, err := afero.ReadFile(ir.Fs, ir.Filename)
	if err != nil {
		return IslandRepository{}, err
	}
	err = json.Unmarshal(data, &ir.data)
	if err != nil {
		return IslandRepository{}, err
	}
	return ir, nil
}

func InitIslandRepository(fs afero.Fs, filename string) (IslandRepository, error) {
	if _, err := fs.Stat(filename); !errors.Is(err, os.ErrNotExist) {
		return IslandRepository{}, err
	}
	file, err := fs.Create(filename)
	file.Close()
	if err != nil {
		return IslandRepository{}, err
	}
	return IslandRepository{fs, filename, islandsJson{Islands: map[uint]Island{}}}, nil
}

func (ir IslandRepository) Islands() map[uint]Island {
	return ir.data.Islands
}
