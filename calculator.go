package main

import (
	"math"
	"strconv"
	"strings"
)

const (
	clearKey     = '\x1b' // ASCII Escape
	backspaceKey = '\b'
	operatorKeys = "+-*/"

	maxEntryLength          = 19 // Includes the decimal point.
	resultSignificantDigits = 15
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

// Press enters one digit or applies one keypad action.
func (c *Calculator) Press(key rune) {
	switch {
	case key == clearKey || (key == backspaceKey && c.state == showingError):
		*c = Calculator{}
	case key == backspaceKey:
		if c.state == entering && len(c.entry) > 0 {
			// Editable input contains only ASCII digits and the decimal point.
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
		if len(c.entry) < maxEntryLength {
			c.entry += string(key)
		}
	case strings.ContainsRune(operatorKeys, key):
		// While waiting for the next operand, a second operator replaces the pending one.
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

// commit records the first operand or applies the pending operation to the current entry.
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
	c.accumulator = value // Chained operations use the value before display rounding.
	if c.operator != 0 {
		c.entry = formatResult(value)
	}
	return true
}

func formatResult(value float64) string {
	if value == 0 { // Display both positive and negative zero as 0.
		return "0"
	}
	return strconv.FormatFloat(value, 'g', resultSignificantDigits, 64)
}
