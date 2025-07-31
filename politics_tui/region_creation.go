package politics_tui

import (
	"fmt"
	"image/color"
	"io"

	"github.com/luca-patrignani/maps/politics"
)

type IO struct {
	In  io.Reader
	Out io.Writer
}

func (io IO) NewPoliticalEntity() (name politics.PoliticalEntity, c color.RGBA, err error) {
	_, err = fmt.Fprintln(io.Out, "This prompt lets you create a new political entity")
	if err != nil {
		return
	}
	_, err = fmt.Fprintln(io.Out, "Name of the new political entity:")
	if err != nil {
		return
	}
	_, err = fmt.Fscanln(io.In, &name)
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
	_, err = fmt.Fscanln(io.In, &c.R, &c.G, &c.B)
	c.A = 255
	return
}
