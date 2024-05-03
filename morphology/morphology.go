package morphology

import (
	"github.com/luca-patrignani/maps/bresenham"
	"github.com/luca-patrignani/maps/geometry"
)

type morphType int

const (
	Sea  morphType = 0
	Land morphType = 1
)

type Morphology struct {
	Data                   map[geometry.Point]morphType
	MinX, MinY, MaxX, MaxY int
}

func New(minX, minY, maxX, maxY int) Morphology {
	return Morphology{
		Data: map[geometry.Point]morphType{},
		MinX: minX,
		MinY: minY,
		MaxX: maxX,
		MaxY: maxY,
	}
}

func (m Morphology) FillWith(p geometry.Point, foreground, background morphType) bool {
	if m.Data[p] != background {
		return false
	}
	stack := []geometry.Point{p}
	for len(stack) > 0 {
		u := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		for _, v := range []geometry.Point{{X: u.X + 1, Y: u.Y}, {X: u.X, Y: u.Y + 1}, {X: u.X - 1, Y: u.Y}, {X: u.X, Y: u.Y - 1}} {
			if v.X >= m.MinX && v.X < m.MaxX && v.Y >= m.MinY && v.Y < m.MaxY && m.Data[v] == background {
				m.Data[v] = foreground
				stack = append(stack, v)
			}
		}
	}
	return true
}

func (m Morphology) DrawLine(p1, p2 geometry.Point, t morphType) {
	for _, pp := range bresenham.Bresenham(p1, p2) {
		m.Data[pp] = t
	}
}
