package main

import (
	"reflect"
	"testing"
)

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
		{"JH", card{faceJack, hearts}},
	}

	for _, c := range cases {
		got := convertToOneCard(c.in)
		if got != c.want {
			t.Errorf("convertToOneCard(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}

func TestCardsConvert(t *testing.T) {
	cases := []struct {
		in   string
		want []card
	}{
		{"AD", []card{card{faceAce, diamonds}}},
		{"TC 5D ", []card{card{faceTen, clubs}, card{face5, diamonds}}},
		{"8S KC QH", []card{card{face8, spades}, card{faceKing, clubs}, card{faceQueen, hearts}}},
		{" AD 2H 3H JC", []card{card{face1Ace, diamonds}, card{face2, hearts}, card{face3, hearts}, card{faceJack, clubs}}},
	}

	for _, c := range cases {
		got := convertToCards(c.in)
		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("convertToOneCard(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}
