package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/luca-patrignani/maps/geometry"
)

type DrawModePencil[T any] struct {
	*State
	geography              geography[T]
	foreground, background *T
	label                  string
}

func (g *DrawModePencil[T]) Update() error {
	x, y := ebiten.CursorPosition()
	p := g.Unscaled(geometry.Point{X: x, Y: y})
	if ebiten.IsMouseButtonPressed(ebiten.MouseButton0) {
		if g.PendingPoint != nil {
			g.geography.DrawLine(p, *g.PendingPoint, *g.foreground)
		}
		g.PendingPoint = &p
	} else {
		g.PendingPoint = nil
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyR) {
		Game.Wrapped = &DrawModeRubber{
			State: g.State,
		}
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		Game.Wrapped = &normalMode
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyS) {
		g.foreground, g.background = g.background, g.foreground
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyF) {
		Game.Wrapped = &DrawFillMode[T]{
			State:      g.State,
			geography:  g.geography,
			foreground: g.foreground,
			background: g.background,
			label:      g.label,
		}
	}
	return nil
}

func (g *DrawModePencil[T]) Draw(screen *ebiten.Image) {
	normalMode.Draw(screen)
}

func (g *DrawModePencil[T]) Layout(outsideWidth, outsideHeight int) (int, int) {
	return outsideWidth, outsideHeight
}

func (g *DrawModePencil[T]) Name() string {
	return g.label + ": Draw pencil"
}
