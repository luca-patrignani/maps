package regions

import (
	"math"

	"github.com/luca-patrignani/maps/geometry"
)

func fillAdj(adj map[geometry.Point][]geometry.Point, region Region) map[geometry.Point][]geometry.Point {
	for i := 1; i < len(region); i++ {
		adj[region[i]] = append(adj[region[i]], region[i-1])
		adj[region[i-1]] = append(adj[region[i-1]], region[i])
	}
	adj[region[0]] = append(adj[region[0]], region[len(region)-1])
	adj[region[len(region)-1]] = append(adj[region[len(region)-1]], region[0])
	return adj
}

func (r Region) Intersection(other Region) (Region, error) {
	adj := map[geometry.Point][]geometry.Point{}
	adj = fillAdj(adj, r)
	adj = fillAdj(adj, other)
	src := r[0]
	_ = findCycles(adj, src)

	return Region{}, nil
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
