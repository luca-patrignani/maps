//go:build exclude

package main

import (
	"fmt"
	"io/fs"
	"os"

	"github.com/luca-patrignani/maps/geometry"
	"github.com/luca-patrignani/maps/island"
	"github.com/luca-patrignani/maps/regions"
	"github.com/spf13/afero"
	"github.com/veandco/go-sdl2/gfx"
	"github.com/veandco/go-sdl2/sdl"
)

func main() {
	window, err := sdl.CreateWindow("test", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, 800, 600, sdl.WINDOW_RESIZABLE)

	if err != nil {
		panic(err)
	}
	defer window.Destroy()

	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		panic(err)
	}
	defer renderer.Destroy()

	renderer.SetDrawColor(0, 255, 0, 255)
	renderer.FillRect(&sdl.Rect{
		X: 10,
		Y: 10,
		W: 100,
		H: 100,
	})
	var scale int32 = 10
	renderer.SetScale(float32(scale), float32(scale))
	rb := regions.RegionBuilder{}
	if err := os.MkdirAll("./maps", fs.ModePerm); err != nil {
		panic(err)
	}
	var ir island.IslandRepository
	ir, err = island.NewIslandRepository(afero.NewBasePathFs(afero.OsFs{}, "./maps"), "test.json")
	if err != nil {
		fmt.Println(err)
		fmt.Println("creating a new repository from scratch")
		ir, err = island.InitIslandRepository(afero.NewBasePathFs(afero.OsFs{}, "./maps"), "test.json")
		if err != nil {
			panic(err)
		}
	}

	running := true
	pressed := false
	for running {
		renderer.SetDrawColor(0, 0, 0, 255)
		renderer.Clear()

		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch t := event.(type) {
			case *sdl.QuitEvent:
				running = false
			case *sdl.MouseMotionEvent:
				if pressed {
					newPoint := geometry.Point{X: t.X / scale, Y: t.Y / scale}
					rb.AddPoint(newPoint)
				}
			case *sdl.MouseButtonEvent:
				if t.State == sdl.PRESSED {
					pressed = true
					fmt.Println("Mouse", t.Which, "button", t.Button, "pressed at", t.X, t.Y)
				} else {
					pressed = false
					fmt.Println("Mouse", t.Which, "button", t.Button, "released at", t.X, t.Y)
					if region, err := rb.Build(); err == nil {
						ir.Save(island.Island{Name: "1", Region: region})
					}
					rb = regions.RegionBuilder{}
				}
			}

		}
		for _, island := range ir.Islands() {
			drawIsland(renderer, island, sdl.Color{R: 255, G: 0, B: 0, A: 255})
		}
		renderer.SetDrawColor(0, 255, 0, 255)
		for _, segment := range rb.Segments {
			renderer.DrawLine(segment.P1.X, segment.P1.Y, segment.P2.X, segment.P2.Y)
		}
		renderer.Present()
	}
}

func drawIsland(renderer *sdl.Renderer, island island.Island, color sdl.Color) {
	vx := []int16{}
	vy := []int16{}
	for _, point := range island.Region {
		vx = append(vx, int16(point.X))
		vy = append(vy, int16(point.Y))
	}
	gfx.FilledPolygonColor(
		renderer,
		vx,
		vy,
		color,
	)
}
