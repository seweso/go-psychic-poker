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
		convertToCards("1H 2H 3H"),
		convertToCards("1S 2S 3S"),
	}

	want := [][]card{
		convertToCards("1H 2H 3H"),
		convertToCards("1S 2H 3H"),
		convertToCards("1H 1S 3H"),
		convertToCards("1S 2S 3H"),
		convertToCards("1H 2H 1S"),
		convertToCards("1S 2H 2S"),
		convertToCards("1H 1S 2S"),
		convertToCards("1S 2S 3S"),
	}

	got := GetAllPossibleHands(in.hand, in.deck)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("GetAllPossibleHands(%q) == %q, want %q", in, got, want)
	}
}

func TestIsStraightFlush(t *testing.T) {
	isHand := [][]card{
		convertToCards("AH KH QH JH TH"),
		convertToCards("KH QH JH TH 9H"),
		convertToCards("QH JH TH 9H 8H"),
		convertToCards("8H 7H 6H 5H 4H"),
		convertToCards("6H 5H 2H 3H 4H"),
		convertToCards("AH 5H 4H 3H 2H"),
	}

	notHand := [][]card{
		convertToCards("AH JH TH 9H 8H"),
		convertToCards("QH JH TS 9H 8H"),
	}

	want := true
	for _, in := range isHand {
		got := IsStraightFlush(in)
		if got != want {
			t.Errorf("IsStraightFlush(%q) == %v, want %v", in, got, want)
		}
	}

	want = false
	for _, in := range notHand {
		got := IsStraightFlush(in)
		if got != want {
			t.Errorf("IsStraightFlush(%q) == %v, want %v", in, got, want)
		}
	}
}
