package main

import (
	"github.com/egoist/mygo"
	"github.com/egoist/mygo/ui"
	"github.com/shangzongyu/Caculator/locales"
)

type keyColors struct {
	top, bottom, edge, text ui.Color
}

type appearance struct {
	name                        func(locales.Messages) string
	chassisTop, chassisBottom   ui.Color
	screenTop, screenBottom     ui.Color
	screenEdge, screenText      ui.Color
	screenMuted, screenError    ui.Color
	focus                       ui.Color
	highlight, keyShade, shadow ui.Color
	digits, actions, equals     keyColors
}

// Warm sand is first so a new window starts with the original appearance.
var appearances = [...]appearance{
	{
		name:       locales.Messages.AppearanceWarmSand,
		chassisTop: ui.Hex("#f4efe6"), chassisBottom: ui.Hex("#e4dbcc"),
		screenTop: ui.Hex("#252a26"), screenBottom: ui.Hex("#373e34"),
		screenEdge: ui.Hex("#20251f"), screenText: ui.Hex("#f3f5df"),
		screenMuted: ui.Hex("#c2c9b1"), screenError: ui.Hex("#ffd0b7"),
		focus: ui.Hex("#8d3f29"), highlight: ui.Hex("#ffffff77"),
		keyShade: ui.Hex("#35261b33"), shadow: ui.Hex("#4d3d2b2e"),
		digits:  newKeyColors("#fffdf7", "#eee7da", "#cbbfae", "#352e25"),
		actions: newKeyColors("#e9e1d2", "#d9cbbb", "#c0ae9a", "#43362a"),
		equals:  newKeyColors("#ac5035", "#8c3b26", "#763322", "#fff6e9"),
	},
	{
		name:       locales.Messages.AppearanceGraphite,
		chassisTop: ui.Hex("#30343a"), chassisBottom: ui.Hex("#1b1e22"),
		screenTop: ui.Hex("#101820"), screenBottom: ui.Hex("#1b2b34"),
		screenEdge: ui.Hex("#0a1117"), screenText: ui.Hex("#d6f3fa"),
		screenMuted: ui.Hex("#9abcc4"), screenError: ui.Hex("#ffbfb5"),
		focus: ui.Hex("#9fe4f1"), highlight: ui.Hex("#d3e4f142"),
		keyShade: ui.Hex("#10141c72"), shadow: ui.Hex("#090d125f"),
		digits:  newKeyColors("#454a52", "#30343b", "#24282e", "#f0f1f4"),
		actions: newKeyColors("#575e69", "#3e444e", "#30353e", "#f4f6fa"),
		equals:  newKeyColors("#9bdee7", "#68b6c6", "#4e93a4", "#0f2930"),
	},
	{
		name:       locales.Messages.AppearancePorcelain,
		chassisTop: ui.Hex("#f8fbff"), chassisBottom: ui.Hex("#dce7f2"),
		screenTop: ui.Hex("#dbe7f2"), screenBottom: ui.Hex("#edf5fc"),
		screenEdge: ui.Hex("#afc3d9"), screenText: ui.Hex("#132c4a"),
		screenMuted: ui.Hex("#3c5772"), screenError: ui.Hex("#9b2938"),
		focus: ui.Hex("#2158a1"), highlight: ui.Hex("#ffffff7a"),
		keyShade: ui.Hex("#162e5633"), shadow: ui.Hex("#243c6029"),
		digits:  newKeyColors("#ffffff", "#f3f8fe", "#c2d2e2", "#1d3855"),
		actions: newKeyColors("#e3edf8", "#d0dfef", "#afc5db", "#203e61"),
		equals:  newKeyColors("#326bb9", "#25518f", "#1a4078", "#fcfdff"),
	},
}

func newKeyColors(top, bottom, edge, text string) keyColors {
	return keyColors{top: ui.Hex(top), bottom: ui.Hex(bottom), edge: ui.Hex(edge), text: ui.Hex(text)}
}

func (app *calculatorView) appearanceMenu(messages locales.Messages, redraw func()) *mygo.MenuItem {
	items := make([]*mygo.MenuItem, 0, len(appearances))
	for index := range appearances {
		items = append(items, &mygo.MenuItem{
			Label:   appearances[index].name(messages),
			Type:    mygo.MenuItemRadio,
			Checked: index == app.appearance,
			Click: func(*mygo.MenuItem, *mygo.Window) {
				app.appearance = index
				redraw()
			},
		})
	}
	return &mygo.MenuItem{Label: messages.AppearanceMenu(), Submenu: items}
}
