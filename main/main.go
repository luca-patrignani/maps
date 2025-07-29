package main

import (
	"bytes"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/luca-patrignani/maps/geometry"
	"github.com/luca-patrignani/maps/morphology"
)

var (
	whiteImage = ebiten.NewImage(3, 3)
)

func init() {
	whiteImage.Fill(color.White)
}

func main() {
	faceSource, err := text.NewGoTextFaceSource(bytes.NewReader(fonts.MPlus1pRegular_ttf))
	if err != nil {
		log.Fatal(err)
	}
	normalMode = NormalMode{State: &State{
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
		ViewScale:     15,
		ViewOrigin:    geometry.Point{X: 0, Y: 0},
		MorphFilename: "morphology.json",
		FaceSource: faceSource,
	}}
	drawModePencil = DrawModePencil(normalMode)
	drawModeRubber = DrawModeRubber(normalMode)
	drawModeFill  = DrawFillMode(normalMode)
	Game = gameWrapper{Wrapped: &normalMode}
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	if err := ebiten.RunGame(&Game); err != nil {
		log.Fatal(err)
	}
}
