package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/luca-patrignani/maps/geometry"
)

type DrawModeFill struct {
	*State
}

func (g *DrawModeFill) Update() error {
	if ebiten.IsMouseButtonPressed(ebiten.MouseButton0) {
		x, y := ebiten.CursorPosition()
		p := g.Unscaled(geometry.Point{X: x, Y: y})
		g.Morph.FillWith(p, g.Fore, g.Back)
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyF) {
		Game.Wrapped = &drawModePencil
	}
	return nil
}

func (g *DrawModeFill) Draw(screen *ebiten.Image) {
	normalMode.Draw(screen)
}

func (g *DrawModeFill) Layout(outsideWidth, outsideHeight int) (int, int) {
	return outsideWidth, outsideHeight
}
