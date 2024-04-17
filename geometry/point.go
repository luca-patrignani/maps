package geometry

import "math"

type Point struct {
	X float64
	Y float64
}

func (p Point) Distance(o Point) float64 {
	return math.Sqrt(math.Pow(p.X-o.X, 2) + math.Pow(p.Y-o.Y, 2))
}
