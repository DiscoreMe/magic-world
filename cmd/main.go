package main

import (
	"github.com/DiscoreMe/magic-world/entity"
	"github.com/DiscoreMe/magic-world/world"
	"time"
)

const worldWidth, worldHeight = 30, 15

func main() {
	w := world.NewWorld(worldWidth, worldHeight)
	w.CreateLand()
	hero := entity.NewHero("Nikita")
	w.AddEntity(0, 0, hero)

	for {
		w.Step()
		if err := w.ExportToJSON("test.world"); err != nil {
			panic(err)
		}
		time.Sleep(1 * time.Second)
	}
}
