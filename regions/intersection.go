package regions

import (
	"errors"
	"math"

	mapset "github.com/deckarep/golang-set/v2"
	"github.com/luca-patrignani/maps/geometry"
)

func (r Region) Intersection(other Region) (Region, error) {
	allSegments := intersectSegments(append(r.Sides(), other.Sides()...))
	segments := []geometry.Segment{}
	for _, s := range allSegments {
		if r.Contains(s.P1) && other.Contains(s.P1) && r.Contains(s.P2) && other.Contains(s.P2) {
			segments = append(segments, s)
		}
	}
	if region, err := NewRegionFromSegments(segments); err == nil {
		return region, nil
	}
	return Region{}, errors.New("the regions are not intersecting")
}

func (r Region) IntersectionPoints(o Region) mapset.Set[geometry.Point] {
	intersections := mapset.NewSet[geometry.Point]()
	for _, rs := range r.Sides() {
		for _, os := range o.Sides() {
			if inter, ok := geometry.Intersection(rs, os).(geometry.Point); ok {
				intersections.Add(inter)
			}
		}
	}
	return intersections
}

func (r Region) Contains(p geometry.Point) bool {
	/*
	The issue is solved as follows: If the intersection point is a vertex of a tested polygon side, then the intersection counts only if the other vertex of the side lies below the ray. This is effectively equivalent to considering vertices on the ray as lying slightly above the ray.
	*/
	for _, side := range r.Sides() {
		if side.Contains(p) {
			return true
		}
	}
	line := geometry.Segment{
		P1: geometry.Point{
			X: p.X,
			Y: p.Y,
		},
		P2: geometry.Point{
			X: math.MaxInt32,
			Y: p.Y,
		},
	}
	var counter uint = 0
	for _, s := range r.Sides() {
		if inter, ok := geometry.Intersection(line, s).(geometry.Point); ok {
			if inter != s.P1 && inter != s.P2 {
				counter++
			} else {
				if inter == s.P1 && s.P2.Y < p.Y {
					counter++
				} else {
					if inter == s.P2 && s.P1.Y < p.Y {
						counter++
					}
				}
			}
		}
	}
	return counter%2 == 1
}
