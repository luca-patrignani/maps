package main

import (
	"fmt"
	"reflect"

	"github.com/luca-patrignani/maps/geometry"
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
	points := []geometry.Point{}
	regions := []geometry.Region{}
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
				// fmt.Println("Mouse", t.Which, "moved by", t.XRel, t.YRel, "at", t.X, t.Y)
				if pressed {
					newPoint := geometry.Point{X: t.X/scale, Y: t.Y/scale}
					if len(points) == 0 || !reflect.DeepEqual(newPoint, points[len(points)-1]) {
						points = append(points, newPoint)
						fmt.Println(points)
						region, err := geometry.NewRegion(points)
						if err == nil {
							regions = append(regions, region)
							points = []geometry.Point{}
						}
					}
				}
			case *sdl.MouseButtonEvent:
				if t.State == sdl.PRESSED {
					pressed = true
					fmt.Println("Mouse", t.Which, "button", t.Button, "pressed at", t.X, t.Y)
				} else {
					pressed = false
					fmt.Println("Mouse", t.Which, "button", t.Button, "released at", t.X, t.Y)
				}
			}
		}
		renderer.SetDrawColor(255, 0, 0, 255)
		for _, points := range regions {
			for i := 0; i < len(points)-1; i++ {
				renderer.DrawLine(points[i].X, points[i].Y, points[i+1].X, points[i+1].Y)
			}
			renderer.DrawLine(points[len(points)-1].X, points[len(points)-1].Y, points[0].X, points[0].Y)
		}
		renderer.SetDrawColor(0, 255, 0, 255)
		for _, point := range points {
			renderer.DrawPoint(point.X, point.Y)
		}
		renderer.Present()
	}
}

