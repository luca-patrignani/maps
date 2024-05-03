package bresenham

import (
	"testing"

	mapset "github.com/deckarep/golang-set/v2"
	"github.com/luca-patrignani/maps/geometry"
)

func TestBresenham(t *testing.T) {
	p0 := geometry.Point{X: 0, Y: 1}
	p1 := geometry.Point{X: 6, Y: 4}
	expected := mapset.NewSet(
		geometry.Point{X: 0, Y: 1},
		geometry.Point{X: 1, Y: 1},
		geometry.Point{X: 2, Y: 2},
		geometry.Point{X: 3, Y: 2},
		geometry.Point{X: 4, Y: 3},
		geometry.Point{X: 5, Y: 3},
		geometry.Point{X: 6, Y: 4},
	)
	actual := mapset.NewSet(Bresenham(p0, p1)...)
	if !expected.Equal(actual) {
		t.Fatal()
	}
}

func TestBresenham2(t *testing.T) {
	p0 := geometry.Point{X: 6, Y: 4}
	p1 := geometry.Point{X: 0, Y: 1}
	expected := mapset.NewSet(
		geometry.Point{X: 0, Y: 1},
		geometry.Point{X: 1, Y: 1},
		geometry.Point{X: 2, Y: 2},
		geometry.Point{X: 3, Y: 2},
		geometry.Point{X: 4, Y: 3},
		geometry.Point{X: 5, Y: 3},
		geometry.Point{X: 6, Y: 4},
	)
	actual := mapset.NewSet(Bresenham(p0, p1)...)
	if !expected.Equal(actual) {
		t.Fatal(expected, actual)
	}
}

func TestBresenham3(t *testing.T) {
	p0 := geometry.Point{X: 0, Y: 5}
	p1 := geometry.Point{X: 5, Y: 0}
	expected := mapset.NewSet(
		geometry.Point{X: 0, Y: 5},
		geometry.Point{X: 1, Y: 4},
		geometry.Point{X: 2, Y: 3},
		geometry.Point{X: 3, Y: 2},
		geometry.Point{X: 4, Y: 1},
		geometry.Point{X: 5, Y: 0},
	)
	actual := mapset.NewSet(Bresenham(p0, p1)...)
	if !expected.Equal(actual) {
		t.Fatal(expected, actual)
	}
}

func TestBresenhamVertical(t *testing.T) {
	p0 := geometry.Point{X: 0, Y: 0}
	p1 := geometry.Point{X: 0, Y: 5}
	expected := mapset.NewSet(
		geometry.Point{X: 0, Y: 0},
		geometry.Point{X: 0, Y: 1},
		geometry.Point{X: 0, Y: 2},
		geometry.Point{X: 0, Y: 3},
		geometry.Point{X: 0, Y: 4},
		geometry.Point{X: 0, Y: 5},
	)
	actual := mapset.NewSet(Bresenham(p0, p1)...)
	if !expected.Equal(actual) {
		t.Fatal(expected, actual)
	}
}
