package main

import (
	"flag"
	"fmt"
	"github.com/nsf/termbox-go"
)

func MaybePanic(err error) {
	if err != nil {
		panic(err)
	}
}

var (
	Quit   chan struct{}
	redraw chan struct{}
)

func main() {
	flag.Parse()
	if len(flag.Args()) == 0 {
		fmt.Println("sux: no commands given")
		fmt.Println("Usage sux [command ...]")
		return
	}

	DefaultMode = InputMode
	CurrentMode = DefaultMode

	Quit = make(chan struct{})
	redraw = make(chan struct{})

	MaybePanic(termbox.Init())

	termbox.SetInputMode(termbox.InputEsc)
	termbox.SetOutputMode(termbox.Output256)

	defer termbox.Close()
	defer EndPanes()

	go InputLoop()
	go OutputLoop()

	MaybePanic(RunPanes())

	<-Quit
}
