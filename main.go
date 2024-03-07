package main

import (
	_ "fmt"

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

	var winWidth, winHeight int32 = 800, 600
	renderer.SetDrawColor(0, 255, 0, 255)
	renderer.FillRect(&sdl.Rect{
		X: 10,
		Y: 10,
		W: 100,
		H: 100,
	})

	points := []sdl.Point{}
	running := true
	for running {
		renderer.SetDrawColor(0, 0, 0, 255)
		renderer.Clear()
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch t := event.(type) {
			case *sdl.QuitEvent:
				running = false
			case *sdl.MouseMotionEvent:
				// fmt.Println("Mouse", t.Which, "moved by", t.XRel, t.YRel, "at", t.X, t.Y)
				points = append(points, sdl.Point{X: t.X, Y: t.Y})
				// case *sdl.MouseButtonEvent:
				// 	if t.State == sdl.PRESSED {
				// 		fmt.Println("Mouse", t.Which, "button", t.Button, "pressed at", t.X, t.Y)
				// 	} else {
				// 		fmt.Println("Mouse", t.Which, "button", t.Button, "released at", t.X, t.Y)
				// 	}
				// }
			}
		}
		var vx, vy = make([]int16, 3), make([]int16, 3)
		vx[0] = int16(winWidth / 3)
		vy[0] = int16(winHeight / 3)
		vx[1] = int16(winWidth * 2 / 3)
		vy[1] = int16(winHeight / 3)
		vx[2] = int16(winWidth / 2)
		vy[2] = int16(winHeight * 2 / 3)
		gfx.FilledPolygonColor(renderer, vx, vy, sdl.Color{R: 255, G: 0, B: 0, A: 255})
		renderer.SetDrawColor(255, 0, 0, 255)
		for i := 0; i < len(points)-1; i++ {
			renderer.DrawLine(points[i].X, points[i].Y, points[i+1].X, points[i+1].Y)
		}
		renderer.Present()
	}
}
