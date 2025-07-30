package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/luca-patrignani/maps/geometry"
)

type DrawFillMode[T any] struct {
	*State
	geography              geography[T]
	foreground, background *T
	label                  string
}

func (g *DrawFillMode[T]) Update() error {
	if ebiten.IsMouseButtonPressed(ebiten.MouseButton0) {
		x, y := ebiten.CursorPosition()
		p := g.Unscaled(geometry.Point{X: x, Y: y})
		g.geography.FillWith(p, *g.foreground, *g.background)
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyF) {
		Game.Wrapped = &DrawModePencil[T]{
			g.State,
			g.geography,
			g.foreground,
			g.background,
			g.label,
		}
	}
	return nil
}

func (g *DrawFillMode[T]) Draw(screen *ebiten.Image) {
	normalMode.Draw(screen)
}

func (g *DrawFillMode[T]) Layout(outsideWidth, outsideHeight int) (int, int) {
	return outsideWidth, outsideHeight
}

func (g DrawFillMode[T]) Name() string {
	return g.label + ": Draw fill"
}
