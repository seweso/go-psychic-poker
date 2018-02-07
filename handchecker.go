package main

import "math"

type rankFunc func([]card) bool

var functions = map[handrank]rankFunc{
	straightFlush: IsStraightFlush,
	fourOfAKind:   IsSequentialHighRules,
	fullHouse:     IsFullHouse,
	flush:         IsFlush,
	straight:      IsStraight,
	threeOfAKind:  IsThreeOfAKind,
	twoPairs:      IsTwoPairs,
	onePair:       IsOnePair,
	highestCard:   IsHighestCard,
}

func GetBestRank(hand, deck []card) handrank {
	return highestCard
}

func GetAllPossibleHands(hand, deck []card) (r [][]card) {

	if len(hand) != len(deck) {
		panic("Hand should have equal amount of cards as deck")
	}

	// Calculate total number of options
	totalNrOfOptions := uint(math.Pow(2, float64(len(hand))))

	// Brute force all options by replacing cards based on bit mask
	for i := uint(0); i < totalNrOfOptions; i++ {
		handCopy := hand
		deckPosition := 0

		for b := uint(0); b < totalNrOfOptions; b++ {

			if (i & (1 << b)) != 0 {
				handCopy[b] = deck[deckPosition]
				deckPosition = deckPosition + 1

				r = append(r, handCopy)
			}
		}
	}
	return
}

func GetBestRankForHand(hand []card) handrank {

	if len(hand) != 5 {
		panic("Hand should contain 5 cards")
	}

	for k, v := range functions {
		if v(hand) {
			return k
		}
	}

	return highestCard
}

func IsStraightFlush(hand []card) bool {
	return false
}

func IsSequentialHighRules(hand []card) bool {
	return false
}

func IsFullHouse(hand []card) bool {
	return false
}

func IsFlush(hand []card) bool {
	return false
}

func IsStraight(hand []card) bool {
	return false
}

func IsThreeOfAKind(hand []card) bool {
	return false
}

func IsTwoPairs(hand []card) bool {
	return false
}

func IsOnePair(hand []card) bool {
	return false
}

func IsHighestCard(hand []card) bool {
	return true
}
