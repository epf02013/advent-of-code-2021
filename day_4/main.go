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
	part1()
	part2()
}

func part1() {
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
	winningBoard, _ := getFirstBoardToWin(pickedNumbers, boards)
	fmt.Println("Part 1 score:", winningBoard.calculateScore())
}

func part2() {
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
	lastBoardToWin, _ := getLastBoardToWin(pickedNumbers, boards)
	fmt.Println("Part 2 score:", lastBoardToWin.calculateScore())
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

func getLastBoardToWin(pickedNumbers []int, boards []Board) (Board,error) {
	candidateBoards := boards
	for _, number := range pickedNumbers {
		var winningBoards []Board
		var losingBoards []Board
		for _, board := range candidateBoards {
			board.mark(number)
			if board.hasWon() {
				winningBoards = append(winningBoards, board)
			} else {
				losingBoards = append(losingBoards, board)
			}
		}
		if len(winningBoards) == 1 && len(candidateBoards) == 1 {
			return winningBoards[0], nil
		}
		candidateBoards = losingBoards
	}
	return Board{}, errors.New("no winning board")
}
