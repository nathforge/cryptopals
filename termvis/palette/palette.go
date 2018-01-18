package palette

import (
	"image/color"
)

type Palette [8]color.RGBA

func (s Palette) Index(color Color) int {
	return int(color)
}

func (s Palette) Get(color Color) *color.RGBA {
	return &s[s.Index(color)]
}
