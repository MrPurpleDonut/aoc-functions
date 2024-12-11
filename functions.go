package aoc_functions

import (
	"os"
	"strconv"
	"strings"
)

func HandleError(err error) {
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

func MakeIntMatrix(filename string) ([][]int, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	vals := strings.Split(string(data), "\n")

	rows := make([][]int, len(vals))
	for i, row := range vals {
		rows[i] = make([]int, len(row))
		for j, char := range row {
			val, err := strconv.Atoi(string(char))
			if err != nil {
				return nil, err
			}
			rows[i][j] = val
		}
	}
	return rows, nil
}

func MakeBoolBoard(rows, cols int) [][]bool {
	board := make([][]bool, rows)
	for i := range board {
		board[i] = make([]bool, cols)
	}
	return board
}

func MakeRows(filename string) ([]string, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	return strings.Split(string(data), "\n"), nil
}
