package cursor

import (
	"errors"
)

var (
	ErrOutOfBounds error = errors.New("Out of bounds")
)

type GetExtent func() (int, int)

type Cursor struct {
	getExtent GetExtent
	pos       int
}

func New(getExtent GetExtent) *Cursor {
	return &Cursor{
		getExtent: getExtent,
	}
}

func (s *Cursor) GetPos() int {
	return s.pos
}

func (s *Cursor) SetPos(pos int) error {
	w, h := s.getExtent()
	if pos < 0 || pos >= w*h {
		return ErrOutOfBounds
	}
	s.pos = pos
	return nil
}

func (s *Cursor) GetXY() (int, int) {
	w, _ := s.getExtent()
	return s.pos % w, s.pos / w
}

func (s *Cursor) SetXY(x, y int) error {
	pos, err := s.posFromXY(x, y)
	if err != nil {
		return err
	}
	s.pos = pos
	return nil
}

func (s *Cursor) GetX() int {
	x, _ := s.GetXY()
	return x
}

func (s *Cursor) GetY() int {
	_, y := s.GetXY()
	return y
}

func (s *Cursor) SetX(x int) error {
	return s.SetXY(x, s.GetY())
}

func (s *Cursor) SetY(y int) error {
	return s.SetXY(s.GetX(), y)
}

func (s *Cursor) posFromXY(x, y int) (int, error) {
	w, h := s.getExtent()
	if x < 0 || x >= w || y < 0 || y >= h {
		return 0, ErrOutOfBounds
	}
	return x + (y * w), nil
}
