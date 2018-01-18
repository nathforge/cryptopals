package cursor

import (
	"fmt"
	"testing"
)

type ExpectedPosFromXY struct {
	pos, x, y   int
	outOfBounds bool
}

type ExpectedXYFromPos struct {
	x, y, pos   int
	outOfBounds bool
}

const (
	width  = 10
	height = 5
)

var expectedPosFromXY = []ExpectedPosFromXY{
	{x: -1, y: 0, outOfBounds: true},
	{x: 0, y: 0, pos: 0},
	{x: 0, y: -1, outOfBounds: true},
	{x: 9, y: 0, pos: 9},
	{x: 10, y: 0, outOfBounds: true},
	{x: 0, y: 1, pos: 10},
	{x: 9, y: 4, pos: 49},
	{x: 0, y: 5, outOfBounds: true},
}

var expectedXYFromPos = []ExpectedXYFromPos{
	{pos: -1, outOfBounds: true},
	{pos: 0, x: 0, y: 0},
	{pos: 9, x: 9, y: 0},
	{pos: 10, x: 0, y: 1},
	{pos: 49, x: 9, y: 4},
	{pos: 50, outOfBounds: true},
}

func getCursor(t *testing.T) *Cursor {
	pos := 0
	return &Cursor{
		GetExtent: func() (int, int) {
			return width, height
		},
		GetPos: func() int {
			t.Logf("GetPos() = %d\n", pos)
			return pos
		},
		SetPos: func(newPos int) error {
			pos = newPos
			t.Logf("SetPos(%d)\n", pos)
			return nil
		},
	}
}

func TestGetPosFromXY(t *testing.T) {
	for _, test := range expectedPosFromXY {
		t.Run(fmt.Sprintf("%+v", test), func(t *testing.T) {
			pos, err := getCursor(t).GetPosFromXY(test.x, test.y)
			if test.outOfBounds && err != ErrOutOfBounds {
				t.Fatalf("Expected ErrOutOfBounds for xy=%d,%d but received err=%v, pos=%d", test.x, test.y, err, pos)
			} else if !test.outOfBounds && pos != test.pos {
				t.Fatalf("Expected pos=%d for xy=%d,%d but received err=%v, pos=%d", test.pos, test.x, test.y, err, pos)
			}
		})
	}
}

func TestGetXYFromPos(t *testing.T) {
	for _, test := range expectedXYFromPos {
		t.Run(fmt.Sprintf("%+v", test), func(t *testing.T) {
			x, y, err := getCursor(t).GetXYFromPos(test.pos)
			if test.outOfBounds && err != ErrOutOfBounds {
				t.Fatalf("Expected ErrOutOfBounds for pos=%d but received err=%v, xy=%d,%d", test.pos, err, x, y)
			} else if !test.outOfBounds && (x != test.x || y != test.y) {
				t.Fatalf("Expected xy=%d,%d for pos=%d but received err=%v, xy=%d,%d", test.x, test.y, test.pos, err, x, y)
			}
		})
	}
}

func TestGetXY(t *testing.T) {
	for _, test := range expectedXYFromPos {
		t.Run(fmt.Sprintf("%+v", test), func(t *testing.T) {
			c := getCursor(t)
			c.SetPos(test.pos)
			x, y, err := c.GetXY()
			if test.outOfBounds && err != ErrOutOfBounds {
				t.Fatalf("Expected ErrOutOfBounds for pos=%d but received err=%v, xy=%d,%d", test.pos, err, x, y)
			} else if !test.outOfBounds && (x != test.x || y != test.y) {
				t.Fatalf("Expected xy=%d,%d for pos=%d but received err=%v, xy=%d,%d", test.x, test.y, test.pos, err, x, y)
			}
		})
	}
}

func TestSetXY(t *testing.T) {
	for _, test := range expectedPosFromXY {
		t.Run(fmt.Sprintf("%+v", test), func(t *testing.T) {
			c := getCursor(t)
			err := c.SetXY(test.x, test.y)
			pos := c.GetPos()
			if test.outOfBounds && err != ErrOutOfBounds {
				t.Fatalf("Expected ErrOutOfBounds for xy=%d,%d but received err=%v, pos=%d", test.x, test.y, err, pos)
			} else if !test.outOfBounds && pos != test.pos {
				t.Fatalf("Expected pos=%d for xy=%d,%d but received err=%v, pos=%d", test.pos, test.x, test.y, err, pos)
			}
		})
	}
}
