package politics

import (
	"encoding/json"
	"io"

	"github.com/luca-patrignani/maps/bresenham"
	"github.com/luca-patrignani/maps/geometry"
	"github.com/luca-patrignani/maps/morphology"
)

type PoliticalEntity string

const None PoliticalEntity = ""

type PoliticalMap struct {
	Morphology *morphology.Morphology
	Data       map[geometry.Point]PoliticalEntity
}

func (pm PoliticalMap) FillWith(p geometry.Point, foreground, background PoliticalEntity) {
	stack := []geometry.Point{p}
	for len(stack) > 0 {
		u := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if pm.Morphology.Data[u] == morphology.Land && pm.Data[u] == background {
			pm.Data[u] = foreground
			stack = append(stack, geometry.Point{X: u.X + 1, Y: u.Y}, geometry.Point{X: u.X, Y: u.Y + 1}, geometry.Point{X: u.X - 1, Y: u.Y}, geometry.Point{X: u.X, Y: u.Y - 1})
		}
	}
}

func (pm PoliticalMap) DrawLine(p1, p2 geometry.Point, t PoliticalEntity) {
	for _, p := range bresenham.Bresenham(p1, p2) {
		if pm.Morphology.Data[p] == morphology.Land {
			pm.Data[p] = t
		}
	}
}

func (m PoliticalMap) Save(w io.Writer) error {
	j, err := json.MarshalIndent(m.Data, "", "  ")
	if err != nil {
		return err
	}
	_, err = w.Write(j)
	return err
}

func (p *PoliticalMap) Load(r io.Reader) error {
	j, err := io.ReadAll(r)
	if err != nil {
		return err
	}
	err = json.Unmarshal(j, &p.Data)
	return err
}
