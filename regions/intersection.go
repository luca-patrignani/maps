package regions

import (
	"errors"
	"math"

	"github.com/luca-patrignani/maps/geometry"
)

func (r Region) Intersection(other Region) (Region, error) {
	for _, region := range NewRegionsFromSegments(append(r.Sides(), other.Sides()...)) {
		found := true
		for _, p := range r.IntersectionPoints(other) {
			if !region.Contains(p) {
				found = false
				break
			}
		}
		if found {
			return region, nil
		}
	}
	return Region{}, errors.New("the regions are not intersecting")
}

func (r Region) IntersectionPoints(o Region) []geometry.Point {
	intersections := []geometry.Point{}
	for _, rs := range r.Sides() {
		for _, os := range o.Sides() {
			if inter, err := geometry.Intersection(rs, os); err == nil {
				intersections = append(intersections, inter)
			}
		}
	}
	return intersections
}

func countIntersection(segment geometry.Segment, segments []geometry.Segment) uint {
	var counter uint = 0
	for _, s := range segments {
		if _, err := geometry.Intersection(segment, s); err == nil {
			counter++
		}
	}
	return counter
}

func (r Region) Contains(p geometry.Point) bool {
	for _, vertex := range r {
		if p == vertex {
			return true
		}
	}
	line := geometry.Segment{
		P1: geometry.Point{
			X: p.X,
			Y: p.Y,
		}, P2: geometry.Point{
			X: math.MaxInt32,
			Y: p.Y,
		}}
	return countIntersection(line, r.Sides())%2 == 1
}
