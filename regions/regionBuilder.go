package regions

import (
	"github.com/luca-patrignani/maps/geometry"
)

type RegionBuilder struct {
	Segments []geometry.Segment
	pendingPoint *geometry.Point
}

func (rb *RegionBuilder) AddPoint(newPoint geometry.Point) {
	if rb.pendingPoint != nil && newPoint != *rb.pendingPoint {
		rb.Segments = append(rb.Segments, geometry.Segment{P1: *rb.pendingPoint, P2: newPoint})
	}
	rb.pendingPoint = &newPoint
}

func (rb RegionBuilder) Build() (Region, error) {
	return NewRegionFromSegments(rb.Segments)
}