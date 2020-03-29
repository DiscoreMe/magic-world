package world

import (
	"encoding/json"
	entity2 "github.com/DiscoreMe/magic-world/entity"
	"os"
	"sync"
)

var worldWidth, worldHeight = 0, 0
var maxWorldSize = 0

// PosZone is ID of each zone.
// The map is an X*Y field. In my opinion, using a slice is inconvenient,
// so the values will be stored in map[PosZone]zoneInfo.
// To avoid confusing positions, there is a formula that calculates the PosZone coordinate:
// x * worldHeight + y
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
	X        int
	Y        int
	Meta     string
	Type     int
	Entities []entity2.Entity
}

// calcZone calculates PosZone
func calcZone(x, y int) PosZone {
	return PosZone(x*maxWorldSize + y)
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
	if x < 0 || x > worldWidth || y < 0 || y > worldHeight {
		return
	}
	info.X = x
	info.Y = y
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

func (z *Zone) addEntity(x, y int, entity entity2.Entity) {
	zone := z.zone(x, y)
	zone.Entities = append(zone.Entities, entity)
	z.setZone(x, y, zone)
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
	herMux        sync.RWMutex
	Heroes        []*entity2.Hero
	days          int64
}

// NewWorld creates new world
func NewWorld(width, height int) *World {
	worldWidth, worldHeight = width, height
	maxWorldSize = width
	if height > width {
		maxWorldSize = height
	}
	return &World{
		Zone:   NewZone(),
		width:  width,
		height: height,
		herMux: sync.RWMutex{},
		Heroes: make([]*entity2.Hero, 0),
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

func (w *World) AddEntity(x, y int, entity entity2.Entity) {
	entity.SetPos(x, y)
	switch v := entity.(type) {
	case *entity2.Hero:
		w.Heroes = append(w.Heroes, v)
	default:
		w.Zone.addEntity(x, y, entity)
	}
}

type ExportZone struct {
	X        int            `json:"x"`
	Y        int            `json:"y"`
	Type     int            `json:"type"`
	Entities []ExportEntity `json:"entities,omitempty"`
}

type ExportEntity struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
	X    int    `json:"x"`
	Y    int    `json:"y"`
}

type ExportData struct {
	Days   int64        `json:"days"`
	Width  int          `json:"width"`
	Height int          `json:"height"`
	Zones  []ExportZone `json:"zones,omitempty"`
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
		Days:   w.days,
	}

	for y := 0; y < worldHeight; y++ {
		for x := 0; x < worldWidth; x++ {
			info := w.Zone.zone(x, y)
			zone := ExportZone{
				X:    x,
				Y:    y,
				Type: info.Type,
			}
			for _, ent := range info.Entities {
				zone.Entities = append(zone.Entities, ExportEntity{
					ID:   ent.ID(),
					Name: ent.Name(),
				})
			}

			for _, h := range w.Heroes {
				xh, yh := h.Pos()
				if x == xh && y == yh {
					zone.Entities = append(zone.Entities, ExportEntity{
						ID:   h.ID(),
						Name: h.Name(),
						X:    x,
						Y:    y,
					})
				}
			}

			exportData.Zones = append(exportData.Zones, zone)
		}
	}

	b, err := json.Marshal(exportData)
	if err != nil {
		return err
	}

	_, err = f.Write(b)
	return err
}

func (w *World) Days() int64 {
	return w.days
}

func (w *World) Step() {
	w.days++
	for _, hero := range w.Heroes {
		hero.Step()
		aroundPos := hero.Around()
		if aroundPos.UpY < 0 {
			hero.Down()
		} else if aroundPos.DownY >= w.height {
			hero.Up()
		} else if aroundPos.RightX >= w.width {
			hero.Left()
		} else if aroundPos.LeftX < 0 {
			hero.Right()
		}
	}
}
