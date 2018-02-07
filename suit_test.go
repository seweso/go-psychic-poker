package main

import "testing"
func TestSuitConvert(t *testing.T) {
	cases := []struct {
		in   byte
		want suit
	}{
		{"S"[0], spades},
		{"D"[0], diamonds},
		{"H"[0], hearts},
		{"C"[0], clubs},
	}

	for _, c := range cases {
		got := convertToSuit(c.in)
		if got != c.want {
			t.Errorf("convertToSuit(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}