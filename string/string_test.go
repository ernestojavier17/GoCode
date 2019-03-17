package string

import "testing"

func TestReverse(t *testing.T) {
	var tests = [] struct {
		s, expected string
	} {
		{"home", "emoh"},
	}
	for _, c := range tests {
		got := Reverse(c.s)
		if got != c.expected {
			t.Errorf("Reverse(%q) == %q, expected %q", c.s, c.expected, c.expected)
		}
	}
}
