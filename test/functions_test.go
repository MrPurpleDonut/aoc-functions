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
