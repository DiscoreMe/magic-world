package world

import (
	"sync"

	"github.com/DiscoreMe/magic-world/entity"
)

// Field types
const (
	ZoneTypeEmpty = iota
	ZoneTypeLand
)

// Zone contains information about all cells and provides methods for working with them
type Zone struct {
	mux     sync.RWMutex
	z       map[ZonePos]ZoneCell
	width   int
	height  int
	maxSide int
}

// NewZone creates new zone
func NewZone(width, height int) *Zone {
	var maxSide int
	if width > height {
		maxSide = width
	} else {
		maxSide = height
	}

	zone := &Zone{
		z:       make(map[ZonePos]ZoneCell),
		width:   width,
		height:  height,
		maxSide: maxSide,
	}

	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			zone.SetCell(x, y, ZoneCell{Type: ZoneTypeLand})
		}
	}

	return zone
}

// ZonePos is ID of each zone.
// The map is an X*Y field. In my opinion, using a slice is inconvenient,
// so the values will be stored in map[ZonePos]ZoneCell.
// To avoid confusing positions, there is a formula that calculates the ZonePos coordinate:
// x * worldHeight + y
type ZonePos int

// pos calculates ZonePos
func (z *Zone) pos(x, y int) ZonePos {
	return ZonePos(x*z.maxSide + y)
}

// ZoneCell contains information about the cell
type ZoneCell struct {
	Type int
	Meta string

	x, y     int
	entities map[int64]entity.Entity
}

func (cell ZoneCell) X() int {
	return cell.x
}

func (cell ZoneCell) Y() int {
	return cell.y
}

// checkCoords returns true if x and y is valid for current zone
func (z *Zone) checkCoords(x, y int) bool {
	return x >= 0 && y >= 0 && x < z.width && y < z.height
}

func (z *Zone) forEach(fn func(cell ZoneCell)) {
	for x := 0; x < z.width; x++ {
		for y := 0; y < z.height; y++ {
			if cell, ok := z.Cell(x, y); ok {
				fn(cell)
			}
		}
	}
}

// Cell gets zone cell
func (z *Zone) Cell(x, y int) (ZoneCell, bool) {
	z.mux.RLock()
	defer z.mux.RUnlock()
	cell, ok := z.z[z.pos(x, y)]
	return cell, ok
}

// SetCell sets zone info
func (z *Zone) SetCell(x, y int, cell ZoneCell) {
	if !z.checkCoords(x, y) {
		return // ignore
	}

	z.mux.Lock()
	defer z.mux.Unlock()

	cell.x, cell.y = x, y
	cell.entities = make(map[int64]entity.Entity)

	z.z[z.pos(x, y)] = cell
}

// AddEntity adds new entity to the cell in specified coords
func (z *Zone) AddEntity(x, y int, ent entity.Entity) {
	if cell, ok := z.Cell(x, y); ok {
		cell.entities[ent.ID()] = ent
	}
}