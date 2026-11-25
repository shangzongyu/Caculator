package main

import (
	"testing"

	"github.com/egoist/mygo"
	"github.com/egoist/mygo/ui"
	"github.com/shangzongyu/Caculator/locales"
)

func TestKeypadAndKeyboardShareState(t *testing.T) {
	var app calculatorView
	window := ui.NewTester(app.view, windowWidth, windowHeight)
	checkKeyboard := func() {
		t.Helper()
		if active, direct := window.TextInputActive(); !active || !direct {
			t.Fatalf("native character input active = %v, direct = %v", active, direct)
		}
	}
	checkKeyboard()
	window.TypeKey(0, ui.Key2, "2")
	window.TypeKey(ui.Shift, ui.KeyEqual, "+")
	click := func(name string) {
		t.Helper()
		if err := window.Click(name); err != nil {
			t.Fatal(err)
		}
	}
	click("3")
	checkKeyboard()
	window.TypeKey(ui.Shift, ui.Key8, "*")
	click("4")
	window.Key(0, ui.KeyEnter)
	if app.calc.Display() != "20" || !window.HasText("20") {
		t.Fatalf("after mixed input, display = %s, visible text = %q", app.calc.Display(), window.Texts())
	}
	window.Key(0, ui.KeyEscape)
	window.Type("12.3")
	window.Key(0, ui.KeyBackspace)
	click("加")
	click("4")
	click("等于")
	if app.calc.Display() != "16" || !window.HasText("16") {
		t.Fatalf("after click on equals, display = %s, visible text = %q", app.calc.Display(), window.Texts())
	}
}

func TestResizeDuringCalculation(t *testing.T) {
	var app calculatorView
	window := ui.NewTester(app.view, windowWidth, windowHeight)
	window.Type("12+")
	window.SetSize(720, 640)
	large, ok := window.Find("等于")
	if !ok {
		t.Fatal("missing equals key after enlarging window")
	}
	if err := window.Click("3"); err != nil {
		t.Fatal(err)
	}

	window.SetSize(minWindowWidth, minWindowHeight)
	small, ok := window.Find("等于")
	if !ok || small.W >= large.W || small.H >= large.H {
		t.Fatalf("equals key did not resize: large %+v, small %+v", large, small)
	}
	if err := window.Click("4"); err != nil {
		t.Fatal(err)
	}
	if err := window.Click("等于"); err != nil {
		t.Fatal(err)
	}
	if got := app.calc.Display(); got != "46" || !window.HasText("46") {
		t.Fatalf("after resizing, display = %q, want 46", got)
	}
}

func TestLongNumbersFitAcrossWindowShapes(t *testing.T) {
	var app calculatorView
	window := ui.NewTester(app.view, windowWidth, windowHeight)
	const entry = "9876543210987654321"
	window.Type("1234567890123456789+" + entry)
	const pending = "1234567890123456789 +"

	for _, size := range [][2]int{{240, 320}, {240, 720}, {720, 320}} {
		window.SetSize(size[0], size[1])
		hint, hintOK := window.Find(pending)
		result, resultOK := window.Find(entry)
		clear, clearOK := window.Find("清空")
		backspace, backspaceOK := window.Find("退格")

		if !hintOK || !resultOK || !clearOK || !backspaceOK {
			t.Fatalf("at %v, missing text or keys: %q", size, window.Texts())
		}
		for _, rect := range []ui.Rect{hint, result, clear, backspace} {
			if rect.W <= 0 || rect.H <= 0 || rect.X < 0 || rect.Y < 0 ||
				rect.X+rect.W > float32(size[0]) || rect.Y+rect.H > float32(size[1]) {
				t.Fatalf("at %v, element outside window: %+v", size, rect)
			}
		}
		const rounding = 1 // DIP
		if hint.Y+hint.H > result.Y+rounding || result.Y+result.H > clear.Y+rounding ||
			clear.Y != backspace.Y || clear.X+clear.W > backspace.X || clear.W+rounding < 3*backspace.W {
			t.Fatalf("at %v, unexpected placement: hint %+v, result %+v, clear %+v, backspace %+v", size, hint, result, clear, backspace)
		}
	}
}

