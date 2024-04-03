package island

import (
	"testing"

	"github.com/luca-patrignani/maps/geometry"
	"github.com/luca-patrignani/maps/regions"
	"github.com/spf13/afero"
)

func TestIslandRepositoryMultipleSaves(t *testing.T) {
	expected := map[uint]Island{
		0: {
			Name:   "First Island",
			Region: regions.Region{
				geometry.Point{X: 0, Y: 0},
				geometry.Point{X: 10, Y: 0},
				geometry.Point{X: 10, Y: 10},
				geometry.Point{X: 0, Y: 10},
			},
		},
		1: {
			Name:   "Second Island",
			Region: regions.Region{
				geometry.Point{X: 10, Y: 0},
				geometry.Point{X: 20, Y: 0},
				geometry.Point{X: 20, Y: 10},
				geometry.Point{X: 10, Y: 10},
			},
		},
	}
	
	fs := afero.NewMemMapFs()
	ir, err := InitIslandRepository(fs, "test.json")
	if err != nil {
		t.Fatal(err)
	}
	for _, island := range expected {
		if _, err := ir.Save(island); err != nil {
			t.Fatal(err)
		}
	}

	actual := ir.Islands()
	if len(actual) != len(expected) {
		t.Fatal("expected and actual are not of the same size")
	}
	for id := range actual {
		if actual[id].Name != expected[id].Name {
			t.Fail()
		}
		if !actual[id].Region.Equals(expected[id].Region) {
			t.Fail()
		}
	}

	ir, err = NewIslandRepository(fs, "test.json")
	if err != nil {
		t.Fatal(err)
	}
	expected[2] = Island{
		Name:   "Third Island",
		Region: regions.Region{
			geometry.Point{X: 100, Y: 100},
			geometry.Point{X: 100, Y: 200},
			geometry.Point{X: 200, Y: 200},
		},
	}
	if _, err := ir.Save(expected[2]); err != nil {
		t.Fatal(err)
	}
	actual = ir.Islands()
	if len(actual) != len(expected) {
		t.Fatal("expected and actual are not of the same size")
	}
	for id := range actual {
		if actual[id].Name != expected[id].Name {
			t.Fail()
		}
		if !actual[id].Region.Equals(expected[id].Region) {
			t.Fail()
		}
	}
}

