package main

import (
	"log"

	"github.com/egoist/mygo"
	"github.com/egoist/mygo/ui"
)

const (
	windowWidth     = 360
	windowHeight    = 492
	minWindowWidth  = 240
	minWindowHeight = 320
)

func main() {
	var app calculatorView
	mygo.App.WhenReady(func() {
		window := mygo.NewWindow(mygo.WindowOptions{
			Title: "计算器", Width: windowWidth, Height: windowHeight,
			MinWidth: minWindowWidth, MinHeight: minWindowHeight, UseContentSize: true,
			Content: ui.View(app.view),
		})
		appearance := app.appearanceMenu(window.Invalidate)
		if menu := mygo.App.Menu(); menu != nil {
			menu.Append(appearance)
		} else {
			mygo.App.SetMenu(mygo.NewMenu([]*mygo.MenuItem{appearance}))
		}
	})
	if err := mygo.App.Run(); err != nil {
		log.Fatal(err)
	}
}
