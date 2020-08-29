package world

import (
	"encoding/json"
	"github.com/DiscoreMe/magic-world/zone"
	"os"
)

type exportWorld struct {
	Days  int          `json:"days"`
	Years int          `json:"years"`
	Zones []*zone.Zone `json:"zones"`
}

func (w *World) exportToJSON() ([]byte, error) {
	var exportData = exportWorld{
		Days:  w.days,
		Years: w.years,
		Zones: w.zones,
	}

	return json.Marshal(exportData)
}

func (w *World) exportToFile(filename string) error {
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	b, err := w.exportToJSON()
	if err != nil {
		return err
	}

	_, err = f.Write(b)
	return err
}

func (w *World) importFromBytes(b []byte) (exportWorld, error) {
	var data exportWorld
	return data, json.Unmarshal(b, &data)
}
