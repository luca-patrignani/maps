package main

import (
	"bytes"
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/luca-patrignani/maps/geometry"
	"github.com/luca-patrignani/maps/morphology"
)

var mplusFaceSource, _ = text.NewGoTextFaceSource(bytes.NewReader(fonts.MPlus1pRegular_ttf))

type mode interface {
	Update() error
	Draw(screen *ebiten.Image)
	Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int)
	Name() string
}

type gameWrapper struct {
	Wrapped mode
}

func (g *gameWrapper) Update() error {
	return g.Wrapped.Update()
}

func (g *gameWrapper) Draw(screen *ebiten.Image) {
	g.Wrapped.Draw(screen)
	text.Draw(screen, fmt.Sprintf("%s mode", g.Wrapped.Name()), &text.GoTextFace{
		Source: mplusFaceSource,
		Size: 20,
	}, nil)
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
	FaceSource    *text.GoTextFaceSource
}

// model -> view
func (s State) Scaled(p geometry.Point) geometry.Point {
	return geometry.Point{X: (p.X - s.ViewOrigin.X) * s.ViewScale, Y: (p.Y - s.ViewOrigin.Y) * s.ViewScale}
}

// view -> model
func (s State) Unscaled(p geometry.Point) geometry.Point {
	return geometry.Point{X: p.X/s.ViewScale + s.ViewOrigin.X, Y: p.Y/s.ViewScale + s.ViewOrigin.Y}
}

var normalMode NormalMode

var drawModePencil DrawModePencil

var drawModeRubber DrawModeRubber

var drawModeFill DrawFillMode

var Game gameWrapper
