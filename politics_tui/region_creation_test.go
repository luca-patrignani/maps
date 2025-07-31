package politics_tui

import (
	"image/color"
	"io"
	"strings"
	"testing"
)

func TestCreation(t *testing.T) {
	r := strings.NewReader("Italy\n100 100 100")
	io := IO{In: r, Out: io.Discard}
	name, c, err := io.NewPoliticalEntity()
	if err != nil {
		t.Fatal(err)
	}
	if name != "Italy" {
		t.Fatal()
	}
	if c != (color.RGBA{100, 100, 100, 255}) {
		t.Fatal(c)
	}
}

func TestEmptyName(t *testing.T) {
	r := strings.NewReader("\n100 100 100")
	io := IO{In: r, Out: io.Discard}
	_, _, err := io.NewPoliticalEntity()
	if err == nil {
		t.Fatal()
	}
}

func TestInvalidColor(t *testing.T) {
	r := strings.NewReader("\n100 100")
	io := IO{In: r, Out: io.Discard}
	_, _, err := io.NewPoliticalEntity()
	if err == nil {
		t.Fatal()
	}
}
