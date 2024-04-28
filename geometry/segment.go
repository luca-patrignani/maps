package geometry

import (
	"math"
)

type Segment struct {
	P1 Point
	P2 Point
}

// it can be nil, a Point or a Segment
type intersectionResult interface{}

func Intersection(s1, s2 Segment) intersectionResult {
	x1 := s1.P1.X
	y1 := s1.P1.Y
	x2 := s1.P2.X
	y2 := s1.P2.Y
	x3 := s2.P1.X
	y3 := s2.P1.Y
	x4 := s2.P2.X
	y4 := s2.P2.Y

	t := ((x1-x3)*(y3-y4) - (y1-y3)*(x3-x4)) / ((x1-x2)*(y3-y4) - (y1-y2)*(x3-x4))

	u := -((x1-x2)*(y1-y3) - (y1-y2)*(x1-x3)) / ((x1-x2)*(y3-y4) - (y1-y2)*(x3-x4))

	if math.IsNaN(t) || math.IsNaN(u) {
		//the segments are parallel
		P1P2 := []Point{}
		for _, p := range []Point{s1.P1, s1.P2, s2.P1, s2.P2} {
			if s1.Contains(p) && s2.Contains(p) {
				P1P2 = append(P1P2, p)
			}
		}
		if len(P1P2) >= 2 && P1P2[0] != P1P2[1] {
			return Segment{P1: P1P2[0], P2: P1P2[1]}
		}
		if len(P1P2) >= 1 {
			return P1P2[0]
		}
		return nil
	}

	if t < 0 || t > 1 || u < 0 || u > 1 {
		return nil
	}
	intersection := Point{x1 + t*(x2-x1), y1 + t*(y2-y1)}
	return intersection
}

func (segment1 Segment) IsParallelTo(segment2 Segment) bool {
	x1 := float64(segment1.P1.X)
	y1 := float64(segment1.P1.Y)
	x2 := float64(segment1.P2.X)
	y2 := float64(segment1.P2.Y)

	x3 := float64(segment2.P1.X)
	y3 := float64(segment2.P1.Y)
	x4 := float64(segment2.P2.X)
	y4 := float64(segment2.P2.Y)
	return (x1-x2)*(y3-y4)-(y1-y2)*(x3-x4) == 0
}

func (s Segment) Length() float64 {
	return s.P1.Distance(s.P2)
}

func (s Segment) Contains(p Point) bool {
	if s.P1.X == s.P2.X {
		return s.P1.X == p.X && math.Min(s.P1.Y, s.P2.Y) <= p.Y && p.Y <= math.Max(s.P1.Y, s.P2.Y)
	}
	m := (s.P1.Y - s.P2.Y) / (s.P1.X - s.P2.X)
	q := s.P1.Y - (m * s.P1.X)
	return p.X*m+q == p.Y && math.Min(s.P1.X, s.P2.X) <= p.X && p.X <= math.Max(s.P1.X, s.P2.X)
}
