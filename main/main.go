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
	"github.com/luca-patrignani/maps/politics"
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
	morph := &morphology.Morphology{
		Data: map[geometry.Point]morphology.MorphType{},
		MinX: 0,
		MinY: 0,
		MaxX: 2000,
		MaxY: 2000,
	}
	normalMode = NormalMode{State: &State{
		Morph: morph,
		Politics: &politics.PoliticalMap{
			Data:       make(map[geometry.Point]politics.PoliticalEntity),
			Morphology: morph,
		},
		PendingPoint:        &geometry.Point{},
		morphForeground:     morphology.Land,
		morphBackground:     morphology.Sea,
		politicalForeground: politics.None,
		politicalBackground: politics.None,
		RubberSize:          40,
		ViewScale:           15,
		ViewOrigin:          geometry.Point{X: 0, Y: 0},
		FaceSource:          faceSource,
		regionToColor:       map[politics.PoliticalEntity]color.RGBA{},
	}}
	Game = gameWrapper{Wrapped: &normalMode}
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	if err := ebiten.RunGame(&Game); err != nil {
		log.Fatal(err)
	}
}
