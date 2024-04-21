package regions

import (
	"testing"

	"github.com/luca-patrignani/maps/geometry"
)

func TestRegionBuilder(t *testing.T) {
	rb := RegionBuilder{}
	rb.AddPoint(geometry.Point{X: 0, Y: 0})
	if _, err := rb.Build(); err == nil {
		t.Fail()
	}
	rb.AddPoint(geometry.Point{X: 1, Y: 0})
	if _, err := rb.Build(); err == nil {
		t.Fail()
	}
	rb.AddPoint(geometry.Point{X: 1, Y: 1})
	if _, err := rb.Build(); err == nil {
		t.Fail()
	}
	rb.AddPoint(geometry.Point{X: 0, Y: 1})
	if _, err := rb.Build(); err == nil {
		t.Fail()
	}
	rb.AddPoint(geometry.Point{X: 0, Y: 0})
	actual, err := rb.Build()
	if err != nil {
		t.Fatal(err)
	}
	expected := Region{geometry.Point{X: 0, Y: 0}, geometry.Point{X: 1, Y: 0}, geometry.Point{X: 1, Y: 1}, geometry.Point{X: 0, Y: 1}}
	if !actual.Equals(expected) {
		t.Fatal(expected, actual)
	}
}

func TestWeld(t *testing.T) {
	s1 := geometry.Segment{P1: geometry.Point{X: 0, Y: 0}, P2: geometry.Point{X: 10, Y: 0}}
	s2 := geometry.Segment{P1: geometry.Point{X: 8, Y: 0}, P2: geometry.Point{X: 20, Y: 0}}
	welded, err := weld(s1, s2)
	if err != nil {
		t.Fatal(err)
	}
	expected := geometry.Segment{P1: geometry.Point{X: 0, Y: 0}, P2: geometry.Point{X: 20, Y: 0}}
	if welded != expected {
		t.Fatal(expected, welded)
	}
}

func (rb *RegionBuilder) addAll(points []geometry.Point) {
	for _, point := range points {
		rb.AddPoint(point)
	}
}

func TestBuildWelded(t *testing.T) {
	rb := RegionBuilder{}
	rb.addAll([]geometry.Point{
		{X: 0, Y: 0},
		{X: 1, Y: 0},
		{X: 2, Y: 0},
		{X: 2, Y: 1},
		{X: 2, Y: 2},
		{X: 1, Y: 2},
		{X: 0, Y: 2},
		{X: 0, Y: 1},
		{X: 0, Y: 0},
		{X: 0, Y: -1},
	})
	expected := Region{
		geometry.Point{X: 0, Y: 0},
		geometry.Point{X: 2, Y: 0},
		geometry.Point{X: 2, Y: 2},
		geometry.Point{X: 0, Y: 2},
	}
	actual, err := rb.Build()
	if err != nil {
		t.Fatal(err)
	}
	if !expected.Equals(actual) {
		t.Fatal(expected, actual)
	}
}

func TestRB(t *testing.T) {
	rb := RegionBuilder{}
	rb.addAll([]geometry.Point{
		{X: 2, Y: 0},
		{X: 2, Y: 2},
		{X: 0, Y: 2},
		{X: 0, Y: 1},
		{X: 5, Y: 1},
	})
	actual, err := rb.Build()
	if err != nil {
		t.Fatal(err)
	}
	expected := Region{
		geometry.Point{X: 2, Y: 1},
		geometry.Point{X: 2, Y: 2},
		geometry.Point{X: 0, Y: 2},
		geometry.Point{X: 0, Y: 1},
	}
	if !actual.Equals(expected) {
		t.Fatal(expected, actual)
	}
}
