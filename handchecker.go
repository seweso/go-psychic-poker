package main

import (
	"math"
	"sort"
)

type rankFunc func(cards) bool

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

func GetBestRank(hand, deck cards) handrank {
	currentRank := highestCard

	// Brute force all possible hands (with this deck)
	for _, v := range GetAllPossibleHands(hand, deck) {
		rankForCurrentHand := GetBestRankForHand(v)

		if rankForCurrentHand == straightFlush {
			return straightFlush
		}

		// Found a beter rank?
		if rankForCurrentHand < currentRank {
			currentRank = rankForCurrentHand
		}
	}

	return currentRank
}

func GetAllPossibleHands(hand, deck cards) (r []cards) {

	if len(hand) != len(deck) {
		panic("Hand should have equal amount of cards as deck")
	}

	// Calculate total number of options
	totalNrOfOptions := uint(math.Pow(2, float64(len(hand))))

	// Brute force all options by replacing cards based on bit mask
	for i := uint(0); i < totalNrOfOptions; i++ {
		handCopy := hand.Copy()
		deckPosition := 0

		for b := uint(0); b < uint(len(hand)); b++ {
			if (i & (1 << b)) != 0 {
				handCopy[b] = deck[deckPosition]
				deckPosition = deckPosition + 1
			}
		}
		r = append(r, handCopy)
	}
	return
}

func GetBestRankForHand(hand cards) handrank {

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

func IsStraightFlush(hand cards) bool {
	if len(hand) != 5 {
		panic("Hand should contain 5 cards")
	}

	// Can't be a straight flush if all aren't the same suit
	if !IsSameSuit(hand) {
		return false
	}

	return IsSequentialHighRules(hand)
}

func IsSameSuit(hand cards) bool {
	suit := hand[0].s
	for _, k := range hand {
		if k.s != suit {
			return false
		}
	}
	return true
}

func IsSequentialHighRules(hand cards) bool {
	// Check whether cards are sequential, or cards are sequential when Ace is regarded as 1
	return IsSequential(hand) || IsSequential(hand.Select(AceAsOne))
}

func AceAsOne(c card) card {
	if c.f == faceAce {
		return card{face1Ace, c.s}
	}
	return c
}

func IsSequential(hand cards) bool {

	sort.Sort(byFace(hand))

	for i := 1; i < len(hand); i++ {
		if hand[i-1].f != hand[i].f-1 {
			return false
		}
	}
	return true
}

func IsFullHouse(hand cards) bool {
	// TODO Implement
	return false
}

func IsFlush(hand cards) bool {
	// TODO Implement
	return false
}

func IsStraight(hand cards) bool {
	// TODO Implement
	return false
}

func IsThreeOfAKind(hand cards) bool {
	// TODO Implement
	return false
}

func IsTwoPairs(hand cards) bool {
	// TODO Implement
	return false
}

func IsOnePair(hand cards) bool {
	// TODO Implement
	return false
}

func IsHighestCard(hand cards) bool {
	// TODO Implement
	return true
}
