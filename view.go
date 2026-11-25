package main

import (
	"strings"

	"github.com/egoist/mygo/ui"
)

var (
	paper = ui.Hex("#efede5")
	ink   = ui.Hex("#1e3430")
	ivory = ui.Hex("#fffef8")
	sage  = ui.Hex("#d1dfd5")
	amber = ui.Hex("#ffbc70")

	keypadRows = [...]string{
		"789/",
		"456*",
		"123-",
		".0=+",
	}
	keyNames = map[rune]string{clearKey: "清空", backspaceKey: "退格", '+': "加", '-': "减", '*': "乘", '/': "除", '=': "等于", '.': "小数点"}
	keyFaces = map[rune]string{clearKey: "清空", backspaceKey: "⌫", '-': "−", '*': "×", '/': "÷"}
)

func keyFace(key rune) string {
	if face, ok := keyFaces[key]; ok {
		return face
	}
	return string(key)
}

func (calc *Calculator) view(c *ui.Context) {
	width, height := c.Size()
	scale := min(width/windowWidth, height/windowHeight, 1.6)
	spacing := min(scale, 1)
	padding := 16 * spacing
	ui.Column(c).Fill().Padding(padding).Gap(14 * spacing).Background(paper).
		Focusable().AutoFocus().TextCaret(ui.Rect{}).InputMethod(false).HandleInput(calc.keyboard).Children(func() {
		calc.display(c, scale, padding)
		calc.keypad(c, scale, 7*spacing)
	})
}

func (calc *Calculator) display(c *ui.Context, scale, padding float32) {
	width, _ := c.Size()
	available := width - 4*padding // Outer margins and display padding on both sides.
	ui.Column(c).Height(114*scale).Padding(padding/2, padding).Radius(12).Justify(ui.End).Background(ink).Children(func() {
		if calc.operator != 0 {
			pending := formatResult(calc.accumulator) + " " + keyFace(calc.operator)
			displayLine(c, pending, 15*scale, available).TextColor(sage)
		}
		displayLine(c, calc.Display(), 48*scale, available).TextColor(ivory).Role(ui.RoleStatus)
	})
}

func displayLine(c *ui.Context, text string, size, available float32) ui.Element {
	textWidth, _ := c.MeasureText(0, ui.Span{Text: text, Size: size})
	if available > 0 && textWidth > available {
		size *= available / textWidth
	}
	return ui.Text(c, text).SingleLine().TextAlign(ui.End).FontSize(size)
}

func (calc *Calculator) keypad(c *ui.Context, scale, gap float32) {
	ui.Grid(c).Columns(4).GridRows(1 + len(keypadRows)).Grow(1).Gap(gap).Children(func() {
		calc.button(c, clearKey, scale).ColumnSpan(3)
		calc.button(c, backspaceKey, scale)
		for _, row := range keypadRows {
			for _, key := range row {
				calc.button(c, key, scale)
			}
		}
	})
}

func (calc *Calculator) button(c *ui.Context, key rune, scale float32) ui.Element {
	face := keyFace(key)
	button := ui.ButtonBase(c.Key(string(key))).Center().Radius(10).TextCaret(ui.Rect{}).InputMethod(false).HandleInput(calc.keyboard)
	name := keyNames[key]
	if name == "" {
		name = face
	}
	button.Label(name)

	color, size := ivory, float32(25)
	switch {
	case key == clearKey:
		color, size = sage, 17
	case key == backspaceKey || strings.ContainsRune(operatorKeys, key):
		color = sage
	case key == '=':
		color = amber
	}
	if button.Pressed() {
		color = color.Mix(ink, 0.18)
	} else if button.Hovered() {
		color = color.Mix(ink, 0.08)
	}
	button.Background(color).Border(1, ink.Mix(color, 0.82)).Children(func() {
		ui.Text(c, face).FontSize(size * scale).TextColor(ink)
	})
	if button.Clicked() {
		calc.Press(key)
	}
	return button
}

func (calc *Calculator) keyboard(event ui.InputEvent) bool {
	switch event.Kind {
	case ui.InputText:
		// Use the character the keyboard layout produced, including keypad operators.
		handled := false
		for _, key := range event.Text {
			if strings.ContainsRune("0123456789.="+operatorKeys, key) {
				calc.Press(key)
				handled = true
			}
		}
		return handled
	case ui.InputKeyDown:
		if event.Mods != 0 {
			return false
		}
		switch event.Key {
		case ui.KeyEnter:
			calc.Press('=')
		case ui.KeyEscape:
			calc.Press(clearKey)
		case ui.KeyBackspace:
			calc.Press(backspaceKey)
		default:
			return false
		}
		return true
	}
	return false
}
