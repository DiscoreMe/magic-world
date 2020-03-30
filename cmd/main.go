package main

import (
	"github.com/brianvoe/gofakeit"
	"log"
	"os"
	"sync"
	"time"

	"github.com/DiscoreMe/magic-world/entity"
	"github.com/DiscoreMe/magic-world/server"
	"github.com/DiscoreMe/magic-world/world"
)

func loadConf() string {
	host, port := os.Getenv("MG_HOST"), os.Getenv("MG_PORT")
	if port == "" {
		port = "7777"
	}
	return host + ":" + port
}

func main() {
	w := world.NewWorld(30, 15)
	for i := 0; i < 3; i++ {
		w.AddEntity(0, 0, entity.NewHero(gofakeit.FirstName()))
	}

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			w.Step()
			time.Sleep(time.Second)
		}
	}()

	address := loadConf()

	serve := server.NewServer(w)
	if err := serve.Listen(address); err != nil {
		log.Fatal(err)
	}

	wg.Wait()
}
