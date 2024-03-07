package main

import (
	"github.com/veandco/go-sdl2/sdl"
)

func main() {
	window, err := sdl.CreateWindow("test", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		800, 600, sdl.WINDOW_SHOWN)
	if err != nil {
		panic(err)
	}
	defer window.Destroy()

	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		panic(err)
	}
	defer renderer.Destroy()

	running := true

	for running {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				running = false
			}
		}
		renderer.SetDrawColor(0, 0, 0, 255)
		renderer.Clear()

		renderer.SetDrawColor(255, 0, 0, 255)
		renderer.DrawRect(&sdl.Rect{
			X: 0,
			Y: 0,
			W: 100,
			H: 100,
		})

		renderer.DrawLines([]sdl.Point{
			{X: 0, Y: 0},
			{X: 0, Y: 20},
			{X: 0, Y: 30},
			{X: 20, Y: 30},
			{X: 0, Y: 0},
		})
		renderer.Present()

	}
}
