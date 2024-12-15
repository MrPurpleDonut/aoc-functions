package test

import (
	"testing"

	aoc "github.com/MrPurpleDonut/aoc-functions"
)

func TestMakeMatrix(t *testing.T) {
	t.Run("error's on no file", func(t *testing.T) {
		_, err := aoc.MakeMatrix("not a file")
		if err == nil {
			t.Errorf("Expected an err, but got nil")
		}
	})

	t.Run("returns a [][]rune on success", func(t *testing.T) {
		data, err := aoc.MakeMatrix("square.txt")
		if err != nil {
			t.Errorf("expected nil err, but got err")
		}
		if data[2][3] != 'a' {
			t.Errorf("expected rune 'a' but got something else")
		}
	})
}

func TestParseAllInts(t *testing.T) {
	t.Run("returns []int on sucess", func(t *testing.T) {
		str := `Button A: X+69, Y+23
				Button B: X+27, Y+71
				Prize: X=18641, Y=10279`
		ints, err := aoc.ParseAllInts(str)
		if err != nil {
			t.Errorf("expected nil err, but got err")
		}
		if len(ints) != 6 {
			t.Error("expected 6 ints, but got different amount: ", ints)
		}
		if ints[5] != 10279 {
			t.Errorf("epected 10279 but got %d", ints[2])
		}
	})

	t.Run("works with negative numbers", func(t *testing.T) {
		str := `p=95,12 v=-78,-79`
		ints, err := aoc.ParseAllInts(str)
		if err != nil {
			t.Errorf("expected nil err, but got err")
		}
		if len(ints) != 4 {
			t.Error("expected 6 ints, but got different amount: ", ints)
		}
		if ints[2] != -78 {
			t.Errorf("epected -78 but got %d", ints[2])
		}
	})
}
