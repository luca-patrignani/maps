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
		geometry.Point{X: 15, Y: 15},
		geometry.Point{X: 5, Y: 15},
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
		t.Fatalf("expected %v and actual %v are not equal", expected, actual)
	}
}

func TestIntersectionPoints(t *testing.T) {
	r1 := Region{
		geometry.Point{X: 0, Y: 0},
		geometry.Point{X: 10, Y: 0},
		geometry.Point{X: 10, Y: 10},
		geometry.Point{X: 0, Y: 10},
	}
	r2 := Region{
		geometry.Point{X: 5, Y: 5},
		geometry.Point{X: 15, Y: 5},
		geometry.Point{X: 15, Y: 15},
		geometry.Point{X: 5, Y: 15},
	}
	expected := []geometry.Point{{X: 10, Y: 5}, {X: 5, Y: 10}}
	actual := r1.IntersectionPoints(r2)
	for i := range actual {
		if actual[i] != expected[i] {
			t.Fatal()
		}
	}
}

func TestContains(t *testing.T) {
	r := Region{
		geometry.Point{X: 0, Y: 0},
		geometry.Point{X: 10, Y: 0},
		geometry.Point{X: 10, Y: 10},
		geometry.Point{X: 0, Y: 10},
	}
	p := geometry.Point{X: 2, Y: 2}
	if !r.Contains(p) {
		t.Fatal()
	}
}

func TestNotContains(t *testing.T) {
	r := Region{
		geometry.Point{X: 0, Y: 0},
		geometry.Point{X: 10, Y: 0},
		geometry.Point{X: 10, Y: 10},
		geometry.Point{X: 0, Y: 10},
	}
	p := geometry.Point{X: 20, Y: 42}
	if r.Contains(p) {
		t.Fatal()
	}
}

func TestContainsVertices(t *testing.T) {
	r := Region{
		geometry.Point{X: 0, Y: 0},
		geometry.Point{X: 10, Y: 0},
		geometry.Point{X: 10, Y: 10},
		geometry.Point{X: 0, Y: 10},
	}
	p := geometry.Point{X: 0, Y: 0}
	if !r.Contains(p) {
		t.Fatal()
	}
}

func TestContainsVertices2(t *testing.T) {
	r := Region{
		geometry.Point{X: 5, Y: 5},
		geometry.Point{X: 10, Y: 5},
		geometry.Point{X: 10, Y: 10},
		geometry.Point{X: 5, Y: 10},
	}
	p := geometry.Point{X: 10, Y: 5}
	if !r.Contains(p) {
		t.Fatal()
	}
}
