package main

import (
	"bytes"
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/luca-patrignani/maps/geometry"
	"github.com/luca-patrignani/maps/morphology"
	"github.com/luca-patrignani/maps/politics"
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
	text.Draw(
		screen,
		fmt.Sprintf("%s mode", g.Wrapped.Name()),
		&text.GoTextFace{
			Source: mplusFaceSource,
			Size:   20,
		},
		nil)
}

func (g *gameWrapper) Layout(outsideWidth, outsideHeight int) (int, int) {
	return g.Wrapped.Layout(outsideWidth, outsideHeight)
}

type State struct {
	morph                                    *morphology.Morphology
	politics                                 *politics.PoliticalMap
	pendingPoint                             *geometry.Point
	morphForeground, morphBackground         morphology.MorphType
	politicalForeground, politicalBackground politics.PoliticalEntity
	rubberSize                               int
	viewScale                                int
	viewOrigin                               geometry.Point
	faceSource                               *text.GoTextFaceSource
	politicalEntities                        []politics.PoliticalEntity
}

// model -> view
func (s State) Scaled(p geometry.Point) geometry.Point {
	return geometry.Point{X: (p.X - s.viewOrigin.X) * s.viewScale, Y: (p.Y - s.viewOrigin.Y) * s.viewScale}
}

// view -> model
func (s State) Unscaled(p geometry.Point) geometry.Point {
	return geometry.Point{X: p.X/s.viewScale + s.viewOrigin.X, Y: p.Y/s.viewScale + s.viewOrigin.Y}
}

var Game gameWrapper

var normalMode NormalMode
