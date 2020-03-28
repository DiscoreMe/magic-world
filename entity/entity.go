package entity

type Entity interface {
	ID() int64
	Name() string
	Health() int
	Step()
	SetPos(int, int)
	Pos() (int, int)
}

var lastEntityID int64 = 0

func nextEntityID() int64 {
	lastEntityID++
	return lastEntityID
}
