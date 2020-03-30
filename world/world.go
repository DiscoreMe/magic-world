package world

import (
	"fmt"
	"math/rand"
	"sync"
	"time"

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

	w.zone.forEachForest(func(cell ZoneCell) {
		rn := int64(rand.Intn(10))
		ln := time.Now().Unix() % 10
		x, y := cell.X(), cell.Y()
		cell.Type = ZoneTypeForest

		type task struct {
			x, y int
		}
		var tasks []task

		if ln > rn {
			tasks = append(tasks, task{x: x + 1, y: y})
			tasks = append(tasks, task{x: x, y: y - 1})
		} else {
			tasks = append(tasks, task{x: x - 1, y: y})
			tasks = append(tasks, task{x: x, y: y + 1})
		}

		for _, t := range tasks {
			if w.zone.IsZoneType(t.x, t.y, ZoneTypeLand) {
				fmt.Println(t.x, t.y)
				w.zone.SetCell(t.x, t.y, cell)
			}
		}
	})

	for _, fn := range defers {
		fn()
	}
}
