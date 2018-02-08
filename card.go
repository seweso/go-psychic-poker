package main

import (
	"bufio"
	"fmt"
	"strings"
)

type card struct {
	f face
	s suit
}

type cards []card

type convert func(card) card

// Convert one card as string to card struct
func convertToOneCard(s string) card {
	face := convertToFace(s[0])
	suit := convertToSuit(s[1])
	return card{face, suit}
}

func convertToCards(s string) (r []card) {
	scanner := bufio.NewScanner(strings.NewReader(s))
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		r = append(r, convertToOneCard(scanner.Text()))
	}
	return
}

func (c card) String() string {
	return fmt.Sprintf("%v of %v", c.f, c.s)
}

func (c card) ShortString() string {
	return fmt.Sprintf("%v%v", string(c.f.String()[0]), string(c.s.String()[0]))
}

func (c cards) Copy() cards {
	return append(cards(nil), c...)
}

func (c cards) Select(f convert) cards {
	n := make(cards, len(c))

	// Use converted function to convert all items
	for i := 0; i < len(c); i++ {
		n[i] = f(c[i])
	}
	return n
}

type byFace cards

func (a byFace) Len() int           { return len(a) }
func (a byFace) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a byFace) Less(i, j int) bool { return a[i].f < a[j].f }
