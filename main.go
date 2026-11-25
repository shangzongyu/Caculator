package main

import (
	"log"

	"github.com/egoist/mygo"
	"github.com/egoist/mygo/ui"
)

const (
	windowWidth  = 360
	windowHeight = 492
)

func main() {
	var calc Calculator
	mygo.App.WhenReady(func() {
		mygo.NewWindow(mygo.WindowOptions{
			Title: "计算器", Width: windowWidth, Height: windowHeight,
			UseContentSize: true, DisableResize: true, DisableMaximize: true,
			Content: ui.View(calc.view),
		})
	})
	if err := mygo.App.Run(); err != nil {
		log.Fatal(err)
	}
}
