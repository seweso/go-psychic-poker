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
