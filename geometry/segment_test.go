package geometry

import "testing"

func TestIntersectionOdd(t *testing.T) {
	segment1 := Segment{Point{0, 0}, Point{5, 5}}
	segment2 := Segment{Point{0, 5}, Point{5, 0}}
	intersection, err := Intersection(segment1, segment2)
	if err != nil {
		t.Error(err)
	}
	if intersection != (Point{2.5, 2.5}) {
		t.Errorf("expected %v, actual %v", Point{2, 2}, intersection)
	}
}

func TestIntersectionEven(t *testing.T) {
	segment1 := Segment{Point{0, 0}, Point{4, 4}}
	segment2 := Segment{Point{0, 4}, Point{4, 0}}
	intersection, err := Intersection(segment1, segment2)
	if err != nil {
		t.Error(err)
	}
	if intersection != (Point{2, 2}) {
		t.Errorf("expected %v, actual %v", Point{2, 2}, intersection)
	}
}

func TestIntersectionNot(t *testing.T) {
	segment1 := Segment{Point{0, 0}, Point{5, 0}}
	segment2 := Segment{Point{0, 5}, Point{5, 5}}
	_, err := Intersection(segment1, segment2)
	if err == nil {
		t.Fail()
	}
}

func TestIntersectionEnd(t *testing.T) {
	segment1 := Segment{Point{0, 0}, Point{5, 0}}
	segment2 := Segment{Point{0, 5}, Point{5, 0}}
	intersection, err := Intersection(segment1, segment2)
	if err != nil {
		t.Error(err)
	}
	if intersection != (Point{5, 0}) {
		t.Fail()
	}
}

func TestIntersection(t *testing.T) {
	segment1 := Segment{Point{0, 0}, Point{1, 0}}
	segment2 := Segment{Point{1, 1}, Point{0, 1}}
	if _, err := Intersection(segment1, segment2); err == nil {
		t.Fail()
	}

}

func TestIntersectionHorizontalParallel(t *testing.T) {
	segment1 := Segment{Point{0, 0}, Point{10, 0}}
	segment2 := Segment{Point{0, 0}, Point{20, 0}}
	_, err := Intersection(segment1, segment2)
	if err != nil {
		t.Error(err)
	}
}

func TestIntersectionVerticalParallel(t *testing.T) {
	segment1 := Segment{Point{0, 0}, Point{0, 10}}
	segment2 := Segment{Point{0, 0}, Point{0, 20}}
	_, err := Intersection(segment1, segment2)
	if err != nil {
		t.Error(err)
	}
}

func TestIntersectionParallel(t *testing.T) {
	segment1 := Segment{Point{0, 0}, Point{5, 5}}
	segment2 := Segment{Point{2, 2}, Point{20, 20}}
	_, err := Intersection(segment1, segment2)
	if err != nil {
		t.Error(err)
	}
}

func TestIntersectionParallelNotIntersect(t *testing.T) {
	segment1 := Segment{Point{0, 0}, Point{3, 0}}
	segment2 := Segment{Point{4, 0}, Point{20, 0}}
	_, err := Intersection(segment1, segment2)
	if err == nil {
		t.Fatal()
	}
}

func TestIntersectionParallelNotIntersect2(t *testing.T) {
	segment2 := Segment{Point{0, 0}, Point{3, 0}}
	segment1 := Segment{Point{4, 0}, Point{20, 0}}
	_, err := Intersection(segment1, segment2)
	if err == nil {
		t.Fatal()
	}
}

func TestIntersection2(t *testing.T) {
	segment1 := Segment{Point{2, 0}, Point{1, 0}}
	segment2 := Segment{Point{1, 0}, Point{0, 0}}
	inter, err := Intersection(segment1, segment2)
	if err != nil {
		t.Fatal(err)
	}
	actual := Point{1, 0}
	if inter != actual {
		t.Fatal()
	}
}

func TestIntersectionParallelButNotIntersecting(t *testing.T) {
	s1 := Segment{P1: Point{X: 31, Y: 4}, P2: Point{X: 31, Y: 5}}
	s2 := Segment{P1: Point{X: 31, Y: 7}, P2: Point{X: 31, Y: 8}}
	inter, err := Intersection(s1, s2)
	if err == nil {
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
