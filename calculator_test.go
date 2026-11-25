package main

import "testing"

func TestKeySequences(t *testing.T) {
	for _, test := range []struct{ keys, want string }{
		{"", "0"},
		{"2+3*4=", "20"},
		{"8+*2=", "16"},
		{"9-12=/2=", "-1.5"},
		{"0.1+0.2=", "0.3"},
		{"2+3==", "5"},
		{"2+3=4", "4"},
		{"2+\b=1=", "3"},
		{"12.30\b", "12.3"},
		{"001..00", "1.00"},
		{"1/0=", "Error"},
		{"0/0=*3", "3"},
		{"7/0=\b", "0"},
		{"7/0=.4", "0.4"},
		{"4*5\x1b2=", "2"},
		{"123456789012345678901", "1234567890123456789"},
		{"1/3=", "0.333333333333333"},
		{"1/3=*3=", "1"},
	} {
		t.Run(test.keys, func(t *testing.T) {
			var calc Calculator
			for _, key := range test.keys {
				calc.Press(key)
			}
			if got := calc.Display(); got != test.want {
				t.Errorf("display = %q, want %q", got, test.want)
			}
		})
	}
}
