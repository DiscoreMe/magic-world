package world

import (
	"sync"

	"github.com/DiscoreMe/magic-world/entity"
)

// World is main data when saves info
// Each cell is a separate object.
// The world structure combines them
type World struct {
	mux  sync.RWMutex
	days int64
	zone *Zone
}

func (w *World) Days() int64 {
	return w.days
}

func (w *World) Width() int {
	return w.zone.width
}

func (w *World) Height() int {
	return w.zone.height
}

// NewWorld creates new world
func NewWorld(width, height int) *World {
	return &World{zone: NewZone(width, height)}
}

// AddEntity adds new entity to the world
func (w *World) AddEntity(x, y int, ent entity.Entity) {
	w.zone.AddEntity(x, y, ent)
}

func (w *World) Step() {
	w.mux.Lock()
	defer w.mux.Unlock()

	w.days++

	var defers []func()
	w.zone.forEach(func(cell ZoneCell) {
		for id, ent := range cell.entities {
			ent.Step()

			around := entity.Around(ent)
			if around.UpY < 0 {
				entity.Down(ent)
			} else if around.DownY >= w.Height() {
				entity.Up(ent)
			} else if around.RightX >= w.Width() {
				entity.Left(ent)
			} else if around.LeftX < 0 {
				entity.Right(ent)
			}

			x, y := ent.X(), ent.Y()
			if x != cell.x || y != cell.y {
				defers = append(defers, func() {
					// change entity position
					// after all steps
					delete(cell.entities, id)
					if cell, ok := w.zone.Cell(x, y); ok {
						cell.entities[id] = ent
					}
				})
			}
		}
	})

	for _, fn := range defers {
		fn()
	}
}
