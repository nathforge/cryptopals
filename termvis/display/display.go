package display

import (
	"image/color"

	ansiterm "github.com/Azure/go-ansiterm"
	"github.com/nathforge/cryptopals/termvis/ansievent"
	"github.com/nathforge/cryptopals/termvis/cursor"
	"github.com/nathforge/cryptopals/termvis/palette"
)

type Cell struct {
	FG      *color.RGBA
	BG      *color.RGBA
	Bold    bool
	Content rune
}

type Display struct {
	Width           int
	Height          int
	Palette         palette.Palette
	DefaultFG       *color.RGBA
	DefaultBG       *color.RGBA
	DefaultBold     bool
	CurrentFG       *color.RGBA
	CurrentBG       *color.RGBA
	CurrentBold     bool
	Cursor          *cursor.Cursor
	TabSize         int
	Cells           []Cell
	BeforeAnsiEvent func(interface{})
	AfterAnsiEvent  func(interface{})
	ansiEvents      []interface{}
	ansiParser      *ansiterm.AnsiParser
}

func New(width, height int) *Display {
	s := &Display{}

	s.Width = width
	s.Height = height

	s.Palette = palette.Default

	s.DefaultFG = s.Palette.Get(palette.White)
	s.DefaultBG = s.Palette.Get(palette.Black)
	s.DefaultBold = false

	s.CurrentFG = s.DefaultFG
	s.CurrentBG = s.DefaultBG
	s.CurrentBold = s.DefaultBold

	s.Cursor = cursor.New(func() (int, int) { return s.Width, s.Height })

	s.TabSize = 8

	for i := 0; i < width*height; i++ {
		s.Cells = append(s.Cells, Cell{
			FG:   s.CurrentFG,
			BG:   s.CurrentBG,
			Bold: s.CurrentBold,
		})
	}

	s.ansiParser = ansiterm.CreateParser("Ground", &ansievent.Handler{Events: &s.ansiEvents})

	return s
}

func (s *Display) WriteBytes(bytes []byte) error {
	s.ansiParser.Parse(bytes)
	for _, event := range s.ansiEvents {
		if s.BeforeAnsiEvent != nil {
			s.BeforeAnsiEvent(event)
		}
		switch v := event.(type) {
		case ansievent.Execute:
			switch v.B {
			case '\t':
				for i := 0; i < s.TabSize; i++ {
					if err := s.WriteContentRune(' '); err != nil {
						return err
					}
				}
			case '\n':
				if err := s.Cursor.SetXY(0, s.Cursor.GetY()+1); err != nil {
					return err
				}
			case '\r':
				if err := s.Cursor.SetX(0); err != nil {
					return err
				}
			}
		case ansievent.Print:
			// TODO: Buffer bytes, decode utf8
			s.WriteContentRune(rune(v.B))
		case ansievent.SGR:
			if len(v.Params) == 0 {
				// Reset
				s.CurrentFG = s.DefaultFG
				s.CurrentBG = s.DefaultBG
				s.CurrentBold = s.DefaultBold
			}
			for _, param := range v.Params {
				switch {
				// Reset/normal
				case param == 0:
					s.CurrentFG = s.DefaultFG
					s.CurrentBG = s.DefaultBG
					s.CurrentBold = s.DefaultBold
				// Bold
				case param == 1:
					s.CurrentBold = true
				// Foreground color
				case param >= 30 && param <= 37:
					s.CurrentFG = s.Palette.Get(palette.Color(param - 30))
				// Background color
				case param >= 40 && param <= 47:
					s.CurrentBG = s.Palette.Get(palette.Color(param - 40))
				}
			}
		case ansievent.ED:
			if v.Param == 2 {
				// TODO: Clear screen
				s.Cursor.SetPos(0)
			}
		}
		if s.AfterAnsiEvent != nil {
			s.AfterAnsiEvent(event)
		}
	}
	return nil
}

func (s *Display) WriteContentRune(r rune) error {
	cell := &s.Cells[s.Cursor.GetPos()]
	cell.FG = s.CurrentFG
	cell.BG = s.CurrentBG
	cell.Bold = s.CurrentBold
	cell.Content = r
	if err := s.Cursor.SetPos(s.Cursor.GetPos() + 1); err != nil {
		return err
	}
	return nil
}
