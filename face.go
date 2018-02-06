package main

import "fmt"

type face int

const (
	faceAce face = 0 + iota
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

var faces = [...]string{
	"Ace",
	"King",
	"Queen",
	"Jack",
	"Ten",
	"9",
	"8",
	"7",
	"6",
	"5",
	"4",
	"3",
	"2",
	"1Ace",
}

func (f face) String() string {
	return fmt.Sprintf("%v", faces[f])
}
