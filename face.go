package main

import "fmt"

type face int

const (
	faceAce face = 1 + iota
	faceKing
	faceQueen
	faceJack
	faceTen
	face9
	face8
	face7
	face6
	face5
	face4
	face3
	face2
	face1Ace
)

var faces = map[face]string{
	faceAce:   "Ace",
	faceKing:  "King",
	faceQueen: "Queen",
	faceJack:  "Jack",
	faceTen:   "Ten",
	face9:     "9",
	face8:     "8",
	face7:     "7",
	face6:     "6",
	face5:     "5",
	face4:     "4",
	face3:     "3",
	face2:     "2",
	face1Ace:  "1Ace",
}

func (f face) String() string {
	return fmt.Sprintf("%v", faces[f])
}

func convertToFace(b byte) face {
	for k := range faces {
		if k.String()[0] == b {
			return k
		}
	}
	panic(fmt.Sprintf("Unknown face found: %v", b))
}
