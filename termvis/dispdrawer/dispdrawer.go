package dispdrawer

import (
	"image"
	"image/draw"

	"github.com/nathforge/cryptopals/termvis/display"
	"golang.org/x/image/font"
	"golang.org/x/image/math/fixed"
)

type DisplayDrawer struct {
	Display           *display.Display
	CellPaddingLeft   int
	CellPaddingTop    int
	CellPaddingRight  int
	CellPaddingBottom int
	NormalFace        font.Face
	BoldFace          font.Face
	CellW             int
	CellH             int
}

func New(display *display.Display, normalFace, boldFace font.Face) *DisplayDrawer {
	s := &DisplayDrawer{}
	s.Display = display
	s.CellPaddingLeft = 0
	s.CellPaddingTop = 0
	s.CellPaddingRight = 0
	s.CellPaddingBottom = 0
	s.NormalFace = normalFace
	s.BoldFace = boldFace
	s.RecalculateCellWH()
	return s
}

func (s *DisplayDrawer) RecalculateCellWH() {
	textW := font.MeasureBytes(s.NormalFace, []byte{'A'}).Round()
	textH := s.NormalFace.Metrics().Height.Round()
	if s.BoldFace != nil {
		textW = maxInt(textW, font.MeasureBytes(s.BoldFace, []byte{'A'}).Round())
		textH = maxInt(textH, s.BoldFace.Metrics().Height.Round())
	}
	s.CellW = s.CellPaddingLeft + s.CellPaddingRight + textW
	s.CellH = s.CellPaddingTop + s.CellPaddingBottom + textH
}

func (s *DisplayDrawer) ImageSize() image.Rectangle {
	s.RecalculateCellWH()
	w := s.CellW * s.Display.Width
	h := s.CellH * s.Display.Height
	return image.Rect(0, 0, w, h)
}

func (s *DisplayDrawer) Draw(img *image.RGBA) {
	s.RecalculateCellWH()

	cellX := 0
	cellY := 0
	for _, cell := range s.Display.Cells {
		x := cellX * s.CellW
		y := cellY * s.CellH

		draw.Draw(img, image.Rect(x, y, x+s.CellW, y+s.CellH), image.NewUniform(cell.BG), image.ZP, draw.Src)

		cellX++
		if cellX >= s.Display.Width {
			cellX = 0
			cellY++
		}
	}

	cellX = 0
	cellY = 0
	for _, cell := range s.Display.Cells {
		x := cellX * s.CellW
		y := cellY * s.CellH

		if cell.Content != 0x00 {
			face := s.NormalFace
			if cell.Bold && s.BoldFace != nil {
				face = s.BoldFace
			}
			metrics := face.Metrics()

			textX := s.CellPaddingLeft + x
			textY := s.CellPaddingTop + y + metrics.Height.Round() - metrics.Descent.Round()

			drawer := &font.Drawer{
				Dst:  img,
				Src:  image.NewUniform(cell.FG),
				Face: face,
				Dot:  fixed.P(textX, textY),
			}
			drawer.DrawString(string(cell.Content))

			if cell.Bold && s.BoldFace == nil {
				drawer.Dot = fixed.P(textX+1, textY)
				drawer.DrawString(string(cell.Content))
			}
		}

		cellX++
		if cellX >= s.Display.Width {
			cellX = 0
			cellY++
		}
	}
}

func maxInt(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
