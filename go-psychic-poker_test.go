package main

import (
	"testing"
)

func TestProcessString(t *testing.T) {
	cases := []struct {
		in, want string
	}{
		{"TH JH QC QD QS QH KH AH 2S 6S", "Hand: TH JH QC QD QS Deck: QH KH AH 2S 6S Best hand: straight-flush"},
		{"2H 2S 3H 3S 3C 2D 3D 6C 9C TH", "Hand: 2H 2S 3H 3S 3C Deck: 2D 3D 6C 9C TH Best hand: four-of-a-kind"},
	}
	for _, c := range cases {
		got := ProcessString(c.in)
		if got != c.want {
			t.Errorf("Reverse(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}
