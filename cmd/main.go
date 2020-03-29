package main

import (
	"time"

	"github.com/DiscoreMe/magic-world/entity"
	"github.com/DiscoreMe/magic-world/world"
)

func main() {
	w := world.NewWorld(30, 15)
	w.AddEntity(0, 0, entity.NewHero("Nikita"))

	for {
		w.Step()

		if err := w.ExportToJSON("test.world"); err != nil {
			panic(err)
		}

		time.Sleep(time.Second)
	}
}
