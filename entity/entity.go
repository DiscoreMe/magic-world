package entity

type Entity interface {
	ID() int64
	Name() string
	Health() int
	Step()
}

var lastEntityID int64 = 0

func nextEntityID() int64 {
	lastEntityID++
	return lastEntityID
}
