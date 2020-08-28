package main

import (
	"fmt"
	"github.com/DiscoreMe/magic-world/world"
	"os"
	"os/exec"
	"runtime"
	"time"
)

const textGameInfo = `[ Magic world ]

На дворе %d год, день %d
`

const textAction = `
1. Обновить статистику
>>> 
`

func main() {
	os.MkdirAll("saves", os.ModePerm)

	w := world.NewWorld()

	var action int
	t := time.NewTicker(1 * time.Second)

	for range t.C {
		CallClear()
		fmt.Printf(textGameInfo, w.Years(), w.Days())
		fmt.Println()
		fmt.Print(textAction)
		fmt.Fscanln(os.Stdin, &action)

		switch action {
		default:
			err := w.Load(world.PathSaveWorld)
			if err != nil {
				fmt.Println(err)
			}
		}
	}
}

var clear map[string]func() //create a map for storing clear funcs

func init() {
	clear = make(map[string]func()) //Initialize it
	clear["linux"] = func() {
		cmd := exec.Command("clear") //Linux example, its tested
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
	clear["windows"] = func() {
		cmd := exec.Command("cmd", "/c", "cls") //Windows example, its tested
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

func CallClear() {
	value, ok := clear[runtime.GOOS] //runtime.GOOS -> linux, windows, darwin etc.
	if ok {                          //if we defined a clear func for that platform:
		value() //we execute it
	} else { //unsupported platform
		panic("Your platform is unsupported! I can't clear terminal screen :(")
	}
}
