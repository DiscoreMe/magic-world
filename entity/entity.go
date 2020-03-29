package entity

import (
	"go.uber.org/atomic"
)

type Entity interface {
	ID() int64
	Name() string
	X() int
	Y() int
	SetX(int)
	SetY(int)
	Step()
}

var nextEntityID atomic.Int64

func Up(ent Entity) {
	ent.SetY(ent.Y() - 1)
}
func Down(ent Entity) {
	ent.SetY(ent.Y() + 1)
}
func Left(ent Entity) {
	ent.SetX(ent.X() - 1)
}
func Right(ent Entity) {
	ent.SetX(ent.X() + 1)
}

type AroundPos struct {
	UpX    int
	UpY    int
	DownX  int
	DownY  int
	LeftX  int
	LeftY  int
	RightX int
	RightY int
}

func Around(ent Entity) AroundPos {
	x, y := ent.X(), ent.Y()
	return AroundPos{
		UpX:    x,
		UpY:    y + 1,
		DownX:  x,
		DownY:  y - 1,
		LeftX:  x - 1,
		LeftY:  y,
		RightX: x + 1,
		RightY: y,
	}
}
