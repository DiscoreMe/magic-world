package world

import (
	"encoding/json"
	"fmt"
	"os"
	"sync"
)

var worldWidth, worldHeight = 0, 0

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

// Field types
const (
	TypeZoneEmpty = iota
	TypeZoneLand
)

// Zone contains information about all cells and provides methods for working with them
type Zone struct {
	zMux sync.RWMutex
	z    map[PosZone]zoneInfo
}

// zoneInfo contains information about the cell
type zoneInfo struct {
	X    int
	Y    int
	Meta string
	Type int
}

// calcZone calculates PosZone
func calcZone(x, y int) PosZone {
	return PosZone(x*worldWidth*worldHeight*2 + y*worldWidth*worldHeight*3)
}

// zone gets zone info
func (z *Zone) zone(x, y int) zoneInfo {
	z.zMux.RLock()
	info := z.z[calcZone(x, y)]
	z.zMux.RUnlock()
	fmt.Println(x, y, info.X, info.Y, info.Type)
	return info
}

// setZone sets zone info
func (z *Zone) setZone(x, y int, info zoneInfo) {
	if x < 0 || x > worldWidth || y < 0 || y > worldHeight {
		return
	}
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

func (z *Zone) Type(x, y int) int {
	return z.zone(x, y).Type
}

// NewZone creates new zone
func NewZone() *Zone {
	return &Zone{
		zMux: sync.RWMutex{},
		z:    make(map[PosZone]zoneInfo),
	}
}

// World is main data when saves info
// Each cell is a separate object.
// The world structure combines them
type World struct {
	Zone          *Zone
	width, height int
}

// NewWorld creates new world
func NewWorld(width, height int) *World {
	worldWidth, worldHeight = width, height
	return &World{
		Zone:   NewZone(),
		width:  width,
		height: height,
	}
}

func (w *World) CreateLand() {
	for y := 0; y < worldHeight; y++ {
		for x := 0; x < worldWidth; x++ {
			w.Zone.setZone(x, y, zoneInfo{
				X:    x,
				Y:    y,
				Type: TypeZoneLand,
			})
		}
	}
}

type ExportZone struct {
	X    int `json:"x"`
	Y    int `json:"y"`
	Type int `json:"type"`
}

type ExportData struct {
	Width  int          `json:"width"`
	Height int          `json:"height"`
	Zones  []ExportZone `json:"zones"`
}

func (w *World) ExportToJSON(filename string) error {
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	exportData := ExportData{
		Width:  w.width,
		Height: w.height,
	}

	for y := 0; y < worldHeight; y++ {
		for x := 0; x < worldWidth; x++ {
			info := w.Zone.zone(x, y)
			exportData.Zones = append(exportData.Zones, ExportZone{
				X:    x,
				Y:    y,
				Type: info.Type,
			})
		}
	}

	b, err := json.Marshal(exportData)
	if err != nil {
		return err
	}

	_, err = f.Write(b)
	return err
}
