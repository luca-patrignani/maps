package politics_tui

import (
	"fmt"
	"io"

	"github.com/luca-patrignani/maps/politics"
)

type IO struct {
	In  io.Reader
	Out io.Writer
}

func (io IO) NewPoliticalEntity() (entity politics.PoliticalEntity, err error) {
	_, err = fmt.Fprintln(io.Out, "This prompt lets you create a new political entity")
	if err != nil {
		return
	}
	_, err = fmt.Fprintln(io.Out, "Name of the new political entity:")
	if err != nil {
		return
	}
	_, err = fmt.Fscanln(io.In, &entity.Name)
	if err != nil {
		return
	}
	_, err = fmt.Fprintln(io.Out, "Color to be used to show the new political entity")
	if err != nil {
		return
	}
	_, err = fmt.Fprintln(io.Out, "The color must be expressed as RBG, with a whitespace between each field")
	if err != nil {
		return
	}
	c := &entity.Color
	_, err = fmt.Fscanln(io.In, &c.R, &c.G, &c.B)
	c.A = 255
	return
}
