package main

import (
	"fmt"
	"image/color"
	"io"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
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
	Load(r io.Reader) error
}

func (g *NormalMode) Update() error {
	_, wheelDy := ebiten.Wheel()
	g.viewScale += int(wheelDy)
	if g.viewScale < 1 {
		g.viewScale = 1
	}
	if inpututil.IsMouseButtonJustReleased(ebiten.MouseButton0) {
		x, y := ebiten.CursorPosition()
		w, h := g.Layout(ebiten.WindowSize())
		g.viewOrigin = g.Unscaled(geometry.Point{X: x - w/2, Y: y - h/2})
	}
	var saverAndFilenames = map[saver]string{
		g.morph:    basePath + "morphology.json",
		g.politics: basePath + "politics.json",
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyS) {
		if err := os.Mkdir(basePath, os.ModePerm); err != nil && !os.IsExist(err) {
			return err
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
		for saver, filename := range saverAndFilenames {
			f, err := os.Open(filename)
			if err != nil {
				fmt.Println(err)
				return err
			}
			err = saver.Load(f)
			if err != nil {
				return err
			}
			if err := f.Close(); err != nil {
				return err
			}
		}
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyI) {
		Game.Wrapped = &DrawModePencil[morphology.MorphType]{
			State:      g.State,
			geography:  g.morph,
			foreground: &g.morphForeground,
			background: &g.morphBackground,
			label:      "Morphology",
		}
		fmt.Println("Entering draw mode")
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyP) {
		Game.Wrapped = &DrawModePencil[politics.PoliticalEntity]{
			State:      g.State,
			geography:  g.politics,
			foreground: &g.politicalForeground,
			background: &g.politicalBackground,
			label:      "Politics",
		}
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyN) {
		entity, err := stdio.NewPoliticalEntity()
		if err != nil {
			fmt.Println(err)
		} else {
			g.politicalForeground = entity
			g.politicalEntities = append(g.politicalEntities, entity)
		}
	}
	return nil
}

func (g *NormalMode) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{0, 0, 255, 255})
	w, h := ebiten.WindowSize()
	rect := ebiten.NewImage(g.viewScale, g.viewScale)
	for pp, t := range g.morph.Data {
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

	for pp, entity := range g.politics.Data {
		p := g.Scaled(pp)
		if p.X < w && p.Y < h {
			if entity != politics.None {
				rect.Fill(entity.Color)
				geoM := ebiten.GeoM{}
				geoM.Translate(float64(p.X), float64(p.Y))
				screen.DrawImage(rect, &ebiten.DrawImageOptions{GeoM: geoM})
			}
		}
	}

	x, y := ebiten.CursorPosition()
	cursor := g.Unscaled(geometry.Point{X: x, Y: y})
	if entity, ok := g.politics.Data[cursor]; ok {
		geoM := ebiten.GeoM{}
		geoM.Translate(float64(x), float64(y+20))
		text.Draw(
			screen,
			entity.Name,
			&text.GoTextFace{
				Source: mplusFaceSource,
				Size:   15,
			},
			&text.DrawOptions{DrawImageOptions: ebiten.DrawImageOptions{GeoM: geoM}},
		)
	}
}

func (g *NormalMode) Layout(outsideWidth, outsideHeight int) (int, int) {
	return outsideWidth, outsideHeight
}

func (g *NormalMode) Name() string {
	return "Normal"
}
