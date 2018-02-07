package main

import (
	"reflect"
	"testing"
)

func TestGetAllPossibleHandsTest(t *testing.T) {

	in := struct {
		hand []card
		deck []card
	}{
		[]card{card{face1Ace, hearts}, card{face1Ace, hearts}, card{face1Ace, hearts}},
		[]card{card{face1Ace, spades}, card{face1Ace, spades}, card{face1Ace, spades}},
	}

	want := [][]card{
		[]card{card{face1Ace, hearts}, card{face1Ace, hearts}, card{face1Ace, hearts}},
		[]card{card{face1Ace, spades}, card{face1Ace, hearts}, card{face1Ace, hearts}},
		[]card{card{face1Ace, hearts}, card{face1Ace, spades}, card{face1Ace, hearts}},
		[]card{card{face1Ace, spades}, card{face1Ace, spades}, card{face1Ace, hearts}},
		[]card{card{face1Ace, hearts}, card{face1Ace, hearts}, card{face1Ace, spades}},
		[]card{card{face1Ace, spades}, card{face1Ace, hearts}, card{face1Ace, spades}},
		[]card{card{face1Ace, hearts}, card{face1Ace, spades}, card{face1Ace, spades}},
		[]card{card{face1Ace, spades}, card{face1Ace, spades}, card{face1Ace, spades}},
	}

	got := GetAllPossibleHands(in.hand, in.deck)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("GetAllPossibleHands(%q) == %q, want %q", in, got, want)
	}
}
