package politics

import (
	"bytes"
	"testing"

	mapset "github.com/deckarep/golang-set/v2"
	"github.com/luca-patrignani/maps/geometry"
	"github.com/luca-patrignani/maps/morphology"
)

const italy PoliticalEntity = "Italy"

func newMorphSquare() *morphology.Morphology {
	m := morphology.New(0, 0, 2000, 2000)
	m.DrawSquare(geometry.Point{X: 0, Y: 0}, 5, morphology.Land)
	return &m
}

func TestDrawLine(t *testing.T) {
	pm := PoliticalMap{
		Morphology: newMorphSquare(),
		Data:       map[geometry.Point]PoliticalEntity{},
	}
	pm.DrawLine(geometry.Point{X: 3, Y: -2}, geometry.Point{X: 3, Y: 7}, italy)
	expected := mapset.NewSet(
		geometry.Point{X: 3, Y: 0},
		geometry.Point{X: 3, Y: 1},
		geometry.Point{X: 3, Y: 2},
		geometry.Point{X: 3, Y: 3},
		geometry.Point{X: 3, Y: 4},
		geometry.Point{X: 3, Y: 5},
	)
	actual := mapset.NewSet[geometry.Point]()
	for p, e := range pm.Data {
		if e == italy {
			actual.Add(p)
		}
	}
	if !expected.Equal(actual) {
		t.Fatal(expected, actual)
	}
}

func TestFillWith(t *testing.T) {
	pm := PoliticalMap{
		Morphology: newMorphSquare(),
		Data:       map[geometry.Point]PoliticalEntity{},
	}
	pm.DrawLine(geometry.Point{X: 3, Y: 0}, geometry.Point{X: 3, Y: 5}, italy)
	//fill the right side
	pm.FillWith(geometry.Point{X: 1, Y: 1}, italy, None)
	actual := mapset.NewSet[geometry.Point]()
	for p, e := range pm.Data {
		if e == italy {
			actual.Add(p)
		}
	}
	for x := 0; x <= 3; x++ {
		for y := 0; y <= 5; y++ {
			if !actual.ContainsOne(geometry.Point{X: x, Y: y}) {
				t.Fatal(x, y)
			}
		}
	}
}

func TestLoadSave(t *testing.T) {
	pm := PoliticalMap{
		Morphology: newMorphSquare(),
		Data:       map[geometry.Point]PoliticalEntity{},
	}
	pm.DrawLine(geometry.Point{X: 3, Y: 0}, geometry.Point{X: 3, Y: 5}, italy)
	pm.FillWith(geometry.Point{X: 1, Y: 1}, italy, None)
	b := bytes.Buffer{}
	if err := pm.Save(&b); err != nil {
		t.Fatal(err)
	}
	pm2, err := NewFromFile(&b, pm.Morphology)
	if err != nil {
		t.Fatal(err)
	}
	for p, e := range pm2.Data {
		if pm.Data[p] != e {
			t.Fatal()
		}
	}
}
