package main

import (
	"fmt"
	"image/color"
	"io"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/luca-patrignani/maps/geometry"
	"github.com/luca-patrignani/maps/morphology"
	"github.com/luca-patrignani/maps/politics"
	"github.com/luca-patrignani/maps/politics_tui"
)

type NormalMode struct {
	*State
}

var stdio politics_tui.IO = politics_tui.IO{
	In:  os.Stdin,
	Out: os.Stdout,
}

var basePath = ".saves/"

type saver interface {
	Save(w io.Writer) error
}

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
		if err := os.Mkdir(basePath, os.ModePerm); err != nil && !os.IsExist(err) {
			return err
		}
		var saverAndFilenames = map[saver]string{
			g.Morph:    basePath + "morphology.json",
			g.Politics: basePath + "politics.json",
		}
		for saver, filename := range saverAndFilenames {
			f, err := os.Create(filename)
			if err != nil {
				return err
			}
			if err := saver.Save(f); err != nil {
				return err
			}
			if err := f.Close(); err != nil {
				return err
			}
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
		Game.Wrapped = &DrawModePencil[morphology.MorphType]{
			State:      g.State,
			geography:  g.Morph,
			foreground: &g.morphForeground,
			background: &g.morphBackground,
			label:      "Morphology",
		}
		fmt.Println("Entering draw mode")
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyP) {
		Game.Wrapped = &DrawModePencil[politics.PoliticalEntity]{
			State:      g.State,
			geography:  g.Politics,
			foreground: &g.politicalForeground,
			background: &g.politicalBackground,
			label:      "Politics",
		}
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyN) {
		name, c, err := stdio.NewPoliticalEntity()
		if err != nil {
			fmt.Println(err)
		} else {
			g.politicalForeground = name
			g.regionToColor[name] = c
		}
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

	for pp, t := range g.Politics.Data {
		p := g.Scaled(pp)
		if p.X < w && p.Y < h {
			if t != politics.None {
				c, ok := g.regionToColor[t]
				if !ok {
					c = color.RGBA{0, 0, 0, 255}
				}
				rect.Fill(c)
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

func (g *NormalMode) Name() string {
	return "Normal"
}
