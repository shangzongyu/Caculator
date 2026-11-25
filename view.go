package main

import (
	"strings"

	"github.com/egoist/mygo/ui"
)

var (
	paper    = ui.Hex("#efede5")
	ink      = ui.Hex("#1e3430")
	ivory    = ui.Hex("#fffef8")
	sage     = ui.Hex("#d1dfd5")
	amber    = ui.Hex("#ffbc70")
	keyOrder = []rune("\x1b\b789/456*123-.0=+")
	keyNames = map[rune]string{'\x1b': "清空", '\b': "退格", '+': "加", '-': "减", '*': "乘", '/': "除", '=': "等于", '.': "小数点"}
	keyFaces = map[rune]string{'\x1b': "清空", '\b': "⌫", '-': "−", '*': "×", '/': "÷"}
)

func keyFace(key rune) string {
	if face, ok := keyFaces[key]; ok {
		return face
	}
	return string(key)
}

func (calc *Calculator) view(c *ui.Context) {
	ui.Column(c).Fill().Padding(16).Gap(14).Background(paper).
		Focusable().AutoFocus().TextCaret(ui.Rect{}).InputMethod(false).HandleInput(calc.keyboard).Children(func() {
		ui.Column(c).Height(114).Padding(12, 16).Radius(12).Justify(ui.End).Background(ink).Children(func() {
			if calc.operator != 0 {
				ui.Text(c, formatResult(calc.accumulator)+" "+keyFace(calc.operator)).
					TextAlign(ui.End).FontSize(15).TextColor(sage)
			}
			result := calc.Display()
			size := min(float32(48), 450/float32(len(result)))
			ui.Text(c, result).SingleLine().TextAlign(ui.End).FontSize(size).TextColor(ivory).Role(ui.RoleStatus)
		})
		ui.Grid(c).Columns(4).GridRows(5).Grow(1).Gap(7).Children(func() {
			for _, key := range keyOrder {
				calc.button(c, key)
			}
		})
	})
}

func (calc *Calculator) button(c *ui.Context, key rune) {
	face := keyFace(key)
	button := ui.ButtonBase(c.Key(string(key))).Center().Radius(10).TextCaret(ui.Rect{}).InputMethod(false).HandleInput(calc.keyboard)
	name := keyNames[key]
	if name == "" {
		name = face
	}
	button.Label(name)

	color, size := ivory, float32(25)
	switch {
	case key == '\x1b':
		button.ColumnSpan(3)
		color, size = sage, 17
	case key == '\b' || strings.ContainsRune("+-*/", key):
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
		ui.Text(c, face).FontSize(size).TextColor(ink)
	})
	if button.Clicked() {
		calc.Press(key)
	}
}

func (calc *Calculator) keyboard(event ui.InputEvent) bool {
	switch event.Kind {
	case ui.InputText:
		// Use the character the keyboard layout produced, including keypad operators.
		handled := false
		for _, key := range event.Text {
			if strings.ContainsRune("0123456789.+-*/=", key) {
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
			calc.Press('\x1b')
		case ui.KeyBackspace:
			calc.Press('\b')
		default:
			return false
		}
		return true
	}
	return false
}
