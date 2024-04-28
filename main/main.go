package main

import (
	"fmt"
	"image"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/luca-patrignani/maps/geometry"
	"github.com/luca-patrignani/maps/island"
	"github.com/luca-patrignani/maps/regions"
	"github.com/spf13/afero"
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

type Game struct {
	ir         island.IslandRepository
	nations    []regions.Region
	drawNation bool
	rb         regions.RegionBuilder
	scale      float32
}

func (g *Game) Update() error {
	if inpututil.IsKeyJustPressed(ebiten.KeyN) {
		g.drawNation = !g.drawNation
		if g.drawNation {
			fmt.Println("Draw nation mode")
		} else {
			fmt.Println("Normal mode")
		}
	}
	if ebiten.IsMouseButtonPressed(ebiten.MouseButton0) {
		x, y := ebiten.CursorPosition()
		g.rb.AddPoint(geometry.Point{
			X: float64(x / int(g.scale)),
			Y: float64(y / int(g.scale)),
		})
	} else {
		if region, err := g.rb.Build(); err == nil {
			if g.drawNation {
				for _, island := range g.ir.Islands() {
					if nation, err := region.Intersection(island.Region); err == nil {
						g.nations = append(g.nations, nation)
						break
					}
				}
				fmt.Println(g.nations)
			} else {
				intersect := false
				for _, island := range g.ir.Islands() {
					if _, err := island.Region.Intersection(region); err == nil {
						intersect = true
						break
					}
				}
				if !intersect {
					if _, err := g.ir.Save(island.Island{Name: "1", Region: region}); err != nil {
						fmt.Println(err)
					}
				}
			}
		}
		g.rb = regions.RegionBuilder{}
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	for _, island := range g.ir.Islands() {
		g.drawRegion(screen, island.Region, color.RGBA64{R: 255, G: 0, B: 0, A: 255})
	}
	for _, s := range g.rb.Segments {
		vector.StrokeLine(screen,
			float32(s.P1.X)*g.scale, float32(s.P1.Y)*g.scale, float32(s.P2.X)*g.scale, float32(s.P2.Y)*g.scale,
			g.scale, color.White, false)
	}
}

func (game *Game) drawRegion(screen *ebiten.Image, region regions.Region, color color.Color) {
	op := &ebiten.DrawTrianglesOptions{}
	op.Address = ebiten.AddressUnsafe
	indices := []uint16{}
	for i := 0; i < len(region); i++ {
		indices = append(indices, uint16(i), uint16(i+1)%uint16(len(region)), uint16(len(region)-1))
	}
	r, g, b, a := color.RGBA()
	vertices := []ebiten.Vertex{}
	for _, point := range region {
		vertices = append(vertices, ebiten.Vertex{
			DstX:   float32(point.X) * game.scale,
			DstY:   float32(point.Y) * game.scale,
			ColorR: float32(r),
			ColorG: float32(g),
			ColorB: float32(b),
			ColorA: float32(a),
		})
	}
	screen.DrawTriangles(vertices, indices, whiteImage.SubImage(image.Rect(1, 1, 2, 2)).(*ebiten.Image), op)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return outsideWidth, outsideHeight
}

func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	ebiten.SetWindowTitle("maps")
	ir, err := island.NewIslandRepository(afero.NewBasePathFs(afero.OsFs{}, "./maps"), "test.json")
	if err != nil {
		fmt.Println(err)
		fmt.Println("creating a new repository from scratch")
		ir, err = island.InitIslandRepository(afero.NewBasePathFs(afero.OsFs{}, "./maps"), "test.json")
		if err != nil {
			panic(err)
		}
	}
	if err := ebiten.RunGame(&Game{ir: ir, scale: 10}); err != nil {
		log.Fatal(err)
	}
}
