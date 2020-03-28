package main

import (
	"encoding/json"
	"github.com/DiscoreMe/magic-world/world"
	"io/ioutil"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	bytes, err := ioutil.ReadFile(r.URL.Path[1:])
	if err != nil {
		_, _ = w.Write([]byte(err.Error()))
		return
	}

	var exportData world.ExportData
	if err := json.Unmarshal(bytes, &exportData); err != nil {
		_, _ = w.Write([]byte(err.Error()))
		return
	}

	zones := make([][]world.ExportZone, exportData.Height)
	for i := range zones {
		points := make([]world.ExportZone, exportData.Width)
		zones[i] = points
	}

	for _, zone := range exportData.Zones {
		zones[zone.X][zone.Y] = zone
	}

	for y := 0; y < exportData.Height; y++ {
		for x := 0; x < exportData.Width; x++ {
			var s string = " "
			switch zones[x][y].Type {
			case world.TypeZoneLand:
				s = "🌳"
			}

			if zones[x][y].Entities != nil {
				s = "🙆🏻"
			}

			_, _ = w.Write([]byte(s))
		}
		_, _ = w.Write([]byte("\n"))
	}
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
