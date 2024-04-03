package regions

import (
	"testing"

	"github.com/luca-patrignani/maps/geometry"
)

func TestIntersection(t *testing.T) {
	r1 := Region{
		geometry.Point{X: 0, Y: 0},
		geometry.Point{X: 10, Y: 0},
		geometry.Point{X: 10, Y: 10},
		geometry.Point{X: 0, Y: 10},
	}
	r2 := Region{
		geometry.Point{X: 5, Y: 5},
		geometry.Point{X: 15, Y: 5},
		geometry.Point{X: 10, Y: 10},
		geometry.Point{X: 0, Y: 10},
	}
	expected := Region{
		geometry.Point{X: 5, Y: 5},
		geometry.Point{X: 10, Y: 5},
		geometry.Point{X: 10, Y: 10},
		geometry.Point{X: 5, Y: 10},
	}
	actual, err := r1.Intersection(r2)
	if err != nil {
		t.Fatal(err)
	}
	if !actual.Equals(expected) {
		t.Fatal("regions are not equal")
	}
}