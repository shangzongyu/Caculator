package main

import (
	"math/big"
	"strings"
)

const (
	clearKey     = '\x1b' // ASCII Escape
	backspaceKey = '\b'
	operatorKeys = "+-*/"

	maxEntryLength          = 19 // Includes the decimal point.
	maxPlainResultLength    = 24
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
	accumulator *big.Rat
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
	value, ok := new(big.Rat).SetString(c.Display())
	if !ok || (c.operator == '/' && value.Sign() == 0) {
		*c = Calculator{entry: "Error", state: showingError}
		return false
	}
	switch c.operator {
	case '+':
		value.Add(c.accumulator, value)
	case '-':
		value.Sub(c.accumulator, value)
	case '*':
		value.Mul(c.accumulator, value)
	case '/':
		value.Quo(c.accumulator, value)
	}
	c.accumulator = value // Chained operations use the value before display rounding.
	if c.operator != 0 {
		c.entry = formatResult(value)
	}
	return true
}

func formatResult(value *big.Rat) string {
	fractionDigits, exact := value.FloatPrec()
	if exact && fractionDigits < maxPlainResultLength {
		text := value.FloatString(fractionDigits)
		if len(text) <= maxPlainResultLength {
			return text
		}
	}
	// Keep repeating decimals and long results readable on the display.
	return new(big.Float).SetRat(value).Text('g', resultSignificantDigits)
}
