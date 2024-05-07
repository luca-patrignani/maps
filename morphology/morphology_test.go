package morphology

import (
	"bytes"
	"reflect"
	"testing"

	mapset "github.com/deckarep/golang-set/v2"
	"github.com/luca-patrignani/maps/geometry"
)

func TestDrawLine(t *testing.T) {
	m := New(0, 0, 0, 0)
	m.DrawLine(geometry.Point{X: 0, Y: 0}, geometry.Point{X: 0, Y: 5}, Land)
	expected := mapset.NewSet(
		geometry.Point{X: 0, Y: 0},
		geometry.Point{X: 0, Y: 1},
		geometry.Point{X: 0, Y: 2},
		geometry.Point{X: 0, Y: 3},
		geometry.Point{X: 0, Y: 4},
		geometry.Point{X: 0, Y: 5},
	)
	actual := mapset.NewSet[geometry.Point]()
	for p, t := range m.Data {
		if t == Land {
			actual.Add(p)
		}
	}
	if !expected.Equal(actual) {
		t.Fatal(expected, actual)
	}
}

func TestFillWith(t *testing.T) {
	m := New(0, 0, 2000, 2000)
	m.DrawLine(geometry.Point{X: 0, Y: 0}, geometry.Point{X: 5, Y: 0}, Land)
	m.DrawLine(geometry.Point{X: 5, Y: 0}, geometry.Point{X: 5, Y: 5}, Land)
	m.DrawLine(geometry.Point{X: 5, Y: 5}, geometry.Point{X: 0, Y: 5}, Land)
	m.DrawLine(geometry.Point{X: 0, Y: 5}, geometry.Point{X: 0, Y: 0}, Land)

	m.FillWith(geometry.Point{X: 1, Y: 1}, Land, Sea)

	expected := mapset.NewSet[geometry.Point]()
	for x := 0; x <= 5; x++ {
		for y := 0; y <= 5; y++ {
			expected.Add(geometry.Point{X: x, Y: y})
		}
	}
	actual := mapset.NewSet[geometry.Point]()
	for p, t := range m.Data {
		if t == Land {
			actual.Add(p)
		}
	}
	if !expected.Equal(actual) {
		t.Fatal(expected, actual)
	}

	m.FillWith(geometry.Point{X: 1, Y: 1}, Land, Sea)
	actual = mapset.NewSet[geometry.Point]()
	for p, t := range m.Data {
		if t == Land {
			actual.Add(p)
		}
	}
	if !expected.Equal(actual) {
		t.Fatal(expected, actual)
	}
}

func TestLoadSave(t *testing.T) {
	m1 := New(0, 0, 2000, 2000)
	m1.DrawLine(geometry.Point{X: 0, Y: 0}, geometry.Point{X: 5, Y: 0}, Land)
	m1.DrawLine(geometry.Point{X: 5, Y: 0}, geometry.Point{X: 5, Y: 5}, Land)
	m1.DrawLine(geometry.Point{X: 5, Y: 5}, geometry.Point{X: 0, Y: 5}, Land)
	m1.DrawLine(geometry.Point{X: 0, Y: 5}, geometry.Point{X: 0, Y: 0}, Land)
	rw := bytes.Buffer{}
	if err := m1.Save(&rw); err != nil {
		t.Fatal(err)
	}
	m2, err := NewFromFile(&rw)
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(m1, m2) {
		t.Fatal()
	}
}
