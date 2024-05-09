package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/luca-patrignani/maps/geometry"
)

type DrawModeRubber struct {
	*State
}

func (g *DrawModeRubber) Update() error {
	if ebiten.IsMouseButtonPressed(ebiten.MouseButton0) {
		x, y := ebiten.CursorPosition()
		p := g.Unscaled(geometry.Point{X: x, Y: y})
		r := geometry.Point{}
		for r.X = p.X; r.X <= p.X+g.RubberSize; r.X++ {
			for r.Y = p.Y; r.Y <= p.Y+g.RubberSize; r.Y++ {
				g.Morph.Data[r] = g.Fore
			}
		}
	}
	_, wheelDy := ebiten.Wheel()
	g.RubberSize += int(wheelDy)
	if g.RubberSize < 1 {
		g.RubberSize = 1
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyR) {
		Game.Wrapped = &drawModePencil
	}
	return nil
}

func (g *DrawModeRubber) Draw(screen *ebiten.Image) {
	normalMode.Draw(screen)
	ebiten.SetCursorMode(ebiten.CursorModeHidden)
	x, y := ebiten.CursorPosition()
	vector.StrokeRect(screen, float32(x), float32(y), float32(g.RubberSize*g.ViewScale), float32(g.RubberSize*g.ViewScale), float32(g.ViewScale), color.Black, true)
	ebiten.SetCursorMode(ebiten.CursorModeVisible)
}

func (g *DrawModeRubber) Layout(outsideWidth, outsideHeight int) (int, int) {
	return outsideWidth, outsideHeight
}