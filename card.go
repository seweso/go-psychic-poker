package main

import "fmt"

type card struct {
	f face
	s suit
}

func (c card) String() string {
	return fmt.Sprintf("%v of %v", c.f, c.s)
}

func (c card) ShortString() string {
	return fmt.Sprintf("%v%v", string(c.f.String()[0]), string(c.s.String()[0]))
}
