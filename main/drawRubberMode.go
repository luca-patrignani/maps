package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/luca-patrignani/maps/geometry"
	"github.com/luca-patrignani/maps/morphology"
)

type DrawModeRubber struct {
	*State
}

func (g *DrawModeRubber) Update() error {
	if ebiten.IsMouseButtonPressed(ebiten.MouseButton0) {
		x, y := ebiten.CursorPosition()
		p := g.Unscaled(geometry.Point{X: x, Y: y})
		g.morph.DrawSquare(p, g.rubberSize, g.morphForeground)
	}
	_, wheelDy := ebiten.Wheel()
	g.rubberSize += int(wheelDy)
	if g.rubberSize < 1 {
		g.rubberSize = 1
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyR) {
		Game.Wrapped = &DrawModePencil[morphology.MorphType]{
			State:     g.State,
			geography: g.morph,
			label:     "morphology",
		}
	}
	return nil
}

func (g *DrawModeRubber) Draw(screen *ebiten.Image) {
	normalMode.Draw(screen)
	ebiten.SetCursorMode(ebiten.CursorModeHidden)
	x, y := ebiten.CursorPosition()
	vector.StrokeRect(screen, float32(x), float32(y), float32(g.rubberSize*g.viewScale), float32(g.rubberSize*g.viewScale), float32(g.viewScale), color.Black, true)
	ebiten.SetCursorMode(ebiten.CursorModeVisible)
}

func (g *DrawModeRubber) Layout(outsideWidth, outsideHeight int) (int, int) {
	return outsideWidth, outsideHeight
}

func (g *DrawModeRubber) Name() string {
	return "Draw rubber"
}
