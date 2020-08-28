package main

import (
	"fmt"
	"github.com/DiscoreMe/magic-world/world"
	"time"
)

func main() {
	w := world.NewWorld()
	if err := w.Load(world.PathSaveWorld); err != nil {
		fmt.Println(err)
	}

	t := time.NewTicker(world.DurationDay)
	for range t.C {
		w.NextDay()
	}
}
