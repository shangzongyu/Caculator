package main

import (
	"testing"

	"github.com/egoist/mygo/ui"
)

func TestKeypadAndKeyboardShareState(t *testing.T) {
	var calc Calculator
	window := ui.NewTester(calc.view, windowWidth, windowHeight)
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
	if calc.Display() != "20" || !window.HasText("20") {
		t.Fatalf("after mixed input, display = %s, visible text = %q", calc.Display(), window.Texts())
	}
	window.Key(0, ui.KeyEscape)
	window.Type("12.3")
	window.Key(0, ui.KeyBackspace)
	click("加")
	click("4")
	click("等于")
	if calc.Display() != "16" || !window.HasText("16") {
		t.Fatalf("after click on equals, display = %s, visible text = %q", calc.Display(), window.Texts())
	}
}

func TestResizeDuringCalculation(t *testing.T) {
	var calc Calculator
	window := ui.NewTester(calc.view, windowWidth, windowHeight)
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
	if got := calc.Display(); got != "46" || !window.HasText("46") {
		t.Fatalf("after resizing, display = %q, want 46", got)
	}
}
