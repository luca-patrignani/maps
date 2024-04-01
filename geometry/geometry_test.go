package geometry

import "testing"

func TestIntersectionOdd(t *testing.T) {
	segment1 := Segment{Point{0, 0}, Point{5, 5}}
	segment2 := Segment{Point{0, 5}, Point{5, 0}}
	intersection, err := Intersection(segment1, segment2)
	if err != nil {
		t.Error(err)
	}
	if intersection != (Point{2, 2}) {
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
