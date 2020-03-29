package main

import (
	"encoding/json"
	"fmt"
	"github.com/DiscoreMe/magic-world/world"
	"io/ioutil"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "<div>")

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

	zones := make([][]world.ExportZone, exportData.Width)
	for i := range zones {
		points := make([]world.ExportZone, exportData.Height)
		zones[i] = points
	}

	for _, zone := range exportData.Zones {
		zones[zone.X][zone.Y] = zone
	}

	entities := make([][]world.ExportEntity, 0)

	for y := 0; y < exportData.Height; y++ {
		for x := 0; x < exportData.Width; x++ {
			var s string = " "
			switch zones[x][y].Type {
			case world.TypeZoneLand:
				s = fmt.Sprintf(`<img src="%s"/>`, picLand)
			}

			if zones[x][y].Entities != nil {
				s = fmt.Sprintf(`<img src="%s"/>`, picHero)
				entities = append(entities, zones[x][y].Entities)
			}

			fmt.Fprint(w, s)
		}
		fmt.Fprint(w, "<br>")
	}
	fmt.Fprintln(w, "<hr>")

	for _, z := range entities {
		for _, e := range z {
			fmt.Fprintf(w, "[ID:%d] %s [X:%d;Y:%d]<br>", e.ID, e.Name, e.X, e.Y)
		}
	}
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

const picLand = `data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAACAAAAAgCAIAAAD8GO2jAAAAAXNSR0IArs4c6QAAAARnQU1BAACxjwv8YQUAAAAJcEhZcwAADsMAAA7DAcdvqGQAAAA2SURBVEhL7c2hAQAwCMTAp6N0ni7LhjX4KFzORKZuv2w60zUOkAPkADlADpAD5AA5QA6QA5B8JnkBX8np5eUAAAAASUVORK5CYII=`
const picHero = `data:image/png;base64,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`
