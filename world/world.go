package world

import (
	"fmt"
	"github.com/DiscoreMe/magic-world/zone"
	"io/ioutil"
	"path"
	"sort"
	"time"
)

const maxDaysInYear = 300
const DurationDay = 5 * time.Second

var PathSaveWorld = path.Join("saves", "world.json")

type World struct {
	days  int
	years int
	zones []*zone.Zone

	debug bool
}

func NewWorld() *World {
	w := &World{
		days:  1,
		years: 1,
		debug: true,
	}
	w.initWorld()
	w.sortZones()

	return w
}

func (w *World) Days() int {
	return w.days
}

func (w *World) Years() int {
	return w.years
}

func debug(message string) {
	fmt.Println(message)
}

func (w *World) NextDay() {
	if w.days >= maxDaysInYear {
		w.days = 1
		w.years += 1
	}
	w.days++

	if w.debug {
		debug(fmt.Sprintf("Current time: %d day, %d year", w.days, w.years))
	}

	// todo log error
	w.Save()
}

func (w *World) Save() error {
	if err := w.exportToFile(PathSaveWorld); err != nil {
		return fmt.Errorf("export to file: %w", err)
	}

	return nil
}

func (w *World) Load(filepath string) error {
	b, err := ioutil.ReadFile(PathSaveWorld)
	if err != nil {
		return fmt.Errorf("ioutil readfile: %w", err)
	}
	data, err := w.importFromBytes(b)
	if err != nil {
		return fmt.Errorf("import from bytes: %w", err)
	}

	w.days = data.Days
	w.years = data.Years

	return nil
}

func (w *World) sortZones() {
	sort.Slice(w.zones, func(i, j int) bool {
		return w.zones[i].Course < w.zones[j].Course
	})
}
