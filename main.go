package main

import (
	"log"

	"github.com/egoist/mygo"
	"github.com/egoist/mygo/ui"
	"github.com/shangzongyu/Caculator/locales"
)

const (
	windowWidth     = 360
	windowHeight    = 492
	minWindowWidth  = 240
	minWindowHeight = 320
)

func main() {
	mygo.App.WhenReady(func() {
		app := calculatorView{appearance: savedAppearance()}
		window := mygo.NewWindow(mygo.WindowOptions{
			Title: locales.WindowTitle(), Width: windowWidth, Height: windowHeight,
			MinWidth: minWindowWidth, MinHeight: minWindowHeight, UseContentSize: true,
			Content: ui.View(app.view),
		})
		selectAppearance := func(index int) {
			if err := app.selectAppearance(index); err != nil {
				log.Printf("save appearance: %v", err)
			}
			window.Invalidate()
		}
		mygo.App.SetMenu(app.menu(mygo.App.Locale(), selectAppearance))
		mygo.App.OnLocaleChanged(func(locale string) {
			window.SetTitle(locales.In(locale).WindowTitle())
			mygo.App.SetMenu(app.menu(locale, selectAppearance))
		})
	})
	if err := mygo.App.Run(); err != nil {
		log.Fatal(err)
	}
}
