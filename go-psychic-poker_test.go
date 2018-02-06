package main

import (
	"bufio"
	"bytes"
	"strings"
	"testing"
)

func TestProcessReader(t *testing.T) {
	in := strings.NewReader("TH JH QC QD QS QH KH AH 2S 6S\n2H 2S 3H 3S 3C 2D 3D 6C 9C TH")
	want := "Hand: TH JH QC QD QS Deck: QH KH AH 2S 6S Best hand: straight-flush\nHand: 2H 2S 3H 3S 3C Deck: 2D 3D 6C 9C TH Best hand: four-of-a-kind\n"

	bufIn := bufio.NewReader(in)
	var out bytes.Buffer
	bufOut := bufio.NewWriter(&out)

	ProcessReader(bufIn, bufOut)

	got := out.String()

	if got != want {
		t.Errorf("ProcessReader(%q) == %q, want %q", in, got, want)
	}
}

func TestProcessString(t *testing.T) {
	cases := []struct {
		in, want string
	}{
		{"TH JH QC QD QS QH KH AH 2S 6S", "Hand: TH JH QC QD QS Deck: QH KH AH 2S 6S Best hand: straight-flush"},
		{"2H 2S 3H 3S 3C 2D 3D 6C 9C TH", "Hand: 2H 2S 3H 3S 3C Deck: 2D 3D 6C 9C TH Best hand: four-of-a-kind"},
		{"2H 2S 3H 3S 3C 2D 9C 3D 6C TH", "Hand: 2H 2S 3H 3S 3C Deck: 2D 9C 3D 6C TH Best hand: full-house"},
		{"2H AD 5H AC 7H AH 6H 9H 4H 3C", "Hand: 2H AD 5H AC 7H Deck: AH 6H 9H 4H 3C Best hand: flush"},
		{"AC 2D 9C 3S KD 5S 4D KS AS 4C", "Hand: AC 2D 9C 3S KD Deck: 5S 4D KS AS 4C Best hand: straight"},
		{"KS AH 2H 3C 4H KC 2C TC 2D AS", "Hand: KS AH 2H 3C 4H Deck: KC 2C TC 2D AS Best hand: three-of-a-kind"},
		{"AH 2C 9S AD 3C QH KS JS JD KD", "Hand: AH 2C 9S AD 3C Deck: QH KS JS JD KD Best hand: two-pairs"},
		{"6C 9C 8C 2D 7C 2H TC 4C 9S AH", "Hand: 6C 9C 8C 2D 7C Deck: 2H TC 4C 9S AH Best hand: one-pair"},
		{"3D 5S 2H QD TD 6S KH 9H AD QH", "Hand: 3D 5S 2H QD TD Deck: 6S KH 9H AD QH Best hand: highest-card"},
	}
	for _, c := range cases {
		got := ProcessString(c.in)
		if got != c.want {
			t.Errorf("ProcessString(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}
