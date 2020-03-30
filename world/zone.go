package world

import (
	"math/rand"
	"sync"
	"time"

	"github.com/DiscoreMe/magic-world/entity"
	"github.com/aquilax/go-perlin"
)

// Field types
const (
	ZoneTypeEmpty = iota
	ZoneTypeStone
	ZoneTypeWater
	ZoneTypeLand
	ZoneTypeForest
	ZoneTypeMax
)

var ZoneTypeNames = map[int]string{
	ZoneTypeStone:  "stone",
	ZoneTypeWater:  "water",
	ZoneTypeLand:   "land",
	ZoneTypeForest: "forest",
}

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

	for i := 0; i < width*height/4; i++ {
		x, y := rand.Intn(width), rand.Intn(height)
		zone.SetCell(x, y, ZoneCell{Type: ZoneTypeForest})
	}

	const alpha, beta, n = 2., 2., 3
	const oldMin, oldMax = -10, 10
	gen := perlin.NewPerlin(alpha, beta, n, time.Now().Unix())

	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			noise := int(gen.Noise2D(float64(x)/10, float64(y)/10) * 10)
			ztype := convertRange(noise, oldMin, oldMax, ZoneTypeEmpty+1, ZoneTypeMax)
			zone.SetCell(x, y, ZoneCell{Type: ztype})
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

func (z *Zone) forEachForest(fn func(cell ZoneCell)) {
	for y := 0; y < z.height; y++ {
		for x := 0; x < z.width; x++ {
			if cell, ok := z.Cell(x, y); ok {
				if cell.Type == ZoneTypeForest {
					fn(cell)
				}
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

func (z *Zone) IsZoneType(x, y, typeZone int) bool {
	cell, exist := z.Cell(x, y)
	if !exist {
		return exist
	}
	return cell.Type == typeZone
}

// AddEntity adds new entity to the cell in specified coords
func (z *Zone) AddEntity(x, y int, ent entity.Entity) {
	if cell, ok := z.Cell(x, y); ok {
		cell.entities[ent.ID()] = ent
	}
}
