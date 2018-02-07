package main

import "testing"

func TestFaceString(t *testing.T) {
	cases := []struct {
		in   face
		want string
	}{
		{face2, "2"},
		{faceJack, "Jack"},
	}

	for _, c := range cases {
		got := c.in.String()
		if got != c.want {
			t.Errorf("face.String(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}

func TestFaceConvert(t *testing.T) {
	cases := []struct {
		in   byte
		want face
	}{
		{"A"[0], faceAce},
		{"T"[0], faceTen},
		{"8"[0], face8},
		{"Q"[0], faceQueen},
	}

	for _, c := range cases {
		got := convertToFace(c.in)
		if got != c.want {
			t.Errorf("convertToFace(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}
