package main

import (
	"fmt"
	"image/color"
	"os"

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

func (g *NormalMode) Update() error {
	_, wheelDy := ebiten.Wheel()
	g.ViewScale += int(wheelDy)
	if g.ViewScale < 1 {
		g.ViewScale = 1
	}
	if inpututil.IsMouseButtonJustReleased(ebiten.MouseButton0) {
		x, y := ebiten.CursorPosition()
		w, h := g.Layout(ebiten.WindowSize())
		g.ViewOrigin = g.Unscaled(geometry.Point{X: x - w/2, Y: y - h/2})
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyS) {
		f, err := os.Create(g.MorphFilename)
		if err != nil {
			fmt.Println(err)
			return err
		}
		if err := g.Morph.Save(f); err != nil {
			return err
		}
		if err := f.Close(); err != nil {
			return err
		}
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyL) {
		f, err := os.Open(g.MorphFilename)
		if err != nil {
			fmt.Println(err)
			return err
		}
		*g.Morph, err = morphology.NewFromFile(f)
		if err != nil {
			return err
		}
		if err := f.Close(); err != nil {
			return err
		}
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyI) {
		Game.Wrapped = &drawModePencil
		fmt.Println("Entering draw mode")
	}
	return nil
}

func (g *NormalMode) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{0, 0, 255, 255})
	w, h := ebiten.WindowSize()
	rect := ebiten.NewImage(g.ViewScale, g.ViewScale)
	for pp, t := range g.Morph.Data {
		p := g.Scaled(pp)
		if p.X < w && p.Y < h {
			if t == morphology.Land {
				rect.Fill(color.RGBA{255, 0, 0, 255})
				geoM := ebiten.GeoM{}
				geoM.Translate(float64(p.X), float64(p.Y))
				screen.DrawImage(rect, &ebiten.DrawImageOptions{GeoM: geoM})
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

type DrawModeRubber struct {
	*State
}

func (g *DrawModeRubber) Update() error {
	if ebiten.IsMouseButtonPressed(ebiten.MouseButton0) {
		x, y := ebiten.CursorPosition()
		p := g.Unscaled(geometry.Point{X: x, Y: y})
		r := geometry.Point{}
		for r.X = p.X; r.X <= p.X+g.RubberSize; r.X++ {
			for r.Y = p.Y; r.Y <= p.Y+g.RubberSize; r.Y++ {
				g.Morph.Data[r] = g.Fore
			}
		}
	}
	_, wheelDy := ebiten.Wheel()
	g.RubberSize += int(wheelDy)
	if g.RubberSize < 1 {
		g.RubberSize = 1
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
	vector.StrokeRect(screen, float32(x), float32(y), float32(g.RubberSize*g.ViewScale), float32(g.RubberSize*g.ViewScale), float32(g.ViewScale), color.Black, true)
	ebiten.SetCursorMode(ebiten.CursorModeVisible)
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
		p := g.Unscaled(geometry.Point{X: x, Y: y})
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
