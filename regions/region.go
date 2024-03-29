package regions

import (
	"errors"

	"github.com/luca-patrignani/maps/geometry"
)

type Region []geometry.Point

func (a Region) notReverseEquals(b Region) bool {
	if len(a) != len(b) {
		return false
	}
	offset := -1
	for j, bb := range b {
		if a[0] == bb {
			offset = j
		}
	}
	if offset == -1 {
		return false
	}
	for i := range a {
		if a[i] != b[(i+offset)%len(b)] {
			return false
		}
	}
	return true
}

func reverse(points []geometry.Point) []geometry.Point {
	for i, j := 0, len(points)-1; i < j; i, j = i+1, j-1 {
		points[i], points[j] = points[j], points[i]
	}
	return points
}

func (a Region) Equals(b Region) bool {
	return a.notReverseEquals(b) || a.notReverseEquals(reverse(b))
}

func NewRegion(points []geometry.Point) (Region, error) {
	edges := []geometry.Segment{}
	for i := 1; i < len(points); i++ {
		newEdge := geometry.Segment{P1: points[i-1], P2: points[i]}
		for j := 0; j < len(edges)-1; j++ {
			inter, err := geometry.Intersection(newEdge, edges[j])
			if err == nil {
				edges = append(edges, newEdge)
				region := Region{inter}
				for k := j; edges[k].P2 != newEdge.P2; k++ {
					region = append(region, edges[k].P2)
				}
				return region, nil
			}
		}
		edges = append(edges, newEdge)
	}
	return Region{}, errors.New("region is not closed")
}

func findCycle(edges map[geometry.Segment]bool) ([]geometry.Point, error) {
	adj := map[geometry.Point][]geometry.Point{}
	var src geometry.Point
	for edge := range edges {
		adj[edge.P1] = append(adj[edge.P1], edge.P2)
		adj[edge.P2] = append(adj[edge.P2], edge.P1)
		src = edge.P1
	}
	queue := []geometry.Point{src}
	visited := map[geometry.Point]bool{src: true}
	pred := map[geometry.Point]*geometry.Point{src: nil}
	for len(queue) > 0 {
		u := queue[0]
		queue = queue[1:]
		for _, v := range adj[u] {
			if pred[u] != nil && v == *pred[u] {
				continue
			}
			if !visited[v] {
				queue = append(queue, v)
				visited[v] = true
				pred[v] = &u
			} else {
				predsU := []geometry.Point{u}
				for uu := pred[u]; uu != nil; uu = pred[*uu] {
					predsU = append(predsU, *uu)
				}
				predsV := []geometry.Point{v}
				for vv := pred[v]; vv != nil; vv = pred[*vv] {
					predsV = append(predsV, *vv)
				}
				for i, uu := range predsU {
					for j, vv := range predsV {
						if uu == vv {
							return append(predsU[:i+1], reverse(predsV[:j])...), nil
						}
					}
				}	
			}
		}
	}

	return []geometry.Point{}, errors.New("cannot find cycle")
}

func NewRegionFromSegments(segments []geometry.Segment) (Region, error) {
	edges := map[geometry.Segment]bool{}
	for _, segment := range segments {
		edges[segment] = true
	}
	for i := 0; i < len(segments); i++ {
		for j := i+1; j < len(segments); j++ {
			if inter, err := geometry.Intersection(segments[i], segments[j]); err == nil {
				if segments[i].P1 != inter && segments[i].P2 != inter {
					edges[geometry.Segment{P1: segments[i].P1, P2: inter}] = true
					edges[geometry.Segment{P1: segments[i].P2, P2: inter}] = true
				}
				if segments[j].P1 != inter && segments[j].P2 != inter {
					edges[geometry.Segment{P1: segments[j].P1, P2: inter}] = true
					edges[geometry.Segment{P1: segments[j].P2, P2: inter}] = true	
				}
				if region, err := findCycle(edges); err == nil {
					return region, nil
				}
			}
		}
	}
	return Region{}, errors.New("region is not closed")
}
