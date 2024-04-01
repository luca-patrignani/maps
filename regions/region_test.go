package regions

import (
	"testing"

	"github.com/luca-patrignani/maps/geometry"
)

func TestEqualsStrict(t *testing.T) {
	a := Region{geometry.Point{X: 0, Y: 0}, geometry.Point{X: 1, Y: 0}, geometry.Point{X: 1, Y: 1}}
	if !a.Equals(a) {
		t.Fatal()
	}
}

func TestEqualsGeneral(t *testing.T) {
	a := Region{geometry.Point{X: 0, Y: 0}, geometry.Point{X: 1, Y: 0}, geometry.Point{X: 1, Y: 1}}
	b := Region{geometry.Point{X: 1, Y: 1}, geometry.Point{X: 0, Y: 0}, geometry.Point{X: 1, Y: 0}}
	if !a.Equals(b) {
		t.Fatal()
	}
}

func TestEqualsDifferentLen(t *testing.T) {
	a := Region{geometry.Point{X: 0, Y: 0}, geometry.Point{X: 1, Y: 0}, geometry.Point{X: 1, Y: 1}}
	b := Region{geometry.Point{X: 0, Y: 0}, geometry.Point{X: 1, Y: 0}, geometry.Point{X: 1, Y: 1}, geometry.Point{X: 1, Y: 1}}
	if a.Equals(b) {
		t.Fatal()
	}
}

func TestEqualsReverse(t *testing.T) {
	a := Region{geometry.Point{X: 0, Y: 0}, geometry.Point{X: 1, Y: 0}, geometry.Point{X: 1, Y: 1}}
	b := Region{geometry.Point{X: 1, Y: 1}, geometry.Point{X: 1, Y: 0}, geometry.Point{X: 0, Y: 0}}
	if !a.Equals(b) {
		t.Fatal()
	}
}

func TestNewRegionFromSegmentsShortEdges(t *testing.T) {
	segments := []geometry.Segment{
		{P1: geometry.Point{X: 0, Y: 0}, P2: geometry.Point{X: 1, Y: 0}},
		{P1: geometry.Point{X: 1, Y: 0}, P2: geometry.Point{X: 1, Y: 1}},
		{P1: geometry.Point{X: 1, Y: 1}, P2: geometry.Point{X: 0, Y: 1}},
		{P1: geometry.Point{X: 0, Y: 1}, P2: geometry.Point{X: 0, Y: 0}},
	}
	actual, err := NewRegionFromSegments(segments)
	if err != nil {
		t.Fatal(err)
	}
	expected := Region{geometry.Point{X: 0, Y: 0}, geometry.Point{X: 1, Y: 0}, geometry.Point{X: 1, Y: 1}, geometry.Point{X: 0, Y: 1}}
	if !actual.Equals(expected) {
		t.Fatal(expected, actual)
	}
}

func TestNewRegionFromSegmentsLongEdges(t *testing.T) {
	segments := []geometry.Segment{
		{P1: geometry.Point{X: 0, Y: 0}, P2: geometry.Point{X: 10, Y: 0}},
		{P1: geometry.Point{X: 10, Y: 0}, P2: geometry.Point{X: 10, Y: 10}},
		{P1: geometry.Point{X: 10, Y: 10}, P2: geometry.Point{X: 0, Y: 10}},
		{P1: geometry.Point{X: 0, Y: 10}, P2: geometry.Point{X: 0, Y: -10}},
	}
	actual, err := NewRegionFromSegments(segments)
	if err != nil {
		t.Fatal(err)
	}
	expected := Region{
		geometry.Point{X: 0, Y: 0},
		geometry.Point{X: 10, Y: 0},
		geometry.Point{X: 10, Y: 10},
		geometry.Point{X: 0, Y: 10},
	}
	if !actual.Equals(expected) {
		t.Fatal(expected, actual)
	}
}
