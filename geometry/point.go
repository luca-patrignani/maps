package geometry

import "encoding/json"

type Point struct {
	X int
	Y int
}

func (p Point) MarshalText() (text []byte, err error) {
    type noMethod Point
    return json.Marshal(noMethod(p))
}

func (p *Point) UnmarshalText(text []byte) error {
    type noMethod Point
    return json.Unmarshal(text, (*noMethod)(p))
}
