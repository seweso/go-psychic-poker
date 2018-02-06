package main

import (
	"fmt"
)

type suit int

const (
	clubs suit = 1 + iota
	diamonds
	hearts
	spades
)

var suits = map[suit]string{
	clubs:    "Clubs",
	diamonds: "Diamonds",
	hearts:   "Hearts",
	spades:   "Spades",
}

func (s suit) String() string {
	return fmt.Sprintf("%v", suits[s])
}

func convertToSuit(b byte) suit {
	for k := range suits {
		if k.String()[0] == b {
			return k
		}
	}
	panic(fmt.Sprintf("Unknown suit found: %v", b))
}
