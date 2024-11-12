package aoc_functions

import (
	"os"
	"strings"
)

func PanicOnErr(err error) {
	if err != nil {
		panic(err)
	}
}

func MakeMatrix(filename string) ([][]rune, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	vals := strings.Split(string(data), "\n")

	rows := make([][]rune, len(vals))
	for i, row := range vals {
		rows[i] = []rune(row)
	}
	return rows, nil
}

func MakeRows(filename string) ([]string, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	return strings.Split(string(data), "\n"), nil
}
