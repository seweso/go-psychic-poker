package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	ProcessReader(reader, writer)
}

// ProcessReader data from reader and write to writer
func ProcessReader(reader *bufio.Reader, writer *bufio.Writer) {
	defer writer.Flush()

	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		cards := scanner.Text()
		if cards == "" {
			return
		}
		result := ProcessString(cards)
		fmt.Fprintln(writer, result)
	}
	if err := scanner.Err(); err != nil {
		panic("error reading input")
	}
}

// ProcessString with cards
func ProcessString(cardString string) string {
	cards := convertToCards(cardString)

	if len(cards) != 10 {
		panic(fmt.Sprintf("Card count incorrect: %v", len(cards)))
	}

	hand := cards[0:5]
	deck := cards[5:5]
	bestHand := GetBestRank(hand, deck)

	return fmt.Sprintf("Hand: %v Deck: %v Best hand: %b", hand, deck, bestHand)
}
