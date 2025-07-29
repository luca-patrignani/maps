package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/luca-patrignani/maps/geometry"
)

type DrawModePencil[T any] struct {
	*State
	geography geography[T]
}

func (g *DrawModePencil[T]) Update() error {
	x, y := ebiten.CursorPosition()
	p := g.Unscaled(geometry.Point{X: x, Y: y})
	if ebiten.IsMouseButtonPressed(ebiten.MouseButton0) {
		if g.PendingPoint != nil {
			g.Morph.DrawLine(p, *g.PendingPoint, g.Fore)
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
		g.State.Fore, g.State.Back = g.State.Back, g.State.Fore
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyF) {
		Game.Wrapped = &DrawFillMode{
			State: g.State,
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
	return "Draw pencil"
}