func TestSwitchAppearanceDuringCalculation(t *testing.T) {
	var app calculatorView
	window := ui.NewTester(app.view, minWindowWidth, minWindowHeight)
	window.Type("12+")
	previous := window.Image().At(0, minWindowHeight-1)
	messages := locales.In("zh-Hans")
	menu := app.appearanceMenu(messages, window.Frame)
	if window.HasText(menu.Label) {
		t.Fatal("the appearance menu is still visible inside the calculator")
	}
	want := []string{"暖砂", "石墨", "瓷白"}
	if menu.Label != "外观" || len(menu.Submenu) != len(want) {
		t.Fatalf("unexpected appearance menu: %q with %d items", menu.Label, len(menu.Submenu))
	}
	for index, item := range menu.Submenu {
		if item.Label != want[index] || item.Type != mygo.MenuItemRadio || item.IsChecked() != (index == 0) {
			t.Fatalf("unexpected menu item %d: %q, %s, checked=%v", index, item.Label, item.Type, item.IsChecked())
		}
	}
	for offset := range menu.Submenu {
		index := (offset + 1) % len(menu.Submenu)
		item := menu.Submenu[index]
		name := want[index]
		item.Click(item, nil)
		if selected := appearances[app.appearance].name(messages); selected != name {
			t.Fatalf("appearance = %q, want %q", selected, name)
		}
		current := window.Image().At(0, minWindowHeight-1)
		if current == previous {
			t.Fatalf("background did not change after selecting %s", name)
		}
		previous = current
		if app.calc.Display() != "12" || !window.HasText("12 +") {
			t.Fatalf("pending calculation changed after selecting %s", name)
		}
	}

	window.Type("3")
	window.Key(0, ui.KeyEnter)
	if got := app.calc.Display(); got != "15" || !window.HasText("15") {
		t.Fatalf("after switching appearances, display = %q, want 15", got)
	}
}

func TestSwitchLanguageDuringCalculation(t *testing.T) {
	var app calculatorView
	window := ui.NewTester(app.view, minWindowWidth, minWindowHeight)
	window.SetLocale("zh-Hans")
	window.Type("12+3")

	menu := languageMenu("zh-Hans", window.SetLocale)
	if menu.Label != "语言" || !menu.Submenu[0].IsChecked() || menu.Submenu[1].IsChecked() {
		t.Fatal("language menu did not mark Simplified Chinese as selected")
	}
	english := menu.Submenu[1]
	english.Click(english, nil)
	if !window.HasText("Clear") || window.HasText("清空") || !window.HasText("12 +") {
		t.Fatalf("switching to English changed the calculation or left old labels: %q", window.Texts())
	}
	menu = languageMenu("en", window.SetLocale)
	if menu.Label != "Language" || menu.Submenu[0].IsChecked() || !menu.Submenu[1].IsChecked() {
		t.Fatal("language menu did not mark English as selected")
	}
	appearance := app.appearanceMenu(locales.In("en"), window.Frame)
	if appearance.Label != "Appearance" || appearance.Submenu[1].Label != "Graphite" {
		t.Fatal("appearance menu did not switch to English")
	}
	if err := window.Click("Equals"); err != nil {
		t.Fatal(err)
	}
	if got := app.calc.Display(); got != "15" {
		t.Fatalf("after switching languages, display = %q, want 15", got)
	}

	window.Type("1/0=")
	if !window.HasText("Error") {
		t.Fatalf("missing English error: %q", window.Texts())
	}
	chinese := menu.Submenu[0]
	chinese.Click(chinese, nil)
	if !window.HasText("错误") || window.HasText("Error") {
		t.Fatalf("error did not switch to Chinese: %q", window.Texts())
	}
}
