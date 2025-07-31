package main

import "github.com/luca-patrignani/maps/geometry"

type geography[T any] interface {
	DrawLine(p1, p2 geometry.Point, t T)
	FillWith(p geometry.Point, foreground, background T)
	At(p geometry.Point) (T, bool)
}
