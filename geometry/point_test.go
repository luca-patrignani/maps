package geometry

import "testing"

func TestDistance(t *testing.T) {
	p1 := Point{X: 0, Y: 3}
	p2 := Point{X: 4, Y: 0}
	if p1.Distance(p2) != 5 {
		t.Fatal()
	}
}
