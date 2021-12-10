package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type LineState struct {
	isCorrupted           bool
	illegalCharacter      string
}


func isOpeningCharacter(c string) bool {
	openingCharacterMap  := map[string]bool{
		"(": true,
		"{": true,
		"[": true,
		"<": true,
	}
	return openingCharacterMap[c]
}
func charactersMatch(a string, b string) bool {
	characterMap  := map[string]string{
		"(": ")",
		"{": "}",
		"[": "]",
		"<": ">",
	}
	return characterMap[a] == b
}
func getLineState(line string)  LineState {
	characters  := strings.Split(line, "")
	stack := Stack{}
	for _, character := range characters {
		if isOpeningCharacter(character) {
			stack = stack.push(character)
		} else if stack.size() > 0 && charactersMatch(stack.peek(), character){
			stack, _, _ = stack.pop()
		} else {
			return LineState{isCorrupted: true, illegalCharacter: character}
		}
	}
	return LineState{
		isCorrupted: false,
	}
}

func getPoints(states []LineState) int {
	pointMap := map[string]int{
		")": 3,
		"]": 57,
		"}": 1197,
		">": 25137,
	}
	sum := 0
	for _, state := range states {
		if state.isCorrupted {
			sum += pointMap[state.illegalCharacter]
		}
	}
	return sum
}

func main() {
	lineStates := getLineStates("input.txt")
	points := getPoints(lineStates)
	fmt.Println("Part 1",points)
}

func getLineStates(fileNmae string) []LineState {
	f, err := os.Open(fileNmae)
	if err != nil {
		fmt.Printf("error opening file: %v\n", err)
		os.Exit(1)
	}
	scanner := bufio.NewScanner(f)
	lineStates := []LineState{}
	for scanner.Scan() {
		text := scanner.Text()
		lineStates = append(lineStates, getLineState(text))
	}
	return lineStates
}
