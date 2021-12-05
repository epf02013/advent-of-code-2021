package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func parsePickedNumbers(numPicks string) []int {
	var pickedNumbers []int
	for _, s := range strings.Split(numPicks,","){
		number, _ := strconv.Atoi(s)
		pickedNumbers = append(pickedNumbers, number)
	}
	return pickedNumbers
}

func parseBoards(scanner *bufio.Scanner) []Board{
	var boards []Board
	for scanner.Scan() {
		boardString := scanner.Text()
		board := ParseBoard(boardString)
		boards = append(boards, board)
	}
	return boards
}

func splitOnTwoNewLines(data []byte, atEOF bool) (advance int, token []byte, err error) {
	if atEOF && len(data) == 0 {
		return 0, nil, nil
	}

	if i := strings.Index(string(data), "\n\n"); i >= 0 {
		return i + 1, data[0:i], nil
	}

	if atEOF {
		return len(data), data, nil
	}

	return
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		fmt.Printf("error opening file: %v\n", err)
		os.Exit(1)
	}
	scanner := bufio.NewScanner(f)

	scanner.Split(splitOnTwoNewLines)
	scanner.Scan()
	pickedNumbers := parsePickedNumbers(scanner.Text())
	boards := parseBoards(scanner)
	winningBoard,_ := getFirstBoardToWin(pickedNumbers, boards)
	fmt.Println(winningBoard.calculateScore())
}

func getFirstBoardToWin(pickedNumbers []int, boards []Board) (Board,error) {
	for _, number := range pickedNumbers {
		for _, board := range boards {
			board.mark(number)
			if board.hasWon() {
				return board, nil
			}
		}
	}
	return Board{}, errors.New("no winning board")
}
