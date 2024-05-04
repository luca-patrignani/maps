package main

import (
	"fmt"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/luca-patrignani/maps/geometry"
	"github.com/luca-patrignani/maps/morphology"
)

var (
	whiteImage = ebiten.NewImage(3, 3)
)

func init() {
	whiteImage.Fill(color.White)
}

type Game struct {
	morph        morphology.Morphology
	pendingPoint *geometry.Point
	fore, back morphology.MorphType
}

func (g *Game) Update() error {
	x, y := ebiten.CursorPosition()
	p := geometry.Point{X: x, Y: y}
	if ebiten.IsMouseButtonPressed(ebiten.MouseButton0) {
		if g.pendingPoint != nil {
			g.morph.DrawLine(p, *g.pendingPoint, g.fore)
		}
		g.pendingPoint = &p
	} else {
		g.pendingPoint = nil
	}
	if ebiten.IsMouseButtonPressed(ebiten.MouseButton2) {
		if g.morph.FillWith(p, g.fore, g.back) {
			fmt.Println("fill")
		}
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyI) {
		g.fore, g.back = g.back, g.fore
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{0, 0, 255, 255})
	w, h := ebiten.WindowSize()
	for p, t := range g.morph.Data {
		if p.X < w && p.Y < h {
			if t == morphology.Land {
				vector.StrokeRect(screen, float32(p.X), float32(p.Y), 1, 1, 1, color.RGBA{255, 0, 0, 255}, true)
			}
		}
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return outsideWidth, outsideHeight
}

func main() {
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	if err := ebiten.RunGame(&Game{morph: morphology.New(0, 0, 2000, 2000), fore: morphology.Land, back: morphology.Sea}); err != nil {
		log.Fatal(err)
	}
}
