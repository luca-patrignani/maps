package bresenham

import "github.com/luca-patrignani/maps/geometry"

// https://en.wikipedia.org/wiki/Bresenham%27s_line_algorithm#All_cases

func bresenhamLow(p0, p1 geometry.Point) []geometry.Point {
	dx := p1.X - p0.X
	dy := p1.Y - p0.Y
	yi := 1
	if dy < 0 {
		yi = -1
		dy = -dy
	}
	D := (2 * dy) - dx
	y := p0.Y
	result := []geometry.Point{}
	for x := p0.X; x <= p1.X; x++ {
		result = append(result, geometry.Point{X: x, Y: y})
		if D > 0 {
			y = y + yi
			D = D + (2 * (dy - dx))
		} else {
			D = D + 2*dy
		}
	}
	return result
}

func bresenhamHigh(p0, p1 geometry.Point) []geometry.Point {
	dx := p1.X - p0.X
	dy := p1.Y - p0.Y
	xi := 1
	if dx < 0 {
		xi = -1
		dx = -dx
	}
	D := (2 * dx) - dy
	x := p0.X
	result := []geometry.Point{}
	for y := p0.Y; y <= p1.Y; y++ {
		result = append(result, geometry.Point{X: x, Y: y})
		if D > 0 {
			x = x + xi
			D = D + (2 * (dx - dy))
		} else {
			D = D + 2*dx
		}
	}
	return result
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func Bresenham(p0 geometry.Point, p1 geometry.Point) []geometry.Point {
	if abs(p1.Y-p0.Y) < abs(p1.X-p0.X) {
		if p0.X > p1.X {
			return bresenhamLow(p1, p0)
		}
		return bresenhamLow(p0, p1)
	}
	if p0.Y > p1.Y {
		return bresenhamHigh(p1, p0)
	}
	return bresenhamHigh(p0, p1)
}
