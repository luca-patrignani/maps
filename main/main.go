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
	var scale int32 = 20
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
	nations := []regions.Region{}

	running := true
	pressed := false
	drawNation := false
	for running {
		renderer.SetDrawColor(0, 0, 0, 255)
		renderer.Clear()

		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch t := event.(type) {
			case *sdl.QuitEvent:
				running = false
			case *sdl.MouseMotionEvent:
				if pressed {
					newPoint := geometry.Point{
						X: float64(t.X / int32(scale)),
						Y: float64(t.Y / int32(scale)),
					}
					rb.AddPoint(newPoint)
				}
			case *sdl.MouseButtonEvent:
				if t.State == sdl.PRESSED {
					pressed = true
				} else {
					pressed = false
					if region, err := rb.Build(); err == nil {
						if drawNation {
							for _, island := range ir.Islands() {
								if nation, err := region.Intersection(island.Region); err == nil {
									nations = append(nations, nation)
									break
								}
							}
							fmt.Println(nations)
						} else {
							intersect := false
							for _, island := range ir.Islands() {
								if _, err := island.Region.Intersection(region); err == nil {
									intersect = true
									break
								}
							}
							if !intersect {
								ir.Save(island.Island{Name: "1", Region: region})
							}
						}
					}
					rb = regions.RegionBuilder{}
				}
			case *sdl.KeyboardEvent:
				if t.Keysym.Sym == sdl.K_n && t.Type == sdl.KEYUP {
					drawNation = !drawNation
					if drawNation {
						fmt.Println("Draw nation mode")
					} else {
						fmt.Println("Normal mode")
					}
				}
			}
		}
		for _, island := range ir.Islands() {
			drawRegion(renderer, island.Region, sdl.Color{R: 255, G: 0, B: 0, A: 255})
		}
		for _, nation := range nations {
			drawRegion(renderer, nation, sdl.Color{R: 0, G: 0, B: 255, A: 255})
		}
		renderer.SetDrawColor(0, 255, 0, 255)
		for _, segment := range rb.Segments {
			renderer.DrawLine(int32(segment.P1.X), int32(segment.P1.Y), int32(segment.P2.X), int32(segment.P2.Y))
		}
		renderer.Present()
	}
}

func drawRegion(renderer *sdl.Renderer, region regions.Region, color sdl.Color) {
	vx := []int16{}
	vy := []int16{}
	for _, point := range region {
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
