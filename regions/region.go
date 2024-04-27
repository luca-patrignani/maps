package regions

import (
	"errors"
	"fmt"

	mapset "github.com/deckarep/golang-set/v2"
	"github.com/luca-patrignani/maps/geometry"
)

type Region []geometry.Point

func (r Region) Sides() []geometry.Segment {
	sides := []geometry.Segment{}
	for i := 0; i < len(r)-1; i++ {
		sides = append(sides, geometry.Segment{P1: r[i], P2: r[i+1]})
	}
	sides = append(sides, geometry.Segment{P1: r[len(r)-1], P2: r[0]})
	return sides
}

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
			if inter, ok := geometry.Intersection(newEdge, edges[j]).(geometry.Point); ok {
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

func findCycle(adj map[geometry.Point][]geometry.Point, src geometry.Point) ([]geometry.Point, error) {
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
							region := append(predsU[:i+1], reverse(predsV[:j])...)
							if len(region) < 3 {
								fmt.Println(region)
							}
							return region, nil
						}
					}
				}
			}
		}
	}
	return []geometry.Point{}, errors.New("cannot find cycle")
}

func intersectSegments(segments []geometry.Segment) []geometry.Segment {
	edges := mapset.NewSet(segments...)
	result := []geometry.Segment{}
	for edges.Cardinality() > 0 {
		si := edges.ToSlice()[0]
		for _, sj := range edges.ToSlice() {
			if si != sj {
				switch inter := geometry.Intersection(si, sj).(type) {
				case geometry.Point:
					if si.P1 != inter && si.P2 != inter {
						edges.Remove(si)
						edges.Append(geometry.Segment{P1: si.P1, P2: inter}, geometry.Segment{P1: si.P2, P2: inter})
					}
					if sj.P1 != inter && sj.P2 != inter {
						edges.Remove(sj)
						edges.Append(geometry.Segment{P1: sj.P1, P2: inter}, geometry.Segment{P1: sj.P2, P2: inter})
					}
				case geometry.Segment:
					edges.Remove(si)
					edges.Remove(sj)
					edges.Add(inter)
					points := []geometry.Point{si.P1, si.P2, sj.P1, sj.P2}					
					for _, p := range points {
						if p != inter.P1 && p != inter.P2 {
							if p.Distance(inter.P1) < p.Distance(inter.P2) {
								edges.Add(geometry.Segment{P1: p, P2: inter.P1})
							} else {
								edges.Add(geometry.Segment{P1: p, P2: inter.P2})
							}
						}
					}
				}
			}
			if !edges.Contains(si) {
				break
			}
		}
		if edges.Contains(si) {
			edges.Remove(si)
			result = append(result, si)
		}
	}
	return result
}

func NewRegionFromSegments(segments []geometry.Segment) (Region, error) {
	edges := intersectSegments(segments)
	adj := map[geometry.Point][]geometry.Point{}
	var src geometry.Point
	for _, edge := range edges {
		adj[edge.P1] = append(adj[edge.P1], edge.P2)
		adj[edge.P2] = append(adj[edge.P2], edge.P1)
		src = edge.P1
	}
	if region, err := findCycle(adj, src); err == nil {
		return region, nil
	}
	return Region{}, errors.New("region is not closed")
}
