package main

import (
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/luca-patrignani/maps/geometry"
	"github.com/luca-patrignani/maps/morphology"
)

type gameWrapper struct {
	Wrapped ebiten.Game
}

func (g *gameWrapper) Update() error {
	return g.Wrapped.Update()
}

func (g *gameWrapper) Draw(screen *ebiten.Image) {
	g.Wrapped.Draw(screen)
}

func (g *gameWrapper) Layout(outsideWidth, outsideHeight int) (int, int) {
	return g.Wrapped.Layout(outsideWidth, outsideHeight)
}

type State struct {
	Morph            *morphology.Morphology
	PendingPoint     *geometry.Point
	Fore, Back       morphology.MorphType
	RubberW, RubberH int
}

type NormalMode struct {
	*State
}

var normalMode NormalMode = NormalMode{State: &State{
	Morph: &morphology.Morphology{
		Data: map[geometry.Point]morphology.MorphType{},
		MinX: 0,
		MinY: 0,
		MaxX: 2000,
		MaxY: 2000,
	},
	PendingPoint: &geometry.Point{},
	Fore:         morphology.Land,
	Back:         morphology.Sea,
	RubberW:      100,
	RubberH:      100,
}}

var drawModePencil DrawModePencil = DrawModePencil(normalMode)

var drawModeRubber DrawModeRubber = DrawModeRubber(normalMode)

var drawModeFill DrawModeFill = DrawModeFill(normalMode)

var Game gameWrapper = gameWrapper{Wrapped: &normalMode}

func (g *NormalMode) Update() error {
	if inpututil.IsKeyJustPressed(ebiten.KeyI) {
		Game.Wrapped = &drawModePencil
		fmt.Println("Entering draw mode")
	}
	return nil
}

func (g *NormalMode) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{0, 0, 255, 255})
	w, h := ebiten.WindowSize()
	for p, t := range g.Morph.Data {
		if p.X < w && p.Y < h {
			if t == morphology.Land {
				vector.StrokeRect(screen, float32(p.X), float32(p.Y), 1, 1, 1, color.RGBA{255, 0, 0, 255}, true)
			}
		}
	}
}

func (g *NormalMode) Layout(outsideWidth, outsideHeight int) (int, int) {
	return outsideWidth, outsideHeight
}

type DrawModePencil struct {
	*State
}

func (g *DrawModePencil) Update() error {
	x, y := ebiten.CursorPosition()
	p := geometry.Point{X: x, Y: y}
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

type DrawModeRubber struct {
	*State
}

func (g *DrawModeRubber) Update() error {
	x, y := ebiten.CursorPosition()
	if ebiten.IsMouseButtonPressed(ebiten.MouseButton0) {
		r := geometry.Point{}
		for r.X = x; r.X <= x+g.RubberW; r.X++ {
			for r.Y = y; r.Y <= y+g.RubberH; r.Y++ {
				g.Morph.Data[r] = g.Fore
			}
		}
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
	vector.StrokeRect(screen, float32(x), float32(y), float32(g.RubberW), float32(g.RubberH), 1, color.Black, true)
}

func (g *DrawModeRubber) Layout(outsideWidth, outsideHeight int) (int, int) {
	return outsideWidth, outsideHeight
}

type DrawModeFill struct {
	*State
}

func (g *DrawModeFill) Update() error {
	if ebiten.IsMouseButtonPressed(ebiten.MouseButton0) {
		x, y := ebiten.CursorPosition()
		p := geometry.Point{X: x, Y: y}
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
