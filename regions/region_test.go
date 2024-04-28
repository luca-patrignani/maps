package regions

import (
	"testing"

	mapset "github.com/deckarep/golang-set/v2"
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

func TestNewRegionFromSegmentsUpRightEnd(t *testing.T) {
	segments := []geometry.Segment{
		{P1: geometry.Point{X: 2, Y: 0}, P2: geometry.Point{X: 1, Y: 0}},
		{P1: geometry.Point{X: 1, Y: 0}, P2: geometry.Point{X: 0, Y: 0}},

		{P1: geometry.Point{X: 0, Y: 0}, P2: geometry.Point{X: 0, Y: 1}},
		{P1: geometry.Point{X: 0, Y: 1}, P2: geometry.Point{X: 0, Y: 2}},

		{P1: geometry.Point{X: 0, Y: 2}, P2: geometry.Point{X: 1, Y: 2}},
		{P1: geometry.Point{X: 1, Y: 2}, P2: geometry.Point{X: 2, Y: 2}},

		{P1: geometry.Point{X: 2, Y: 2}, P2: geometry.Point{X: 2, Y: 1}},
		{P1: geometry.Point{X: 2, Y: 1}, P2: geometry.Point{X: 2, Y: 0}},
		{P1: geometry.Point{X: 2, Y: 0}, P2: geometry.Point{X: 2, Y: -1}},
	}
	actual, err := NewRegionFromSegments(segments)
	if err != nil {
		t.Fatal(err)
	}
	expected := Region{
		geometry.Point{X: 0, Y: 0},
		geometry.Point{X: 1, Y: 0},
		geometry.Point{X: 2, Y: 0},
		geometry.Point{X: 2, Y: 1},
		geometry.Point{X: 2, Y: 2},
		geometry.Point{X: 1, Y: 2},
		geometry.Point{X: 0, Y: 2},
		geometry.Point{X: 0, Y: 1},
	}
	if !actual.Equals(expected) {
		t.Fatal(expected, actual)
	}
}

func TestOpenSpiral(t *testing.T) {
	segments := []geometry.Segment{
		{P1: geometry.Point{X: 27, Y: 3}, P2: geometry.Point{X: 30, Y: 3}},
		{P1: geometry.Point{X: 30, Y: 3}, P2: geometry.Point{X: 30, Y: 4}},
		{P1: geometry.Point{X: 30, Y: 4}, P2: geometry.Point{X: 31, Y: 4}},
		{P1: geometry.Point{X: 31, Y: 4}, P2: geometry.Point{X: 31, Y: 5}},
		{P1: geometry.Point{X: 31, Y: 5}, P2: geometry.Point{X: 32, Y: 5}},
		{P1: geometry.Point{X: 32, Y: 5}, P2: geometry.Point{X: 32, Y: 7}},
		{P1: geometry.Point{X: 32, Y: 7}, P2: geometry.Point{X: 31, Y: 7}},
		{P1: geometry.Point{X: 31, Y: 7}, P2: geometry.Point{X: 31, Y: 8}},
		{P1: geometry.Point{X: 31, Y: 8}, P2: geometry.Point{X: 30, Y: 9}},
		{P1: geometry.Point{X: 30, Y: 9}, P2: geometry.Point{X: 28, Y: 9}},
		{P1: geometry.Point{X: 28, Y: 9}, P2: geometry.Point{X: 28, Y: 8}},
		{P1: geometry.Point{X: 28, Y: 8}, P2: geometry.Point{X: 27, Y: 8}},
		{P1: geometry.Point{X: 28, Y: 8}, P2: geometry.Point{X: 28, Y: 7}},
		{P1: geometry.Point{X: 28, Y: 7}, P2: geometry.Point{X: 29, Y: 7}},
	}
	_, err := NewRegionFromSegments(segments)
	if err == nil {
		t.Fatal()
	}
}

func TestSides(t *testing.T) {
	r := Region{
		geometry.Point{X: 0, Y: 0},
		geometry.Point{X: 10, Y: 0},
		geometry.Point{X: 10, Y: 10},
		geometry.Point{X: 0, Y: 10},
	}
	expected := []geometry.Segment{
		{P1: geometry.Point{X: 0, Y: 0}, P2: geometry.Point{X: 10, Y: 0}},
		{P1: geometry.Point{X: 10, Y: 0}, P2: geometry.Point{X: 10, Y: 10}},
		{P1: geometry.Point{X: 10, Y: 10}, P2: geometry.Point{X: 0, Y: 10}},
		{P1: geometry.Point{X: 0, Y: 10}, P2: geometry.Point{X: 0, Y: 0}},
	}
	actual := r.Sides()
	if len(actual) != len(expected) {
		t.Fatal()
	}
	for i := range actual {
		if actual[i] != expected[i] {
			t.Fatal(actual[i], "!=", expected[i])
		}
	}
}

func TestIntersectSegments(t *testing.T) {
	segments := []geometry.Segment{
		{P1: geometry.Point{X: 0, Y: 0}, P2: geometry.Point{X: 10, Y: 0}},
		{P1: geometry.Point{X: 9, Y: 0}, P2: geometry.Point{X: 10, Y: 0}},
	}
	expected := mapset.NewSet(
		geometry.Segment{P1: geometry.Point{X: 0, Y: 0}, P2: geometry.Point{X: 9, Y: 0}},
		geometry.Segment{P1: geometry.Point{X: 9, Y: 0}, P2: geometry.Point{X: 10, Y: 0}},
	)
	actual := mapset.NewSet(intersectSegments(segments)...)
	if !expected.Equal(actual) {
		t.Fatal(expected, actual)
	}
}
