package regions

import "github.com/luca-patrignani/maps/geometry"

type RegionBuilder struct {
	segments []geometry.Segment
}

func NewRegionBuilder() RegionBuilder {
	return RegionBuilder{}
}