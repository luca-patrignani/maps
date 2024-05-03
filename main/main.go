package main

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/luca-patrignani/maps/bresenham"
	"github.com/luca-patrignani/maps/geometry"
)

const (
	screenWidth  = 640
	screenHeight = 480
)

var (
	whiteImage = ebiten.NewImage(3, 3)
)

func init() {
	whiteImage.Fill(color.White)
}

const (
	sea  = 0
	land = 1
	lake = 2
)

type Game struct {
	morph        map[geometry.Point]int
	pendingPoint *geometry.Point
}

func (g *Game) Update() error {
	if ebiten.IsMouseButtonPressed(ebiten.MouseButton0) {
		x, y := ebiten.CursorPosition()
		p := geometry.Point{X: x, Y: y}
		if g.pendingPoint != nil {
			for _, pp := range bresenham.Bresenham(*g.pendingPoint, p) {
				g.morph[pp] = land
			}
		}
		g.pendingPoint = &p
	} else {
		g.pendingPoint = nil
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{0, 0, 255, 255})
	w, h := ebiten.WindowSize()
	for i := 0; i < w; i++ {
		for j := 0; j < h; j++ {
			if g.morph[(geometry.Point{X: i, Y: j})] == land {
				vector.StrokeRect(screen, float32(i), float32(j), 1, 1, 1, color.RGBA{255, 0, 0, 255}, true)
			}
		}
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	if err := ebiten.RunGame(&Game{morph: make(map[geometry.Point]int)}); err != nil {
		log.Fatal(err)
	}
}
