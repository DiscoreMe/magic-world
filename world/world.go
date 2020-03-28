package world

import (
	"sync"
)

const worldWidth, worldHeight = 100, 100

// PosZone is ID of each zone.
// The map is an X*Y field. In my opinion, using a slice is inconvenient,
// so the values will be stored in map[PosZone]zoneInfo.
// To avoid confusing positions, there is a formula that calculates the PosZone coordinate:
// X = X * width_field * height_field * 2
// Y = Y * width_field * height_field * 3
//
// For example, we need to get data for coordinates 5 and 6, and the field size is 100 by 50:
// PosZone = 5 * 100 * 50 * 2 + 6 * 100 * 50 * 3
type PosZone int

// World is main data when saves info
// Each cell is a separate object.
// The world structure combines them
type World struct {
	Zone *Zone
}

// Zone contains information about all cells and provides methods for working with them
type Zone struct {
	zMux sync.RWMutex
	z map[PosZone]zoneInfo
}

// zoneInfo contains information about the cell
type zoneInfo struct {
	X int
	Y int
	Meta string
}

// calcZone calculates PosZone
func calcZone(x, y int) PosZone {
	return PosZone(x * worldWidth * worldHeight * 2 + y * worldWidth * worldHeight * 3)
}

// zone gets zone info
func (z *Zone) zone(x, y int) zoneInfo {
	z.zMux.RLock()
	info := z.z[calcZone(x, y)]
	z.zMux.RUnlock()
	return info
}

// setZone sets zone info
func (z *Zone) setZone(x, y int, info zoneInfo) {
	z.zMux.Lock()
	z.z[calcZone(x, y)] = info
	z.zMux.Unlock()
}

// Meta gets meta of the zone info
func (z *Zone) Meta(x, y int) string {
	return z.zone(x, y).Meta
}

// Meta sets meta of the zone info
func (z *Zone) SetMeta(x, y int, v string) {
	zone := z.zone(x, y)
	zone.Meta = v
	z.setZone(x, y, zone)
}

// NewZone creates new zone
func NewZone() *Zone {
	return &Zone{
		zMux: sync.RWMutex{},
		z: make(map[PosZone]zoneInfo),
	}
}

// NewWorld creates new world
func NewWorld() *World {
	return &World{
		Zone: NewZone(),
	}
}