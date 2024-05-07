package morphology

import (
	"encoding/json"
	"io"

	"github.com/luca-patrignani/maps/bresenham"
	"github.com/luca-patrignani/maps/geometry"
)

type MorphType int

const (
	Sea  MorphType = 0
	Land MorphType = 1
)

type Morphology struct {
	Data                   map[geometry.Point]MorphType
	MinX, MinY, MaxX, MaxY int
}

func New(minX, minY, maxX, maxY int) Morphology {
	return Morphology{
		Data: map[geometry.Point]MorphType{},
		MinX: minX,
		MinY: minY,
		MaxX: maxX,
		MaxY: maxY,
	}
}

func (m Morphology) FillWith(p geometry.Point, foreground, background MorphType) bool {
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

func (m Morphology) DrawLine(p1, p2 geometry.Point, t MorphType) {
	for _, pp := range bresenham.Bresenham(p1, p2) {
		m.Data[pp] = t
	}
}

func (m Morphology) Save(w io.Writer) error {
	j, err := json.MarshalIndent(m, "", "  ")
	if err != nil {
		return err
	}
	_, err = w.Write(j)
	return err
}

func NewFromFile(r io.Reader) (Morphology, error) {
	j, err := io.ReadAll(r)
	if err != nil {
		return Morphology{}, err
	}
	m := Morphology{}
	err = json.Unmarshal(j, &m)
	return m, err
}
