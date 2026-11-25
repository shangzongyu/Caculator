package main

import (
	"strings"
	"time"

	"github.com/egoist/mygo/ui"
	"github.com/shangzongyu/Caculator/locales"
)

const (
	maxContentScale  = 1.6
	maxContentWidth  = 540
	maxContentHeight = 660
	windowPadding    = 18
	sectionGap       = 14
	keyGap           = 9
	displayHeight    = 120
	displayRadius    = 17
	pendingFontSize  = 14
	resultFontSize   = 52
	displayFont      = "monospace"
	keypadColumns    = 4
	keyRadius        = 11
	keyTravel        = 2
	keyPressTime     = 120 * time.Millisecond
	topToBottom      = 180
)

type calculatorView struct {
	calc       Calculator
	appearance int
}

var (
	keypadRows = [...]string{
		"789/",
		"456*",
		"123-",
		".0=+",
	}
	keyNames = map[rune]func(locales.Messages) string{
		clearKey: locales.Messages.KeyClear, backspaceKey: locales.Messages.KeyBackspace,
		'+': locales.Messages.KeyAdd, '-': locales.Messages.KeySubtract,
		'*': locales.Messages.KeyMultiply, '/': locales.Messages.KeyDivide,
		'=': locales.Messages.KeyEquals, '.': locales.Messages.KeyDecimalPoint,
	}
	keyFaces = map[rune]string{backspaceKey: "⌫", '-': "−", '*': "×", '/': "÷"}
)

func keyFace(key rune) string {
	if face, ok := keyFaces[key]; ok {
		return face
	}
	return string(key)
}

func (app *calculatorView) view(c *ui.Context) {
	style := &appearances[app.appearance]
	width, height := c.Size()
	width, height = min(width, maxContentWidth), min(height, maxContentHeight)
	scale := min(width/windowWidth, height/windowHeight, maxContentScale)
	// Shrink gaps in a small window to leave more room for the keypad.
	spacing := min(scale, 1)
	padding := windowPadding * spacing
	ui.Column(c).Fill().Center().Gradient(style.chassisTop, style.chassisBottom, topToBottom).
		Focusable().AutoFocus().TextCaret(ui.Rect{}).InputMethod(false).HandleInput(app.calc.keyboard).Children(func() {
		ui.Column(c).Size(width, height).Padding(padding).Gap(sectionGap * spacing).Children(func() {
			app.display(c, style, width-2*padding, scale, padding)
			app.keypad(c, style, scale, keyGap*spacing)
		})
	})
}

func (app *calculatorView) display(c *ui.Context, style *appearance, width, scale, padding float32) {
	const borderWidth = 1
	available := width - 2*(padding+borderWidth)
	calc := &app.calc
	ui.Column(c).Height(displayHeight*scale).Padding(padding/2, padding).Radius(displayRadius*scale).
		Justify(ui.End).Gradient(style.screenTop, style.screenBottom, topToBottom).Border(borderWidth, style.screenEdge).
		InsetShadow(0, 2*scale, 6*scale, 0, style.screenEdge).Children(func() {
		if calc.operator != 0 {
			pending := formatResult(calc.accumulator) + " " + keyFace(calc.operator)
			displayLine(c, pending, pendingFontSize*scale, available).TextColor(style.screenMuted)
		}
		result, color := calc.Display(), style.screenText
		if calc.state == showingError {
			result = locales.In(c.Locale()).CalculationError()
			color = style.screenError
		}
		displayLine(c, result, resultFontSize*scale, available).TextColor(color).Role(ui.RoleStatus)
	})
}

func displayLine(c *ui.Context, text string, size, available float32) ui.Element {
	textWidth, _ := c.MeasureText(0, ui.Span{Text: text, Font: displayFont, Size: size})
	if available > 0 && textWidth > available {
		size *= available / textWidth
	}
	return ui.Text(c, text).SingleLine().TextAlign(ui.End).Font(displayFont).FontSize(size)
}

func (app *calculatorView) keypad(c *ui.Context, style *appearance, scale, gap float32) {
	ui.Grid(c).Columns(keypadColumns).GridRows(1 + len(keypadRows)).Grow(1).Gap(gap).Children(func() {
		app.button(c, style, clearKey, scale).ColumnSpan(keypadColumns - 1)
		app.button(c, style, backspaceKey, scale)
		for _, row := range keypadRows {
			for _, key := range row {
				app.button(c, style, key, scale)
			}
		}
	})
}

func (app *calculatorView) button(c *ui.Context, style *appearance, key rune, scale float32) ui.Element {
	face := keyFace(key)
	button := ui.ButtonBase(c.Key(string(key))).Center().Radius(keyRadius * scale).FocusRingColor(style.focus).
		TextCaret(ui.Rect{}).InputMethod(false).HandleInput(app.calc.keyboard)
	name := face
	if translate := keyNames[key]; translate != nil {
		name = translate(locales.In(c.Locale()))
	}
	button.Label(name)

	colors, size := style.digits, float32(25)
	switch {
	case key == clearKey:
		colors, size = style.actions, 16
		face = name
	case key == backspaceKey || strings.ContainsRune(operatorKeys, key):
		colors = style.actions
	case key == '=':
		colors = style.equals
	}
	drawKeycap(c, button, style, colors, face, size, scale)
	if button.Clicked() {
		app.calc.Press(key)
	}
	return button
}

func drawKeycap(c *ui.Context, button ui.Element, style *appearance, colors keyColors, face string, size, scale float32) {
	var pressed float32
	if button.Pressed() {
		pressed = 1
	}
	depth := button.Animate("press", pressed, keyPressTime)
	rest := 1 - depth
	top, bottom := colors.top, colors.bottom
	if button.Hovered() {
		white := ui.Hex("#ffffff")
		top, bottom = top.Mix(white, 0.08), bottom.Mix(white, 0.08)
	}
	// Move the cap while the button's pointer target stays in its grid cell.
	button.Children(func() {
		ui.Box(c).Fill().Center().Radius(keyRadius*scale).Translate(0, keyTravel*depth*scale).
			Gradient(top.Mix(bottom, depth), bottom.Mix(top, depth), topToBottom).
			Border(1, colors.edge).
			InsetShadow(0, scale, 0, 0, style.highlight.Mix(style.keyShade, depth)).
			InsetShadow(0, -scale, 0, 0, style.keyShade.Mix(style.highlight, depth)).
			Shadow(0, (1+rest)*scale, (1+2*rest)*scale, 0, style.shadow).
			Children(func() {
				ui.Text(c, face).FontSize(size * scale).FontWeight(550).TextColor(colors.text)
			})
	})
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
