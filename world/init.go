package world

import (
	"github.com/DiscoreMe/magic-world/zone"
)

func (w *World) initWorld() {

	zones := make([]*zone.Zone, 0)
	zones = append(
		zones,
		zone.NewZone(zone.ZoneTypeCity, "Святые земли", 1),
		zone.NewZone(zone.ZoneTypeCity, "Деревня травников", 3),
		zone.NewZone(zone.ZoneTypeCity, "Сад роз", 7),
		zone.NewZone(zone.ZoneTypeCity, "Крепость им. Розы", 13),
	)

	w.zones = zones
}
