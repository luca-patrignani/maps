package geometry

import (
	"reflect"
	"testing"
)

func TestNewRegionSameEnd(t *testing.T) {
	region, err := NewRegion([]Point{{0, 0}, {0, 3}, {3, 3}, {3, 0}, {0, 0}})
	if err != nil {
		t.Error(err)
	}
	if !reflect.DeepEqual(region, Region{{0, 0}, {0, 3}, {3, 3}, {3, 0}}) {
		t.Fail()
	}
}

func TestNewRegionIntersectionEnd(t *testing.T) {
	region, err := NewRegion([]Point{{0, -2}, {0, 3}, {3, 3}, {3, 0}, {0, 0}})
	if err != nil {
		t.Error(err)
	}
	if !reflect.DeepEqual(region, Region{{0, 0}, {0, 3}, {3, 3}, {3, 0}}) {
		t.Fail()
	}
}

func TestNewRegionIntersectionMiddle(t *testing.T) {
	region, err := NewRegion([]Point{{0, -2}, {0, 3}, {3, 3}, {3, 0}, {0, 0}, {10, 10}, {20, 20}})
	if err != nil {
		t.Error(err)
	}
	if !reflect.DeepEqual(region, Region{{0, 0}, {0, 3}, {3, 3}, {3, 0}}) {
		t.Fail()
	}
}

func TestNewRegionOpen(t *testing.T) {
	_, err := NewRegion([]Point{{0, 0}, {0, 3}, {3, 3}, {3, 0}})
	if err == nil {
		t.Error("region is supposed to be open")
	}
}

func TestNewRegion(t *testing.T) {
	_, err := NewRegion([]Point{{15, 20}, {16, 20}, {17, 20}, {18, 20}})
	if err == nil {
		t.Error("region is supposed to be open")
	}
}