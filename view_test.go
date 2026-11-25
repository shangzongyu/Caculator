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

func TestLongNumbersFitAcrossWindowShapes(t *testing.T) {
	var calc Calculator
	window := ui.NewTester(calc.view, windowWidth, windowHeight)
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
			clear.Y != backspace.Y || clear.X+clear.W > backspace.X || clear.W <= backspace.W {
			t.Fatalf("at %v, unexpected placement: hint %+v, result %+v, clear %+v, backspace %+v", size, hint, result, clear, backspace)
		}
	}
}
