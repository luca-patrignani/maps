package main

import (
	"github.com/hajimehoshi/ebiten/v2"
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
	Morph         *morphology.Morphology
	MorphFilename string
	PendingPoint  *geometry.Point
	Fore, Back    morphology.MorphType
	RubberSize    int
	ViewScale     int
	ViewOrigin    geometry.Point
}

// model -> view
func (s State) Scaled(p geometry.Point) geometry.Point {
	return geometry.Point{X: (p.X - s.ViewOrigin.X) * s.ViewScale, Y: (p.Y - s.ViewOrigin.Y) * s.ViewScale}
}

// view -> model
func (s State) Unscaled(p geometry.Point) geometry.Point {
	return geometry.Point{X: p.X/s.ViewScale + s.ViewOrigin.X, Y: p.Y/s.ViewScale + s.ViewOrigin.Y}
}

var normalMode NormalMode = NormalMode{State: &State{
	Morph: &morphology.Morphology{
		Data: map[geometry.Point]morphology.MorphType{},
		MinX: 0,
		MinY: 0,
		MaxX: 2000,
		MaxY: 2000,
	},
	PendingPoint:  &geometry.Point{},
	Fore:          morphology.Land,
	Back:          morphology.Sea,
	RubberSize:    40,
	ViewScale:     1,
	ViewOrigin:    geometry.Point{X: 0, Y: 0},
	MorphFilename: "morphology.json",
}}

var drawModePencil DrawModePencil = DrawModePencil(normalMode)

var drawModeRubber DrawModeRubber = DrawModeRubber(normalMode)

var drawModeFill DrawModeFill = DrawModeFill(normalMode)

var Game gameWrapper = gameWrapper{Wrapped: &normalMode}
