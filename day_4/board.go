package main

import (
	"strconv"
	"strings"
)

type Item struct {
	row    int
	col    int
	value  int
	marked bool
}

type Board struct {
	rowCounts     map[int]int
	colCounts     map[int]int
	items         map[int]Item
	winningNumber int
}

func ParseBoard(boardString string) Board {
	boardString = strings.TrimPrefix(boardString, "\n\n")
	rowCounts := make(map[int]int)
	colCounts := make(map[int]int)
	for i := 0; i < 5; i++ {
		rowCounts[i] = 0
		colCounts[i] = 0
	}
	items := make(map[int]Item)
	for row, lineString := range strings.Split(boardString, "\n") {
		numberStrings := strings.Split(lineString, " ")
		col := 0
		for _, s := range numberStrings {
			number, err := strconv.Atoi(s)
			if err == nil {
				item := Item{
					row:   row - 1,
					col:   col,
					value: number,
				}
				items[number] = item
				col += 1
			}
		}
	}
	board := Board{
		rowCounts: rowCounts,
		colCounts: colCounts,
		items:     items,
	}
	return board
}

func (board *Board) mark(pickedNumber int) {
	value, ok := board.items[pickedNumber]
	if ok {
		value.marked = true
		board.items[pickedNumber] = value
		board.rowCounts[value.row] += 1
		board.colCounts[value.col] += 1
		if board.hasWon() {
			board.winningNumber = pickedNumber
		}
	}
}

func (board Board) hasWon() bool {
	for _, v := range board.colCounts {
		if v == 5 {
			return true
		}
	}
	for _, v := range board.rowCounts {
		if v == 5 {
			return true
		}
	}
	return false
}
func (board Board) calculateScore() int {
	sumOfUnmarkedNumbers := 0
	for _, item := range board.items {
		if !item.marked {
			sumOfUnmarkedNumbers += item.value
		}
	}
	return board.winningNumber * sumOfUnmarkedNumbers
}
