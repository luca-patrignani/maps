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
