package maps

import (
	"testing"

	"github.com/luca-patrignani/maps/geometry"
	"github.com/luca-patrignani/maps/regions"
	"github.com/spf13/afero"
)

func TestRegionRepositoryMultipleSaves(t *testing.T) {
	expected := []regions.Region{
		{
			geometry.Point{X: 0, Y: 0},
			geometry.Point{X: 10, Y: 0},
			geometry.Point{X: 10, Y: 10},
			geometry.Point{X: 0, Y: 10},
		},
		{
			geometry.Point{X: 0, Y: 0},
			geometry.Point{X: 1, Y: 0},
			geometry.Point{X: 1, Y: 1},
			geometry.Point{X: 0, Y: 1},
		},
	}
	fs := afero.NewMemMapFs()
	rr := RegionRepository{fs, "test.json"}

	if err := rr.Save(expected); err != nil {
		t.Fatal(err)
	}
	actual, err := rr.Load()
	if err != nil {
		t.Fatal(err)
	}
	if len(actual) != len(expected) {
		t.Fatal("expected and actual are not of the same size")
	}
	for i := range actual {
		if !actual[i].Equals(expected[i]) {
			t.Fail()
		}
	}

	if err := rr.Save(expected); err != nil {
		t.Fatal(err)
	}
	actual, err = rr.Load()
	if err != nil {
		t.Fatal(err)
	}
	if len(actual) != len(expected) {
		t.Fatal("expected and actual are not of the same size")
	}
	for i := range actual {
		if !actual[i].Equals(expected[i]) {
			t.Fail()
		}
	}
}
