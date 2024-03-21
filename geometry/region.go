package geometry

import "errors"

type Region []Point

func NewRegion(points []Point) (Region, error) {
	edges := []Segment{}
	for i := 1; i < len(points); i++ {
		newEdge := Segment{points[i-1], points[i]}
		for j := 0; j < len(edges)-1; j++ {
			inter, err := Intersection(newEdge, edges[j])
			if err == nil {
				edges = append(edges, newEdge)
				region := Region{inter}
				for k := j; edges[k].P2 != newEdge.P2; k++ {
					region = append(region, edges[k].P2)
				}
				// region = append(region, edges[j].P2)
				return region, nil
			}
		}
		edges = append(edges, newEdge)
	}
	return Region{}, errors.New("region is not closed")
}
