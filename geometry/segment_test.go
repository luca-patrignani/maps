package geometry

import "testing"

func TestIntersectionOdd(t *testing.T) {
	segment1 := Segment{Point{0, 0}, Point{5, 5}}
	segment2 := Segment{Point{0, 5}, Point{5, 0}}
	intersection := Intersection(segment1, segment2)
	expected := Point{2.5, 2.5}
	if intersection != expected {
		t.Fatalf("expected %v, actual %v", expected, intersection)
	}
}

func TestIntersectionEven(t *testing.T) {
	segment1 := Segment{Point{0, 0}, Point{4, 4}}
	segment2 := Segment{Point{0, 4}, Point{4, 0}}
	intersection := Intersection(segment1, segment2)
	if intersection != (Point{2, 2}) {
		t.Fatalf("expected %v, actual %v", Point{2, 2}, intersection)
	}
}

func TestIntersectionNot(t *testing.T) {
	segment1 := Segment{Point{0, 0}, Point{5, 0}}
	segment2 := Segment{Point{0, 5}, Point{5, 5}}
	inter := Intersection(segment1, segment2)
	if inter != nil {
		t.Fail()
	}
}

func TestIntersectionEnd(t *testing.T) {
	segment1 := Segment{Point{0, 0}, Point{5, 0}}
	segment2 := Segment{Point{0, 5}, Point{5, 0}}
	intersection := Intersection(segment1, segment2)
	if intersection != (Point{5, 0}) {
		t.Fail()
	}
}

func TestIntersection(t *testing.T) {
	segment1 := Segment{Point{0, 0}, Point{1, 0}}
	segment2 := Segment{Point{1, 1}, Point{0, 1}}
	if inter := Intersection(segment1, segment2); inter != nil {
		t.Fail()
	}
}

func TestIntersectionHorizontalParallel(t *testing.T) {
	segment1 := Segment{Point{0, 0}, Point{10, 0}}
	segment2 := Segment{Point{0, 0}, Point{20, 0}}
	if inter := Intersection(segment1, segment2); inter != segment1 {
		t.Fatal(inter)
	}
}

func TestIntersectionVerticalParallel(t *testing.T) {
	segment1 := Segment{Point{0, 0}, Point{0, 10}}
	segment2 := Segment{Point{0, 0}, Point{0, 20}}
	if inter := Intersection(segment1, segment2); inter != segment1 {
		t.Fatal(inter)
	}
}

func TestIntersectionParallel(t *testing.T) {
	segment1 := Segment{Point{0, 0}, Point{5, 5}}
	segment2 := Segment{Point{2, 2}, Point{20, 20}}
	if inter := Intersection(segment1, segment2); inter != (Segment{Point{5, 5}, Point{2, 2}}) {
		t.Fatal(inter)
	}
}

func TestIntersectionParallelNotIntersect(t *testing.T) {
	segment1 := Segment{Point{0, 0}, Point{3, 0}}
	segment2 := Segment{Point{4, 0}, Point{20, 0}}
	if inter := Intersection(segment1, segment2); inter != nil {
		t.Fatal(inter)
	}
}

func TestIntersectionParallelNotIntersect2(t *testing.T) {
	segment2 := Segment{Point{0, 0}, Point{3, 0}}
	segment1 := Segment{Point{4, 0}, Point{20, 0}}
	if inter := Intersection(segment1, segment2); inter != nil {
		t.Fatal()
	}
}

func TestIntersection2(t *testing.T) {
	segment1 := Segment{Point{2, 0}, Point{1, 0}}
	segment2 := Segment{Point{1, 0}, Point{0, 0}}
	inter := Intersection(segment1, segment2)
	if inter != (Point{1, 0}) {
		t.Fatal(inter)
	}
}

func TestIntersectionParallelButNotIntersecting(t *testing.T) {
	s1 := Segment{P1: Point{X: 31, Y: 4}, P2: Point{X: 31, Y: 5}}
	s2 := Segment{P1: Point{X: 31, Y: 7}, P2: Point{X: 31, Y: 8}}
	inter := Intersection(s1, s2)
	if inter != nil {
		t.Fatal(inter)
	}
}

func TestIntersectionParallelEnd(t *testing.T) {
	s1 := Segment{P1: Point{X: 0, Y: 0}, P2: Point{X: 10, Y: 0}}
	s2 := Segment{P1: Point{X: 9, Y: 0}, P2: Point{X: 10, Y: 0}}
	if inter := Intersection(s1, s2); inter != (Segment{P1: Point{X: 10, Y: 0}, P2: Point{X: 9, Y: 0}}) {
		t.Fatal(inter)
	}
}

func TestIntersectionParallelEnd2(t *testing.T) {
	s1 := Segment{P1: Point{X: 0, Y: 0}, P2: Point{X: 9, Y: 0}}
	s2 := Segment{P1: Point{X: 9, Y: 0}, P2: Point{X: 10, Y: 0}}
	if inter := Intersection(s1, s2); inter != (Point{9, 0}) {
		t.Fatal(inter)
	}
}

func TestIntersectionParallelEnd3(t *testing.T) {
	s1 := Segment{P1: Point{X: 0, Y: 0}, P2: Point{X: 10, Y: 0}}
	s2 := Segment{P1: Point{X: 0, Y: 0}, P2: Point{X: 10, Y: 0}}
	if inter := Intersection(s1, s2); inter != s1 {
		t.Fatal(inter)
	}
}

func TestIsParallelToDiagonal(t *testing.T) {
	segment1 := Segment{Point{1, 1}, Point{2, 2}}
	segment2 := Segment{Point{5, 5}, Point{-2, -2}}
	if !segment1.IsParallelTo(segment2) {
		t.Fatal()
	}
}

func TestIsParallelToHorizontal(t *testing.T) {
	segment1 := Segment{Point{1, 0}, Point{2, 0}}
	segment2 := Segment{Point{5, 0}, Point{-2, 0}}
	if !segment1.IsParallelTo(segment2) {
		t.Fatal()
	}
}

func TestIsParallelToVertical(t *testing.T) {
	segment1 := Segment{Point{1, 0}, Point{1, 2}}
	segment2 := Segment{Point{5, 0}, Point{5, -2}}
	if !segment1.IsParallelTo(segment2) {
		t.Fatal()
	}
}

func TestNotIsParallelTo(t *testing.T) {
	segment1 := Segment{Point{0, 0}, Point{1, 0}}
	segment2 := Segment{Point{0, 0}, Point{0, 1}}
	if segment1.IsParallelTo(segment2) {
		t.Fatal()
	}
}

func TestContains(t *testing.T) {
	s := Segment{Point{1, 1}, Point{10, 10}}
	p := Point{4, 4}
	if !s.Contains(p) {
		t.Fatal(s, p)
	}
}

func TestContainsVertical(t *testing.T) {
	s := Segment{Point{1, 1}, Point{1, 10}}
	p := Point{1, 4}
	if !s.Contains(p) {
		t.Fatal(s, p)
	}
}

func TestNotContains(t *testing.T) {
	s := Segment{Point{1, 1}, Point{10, 10}}
	p := Point{20, 14}
	if s.Contains(p) {
		t.Fatal(s, p)
	}
}

func TestNotContainsVertical(t *testing.T) {
	s := Segment{Point{1, 1}, Point{1, 10}}
	p := Point{1, 40}
	if s.Contains(p) {
		t.Fatal(s, p)
	}
}
