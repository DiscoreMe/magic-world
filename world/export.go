package world

import (
	"encoding/json"
	"os"
)

type ExportWorld struct {
	Days   int64        `json:"days"`
	Width  int          `json:"width"`
	Height int          `json:"height"`
	Cells  []ExportCell `json:"cells,omitempty"`
}

type ExportCell struct {
	X        int            `json:"x"`
	Y        int            `json:"y"`
	Type     int            `json:"type"`
	Meta     string         `json:"meta,omitempty"`
	Entities []ExportEntity `json:"entities,omitempty"`
}

type ExportEntity struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
	X    int    `json:"x"`
	Y    int    `json:"y"`
}

func (w *World) ExportToJSON(filename string) error {
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	exportWorld := ExportWorld{
		Width:  w.Width(),
		Height: w.Height(),
		Days:   w.days,
	}

	for y := 0; y < exportWorld.Height; y++ {
		for x := 0; x < exportWorld.Width; x++ {
			cell, _ := w.zone.Cell(x, y)
			exportCell := ExportCell{
				X:    x,
				Y:    y,
				Type: cell.Type,
				Meta: cell.Meta,
			}
			for _, ent := range cell.entities {
				exportCell.Entities = append(exportCell.Entities, ExportEntity{
					ID:   ent.ID(),
					Name: ent.Name(),
					X:    ent.X(),
					Y:    ent.Y(),
				})
			}

			exportWorld.Cells = append(exportWorld.Cells, exportCell)
		}
	}

	data, err := json.Marshal(exportWorld)
	if err != nil {
		return err
	}

	_, err = f.Write(data)
	return err
}
