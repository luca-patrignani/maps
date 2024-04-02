//go:build exclude
package main

import (
	"fmt"
	"io/fs"
	"os"

	"github.com/luca-patrignani/maps"
	"github.com/luca-patrignani/maps/geometry"
	"github.com/luca-patrignani/maps/regions"
	"github.com/spf13/afero"
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
	rr := maps.RegionRepository{Fs: afero.NewBasePathFs(afero.OsFs{}, "./maps"), Filename: "test.json"}
	rs, err := rr.Load()
	if err != nil {
		fmt.Println(err)
		fmt.Println("Creating new region file")
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
						rs = append(rs, region)
					}
					rb = regions.RegionBuilder{}
				}
			case *sdl.KeyboardEvent:
				switch t.Keysym.Sym {
				case sdl.K_s:
					if t.State == sdl.PRESSED {
						err := rr.Save(rs)
						if err != nil {
							fmt.Println(err)
						}
					}
				}
			}

		}
		renderer.SetDrawColor(255, 0, 0, 255)
		for _, points := range rs {
			for i := 0; i < len(points)-1; i++ {
				renderer.DrawLine(points[i].X, points[i].Y, points[i+1].X, points[i+1].Y)
			}
			renderer.DrawLine(points[len(points)-1].X, points[len(points)-1].Y, points[0].X, points[0].Y)
		}
		renderer.SetDrawColor(0, 255, 0, 255)
		for _, segment := range rb.Segments {
			renderer.DrawLine(segment.P1.X, segment.P1.Y, segment.P2.X, segment.P2.Y)
		}
		renderer.Present()
	}
}
