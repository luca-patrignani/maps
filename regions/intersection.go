package regions

import (
	"errors"
	"math"

	mapset "github.com/deckarep/golang-set/v2"
	"github.com/luca-patrignani/maps/geometry"
)

func (r Region) containsAll(points []geometry.Point) bool {
	for _, p := range points {
		if !r.Contains(p) {
			return false
		}
	}
	return true
}

func (r Region) Intersection(other Region) (Region, error) {
	union := mapset.NewSet[geometry.Point]()
	for _, p := range append(r, other...) {
		if r.Contains(p) && other.Contains(p) {
			union.Add(p)
		}
	}
	union.Append(r.IntersectionPoints(other).ToSlice()...)
	segments := []geometry.Segment{}
	for _, side := range append(r.Sides(), other.Sides()...) {
		if union.Contains(side.P1) || union.Contains(side.P2) {
			segments = append(segments, side)
		}
	}
	if region, err := NewRegionFromSegments(segments); err == nil {
		return region, nil
	}
	return Region{}, errors.New("the regions are not intersecting")
}

func (r Region) fillAdj(points mapset.Set[geometry.Point], adj map[geometry.Point][]geometry.Point) {
	for _, side := range r.Sides() {
		if points.Contains(side.P1) && points.Contains(side.P2) {
			adj[side.P1] = append(adj[side.P1], side.P2)
			adj[side.P2] = append(adj[side.P1], side.P1)
		}
	}
}

func (r Region) IntersectionPoints(o Region) mapset.Set[geometry.Point] {
	intersections := mapset.NewSet[geometry.Point]()
	for _, rs := range r.Sides() {
		for _, os := range o.Sides() {
			if inter, err := geometry.Intersection(rs, os); err == nil {
				intersections.Add(inter)
			}
		}
	}
	return intersections
}

func (region Region) countIntersections(segment geometry.Segment) uint {
	var counter uint = 0
	for _, s := range region.Sides() {
		if inter, err := geometry.Intersection(segment, s); err == nil {
			if inter == s.P1 {

			}
			counter++
		}
	}
	return counter
}

func (r Region) Contains(p geometry.Point) bool {
	/*
	The issue is solved as follows: If the intersection point is a vertex of a tested polygon side, then the intersection counts only if the other vertex of the side lies below the ray. This is effectively equivalent to considering vertices on the ray as lying slightly above the ray.
	*/
	for _, vertex := range r {
		if p == vertex {
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
		if inter, err := geometry.Intersection(line, s); err == nil {
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
