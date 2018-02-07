package main

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
