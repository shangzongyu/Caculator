package main

import (
	"math"
	"strconv"
	"strings"
)

type entryState int

const (
	entering entryState = iota
	awaitingOperand
	showingResult
	showingError
)

// Calculator evaluates operations in the order their keys are pressed.
type Calculator struct {
	entry       string
	state       entryState
	operator    rune
	accumulator float64
}

func (c *Calculator) Display() string {
	if c.entry == "" {
		return "0"
	}
	return c.entry
}

// Press accepts a keypad character, backspace (\b), or clear (Escape).
func (c *Calculator) Press(key rune) {
	switch {
	case key == '\x1b' || key == '\b' && c.state == showingError:
		*c = Calculator{}
	case key == '\b':
		if c.state == entering && len(c.entry) > 0 {
			c.entry = c.entry[:len(c.entry)-1]
		}
	case key >= '0' && key <= '9' || key == '.':
		if c.state != entering {
			c.entry = ""
			c.state = entering
		}
		if key == '.' {
			if strings.Contains(c.entry, ".") {
				return
			}
			c.entry = c.Display()
		} else if c.entry == "0" {
			c.entry = ""
		}
		if len(c.entry) < 19 {
			c.entry += string(key)
		}
	case strings.ContainsRune("+-*/", key):
		if c.state == showingError || c.state == entering && !c.commit() {
			return
		}
		c.operator = key
		c.state = awaitingOperand
	case key == '=':
		if c.state == entering && c.operator != 0 && c.commit() {
			c.operator = 0
			c.state = showingResult
		}
	}
}

func (c *Calculator) commit() bool {
	value, err := strconv.ParseFloat(c.Display(), 64)
	switch c.operator {
	case '+':
		value = c.accumulator + value
	case '-':
		value = c.accumulator - value
	case '*':
		value = c.accumulator * value
	case '/':
		value = c.accumulator / value
	}
	if err != nil || math.IsInf(value, 0) || math.IsNaN(value) {
		*c = Calculator{entry: "Error", state: showingError}
		return false
	}
	c.accumulator = value
	if c.operator != 0 {
		c.entry = formatResult(value)
	}
	return true
}

func formatResult(value float64) string {
	if value == 0 {
		return "0"
	}
	return strconv.FormatFloat(value, 'g', 15, 64)
}
