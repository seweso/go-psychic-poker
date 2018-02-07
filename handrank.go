package main

import "fmt"

type handrank int

const (
	straightFlush handrank = 1 + iota
	fourOfAKind
	fullHouse
	flush
	straight
	threeOfAKind
	twoPairs
	onePair
	highestCard
)

var handranks = map[handrank]string{
	straightFlush: "straight-flush",
	fourOfAKind:   "four-of-a-kind",
	fullHouse:     "full-house",
	flush:         "flush",
	straight:      "straight",
	threeOfAKind:  "three-of-a-kind",
	twoPairs:      "two-pairs",
	onePair:       "one-pair",
	highestCard:   "highest-card",
}

func (f handrank) String() string {
	return fmt.Sprintf("%v", handranks[f])
}
