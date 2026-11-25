package main

import (
	"runtime"

	"github.com/egoist/mygo"
	"github.com/shangzongyu/Caculator/locales"
)

func (app *calculatorView) menu(locale string, redraw func()) *mygo.Menu {
	var items []*mygo.MenuItem
	if runtime.GOOS == "darwin" {
		// Keep MyGo's standard roles and their localized labels and shortcuts.
		items = []*mygo.MenuItem{
			{Role: mygo.RoleAppMenu},
			{Role: mygo.RoleFileMenu},
			{Role: mygo.RoleEditMenu},
			{Role: mygo.RoleViewMenu},
			{Role: mygo.RoleWindowMenu},
		}
	}
	items = append(items, app.appearanceMenu(locales.In(locale), redraw), languageMenu(locale, mygo.App.SetLocale))
	return mygo.NewMenu(items)
}

func languageMenu(locale string, selectLanguage func(string)) *mygo.MenuItem {
	// Use each language's own name so users can always recognize their choice.
	options := [...]struct{ tag, name string }{
		{"zh-Hans", "简体中文"},
		{"en", "English"},
	}
	items := make([]*mygo.MenuItem, 0, len(options))
	for _, option := range options {
		items = append(items, &mygo.MenuItem{
			Label:   option.name,
			Type:    mygo.MenuItemRadio,
			Checked: locale == option.tag,
			Click:   func(*mygo.MenuItem, *mygo.Window) { selectLanguage(option.tag) },
		})
	}
	return &mygo.MenuItem{Label: locales.In(locale).LanguageMenu(), Submenu: items}
}
