package regions

import (
	"errors"

	"github.com/luca-patrignani/maps/geometry"
)

type RegionBuilder struct {
	Segments     []geometry.Segment
	pendingPoint *geometry.Point
}

func (rb *RegionBuilder) AddPoint(newPoint geometry.Point) {
	if rb.pendingPoint != nil && newPoint != *rb.pendingPoint {
		rb.Segments = append(rb.Segments, geometry.Segment{P1: *rb.pendingPoint, P2: newPoint})
	}
	rb.pendingPoint = &newPoint
}

func (rb RegionBuilder) Build() (Region, error) {
	if len(rb.Segments) < 3 {
		return Region{}, errors.New("not enough segments for closing a region")
	}
	simplified := rb.Simplify()
	return NewRegionFromSegments(simplified)
}

func (rb RegionBuilder) Simplify() []geometry.Segment {
	simplified := []geometry.Segment{rb.Segments[0]}
	for _, s := range rb.Segments[1:] {
		if welded, err := weld(simplified[len(simplified)-1], s); err == nil {
			simplified[len(simplified)-1] = welded
		} else {
			simplified = append(simplified, s)
		}
	}
	return simplified
}

func weld(s1 geometry.Segment, s2 geometry.Segment) (geometry.Segment, error) {
	if geometry.Intersection(s1, s2) != nil && s1.IsParallelTo(s2) {
		segments := []geometry.Segment{
			{P1: s1.P1, P2: s2.P1},
			{P1: s1.P1, P2: s2.P2},
			{P1: s1.P2, P2: s2.P1},
			{P1: s1.P2, P2: s2.P2},
		}
		ans := segments[0]
		for _, s := range segments[1:] {
			if s.Length() > ans.Length() {
				ans = s
			}
		}
		return ans, nil
	}
	return geometry.Segment{}, errors.New("cannot fuze these segments")
}
