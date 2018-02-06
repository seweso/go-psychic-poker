package main

import "testing"

func TestFaceString(t *testing.T) {
	cases := []struct {
		in   face
		want string
	}{
		{face1Ace, "1Ace"},
		{faceJack, "Jack"},
	}

	for _, c := range cases {
		got := c.in.String()
		if got != c.want {
			t.Errorf("face.String(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}

func TestCardString(t *testing.T) {
	cases := []struct {
		in   card
		want string
	}{
		{card{faceAce, diamonds}, "Ace of Diamonds"},
		{card{faceQueen, spades}, "Queen of Spades"},
	}

	for _, c := range cases {
		got := c.in.String()
		if got != c.want {
			t.Errorf("card.String(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}

func TestCardStringShort(t *testing.T) {
	cases := []struct {
		in   card
		want string
	}{
		{card{faceAce, diamonds}, "AD"},
		{card{faceQueen, spades}, "QS"},
	}

	for _, c := range cases {
		got := c.in.ShortString()
		if got != c.want {
			t.Errorf("card.String(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}

func TestCardConvert(t *testing.T) {
	cases := []struct {
		in   string
		want card
	}{
		{"AD", card{faceAce, diamonds}},
		{"TC", card{faceTen, clubs}},
		{"8S", card{face8, spades}},
		{"1H", card{face1Ace, hearts}},
	}

	for _, c := range cases {
		got := convertToOneCard(c.in)
		if got != c.want {
			t.Errorf("convertToFace(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}

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

func TestFaceConvert(t *testing.T) {
	cases := []struct {
		in   byte
		want face
	}{
		{"1"[0], face1Ace},
		{"T"[0], faceTen},
		{"8"[0], face8},
		{"Q"[0], faceQueen},
	}

	for _, c := range cases {
		got := convertToFace(c.in)
		if got != c.want {
			t.Errorf("convertToFace(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}
