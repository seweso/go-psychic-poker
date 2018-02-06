package main

import "testing"

func TestCard(t *testing.T) {
	cases := []struct {
		in   face
		want string
	}{
		{face1Ace, "1Ace"},
		{faceJack, "Jack"},
	}

	for _, c := range cases {
		got := c.in.String()
		if got != c.want {
			t.Errorf("face.String(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}
