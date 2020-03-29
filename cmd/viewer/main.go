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

	zones := make([][]world.ExportZone, exportData.Height)
	for i := range zones {
		points := make([]world.ExportZone, exportData.Width)
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

const picLand = `data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAEAAAABACAIAAAAlC+aJAAAAAXNSR0IArs4c6QAAAARnQU1BAACxjwv8YQUAAAAJcEhZcwAADsMAAA7DAcdvqGQAAABlSURBVGhD7c8xDcAwEMDATwGUP6WiSpdwOEXyLZ699vfOzZ7TazWgNaA1oDWgNaA1oDWgNaA1oDWgNaA1oDWgNaA1oDWgNaA1oDWgNaA1oDWgNaA1oDWgNaA1oDWgNaA1oDVgzfzY4QJWa5T6igAAAABJRU5ErkJggg==`
const picHero = `data:image/png;base64,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`
