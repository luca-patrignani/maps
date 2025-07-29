package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/luca-patrignani/maps/geometry"
)

type DrawModePencil struct {
	*State
}

func (g *DrawModePencil) Update() error {
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
		Game.Wrapped = &drawModeRubber
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		Game.Wrapped = &normalMode
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyS) {
		g.State.Fore, g.State.Back = g.State.Back, g.State.Fore
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyF) {
		Game.Wrapped = &drawModeFill
	}
	return nil
}

func (g *DrawModePencil) Draw(screen *ebiten.Image) {
	normalMode.Draw(screen)
}

func (g *DrawModePencil) Layout(outsideWidth, outsideHeight int) (int, int) {
	return outsideWidth, outsideHeight
}

func (g *DrawModePencil) Name() string {
	return "Draw pencil"
}
