package main

import (
	"bufio"
	"fmt"
	"github.com/montanaflynn/stats"
	"os"
	"strings"
)

type LineState struct {
	isCorrupted      bool
	illegalCharacter     string
	incomplete           bool
	completionCharacters []string
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
func getClosingCharacter(c string) string{

	characterMap  := map[string]string{
		"(": ")",
		"{": "}",
		"[": "]",
		"<": ">",
	}
	return characterMap[c]
}
func charactersMatch(a string, b string) bool {
	return getClosingCharacter(a) == b
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
	if stack.size() == 0 {
		return LineState{
			isCorrupted: false,
			incomplete:  false,
		}
	}
	completionCharacters  := []string{}

	for stack.size() > 0 {
		item := ""
		stack, item, _ = stack.pop()
		completionCharacters = append(completionCharacters, getClosingCharacter(item))
	}
	return LineState{incomplete: true, completionCharacters: completionCharacters}
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

func getPointsForPart1(states []LineState) int {
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
func getPointsForPart2(states []LineState) int {
	pointMap := map[string]int{
		")": 1,
		"]": 2,
		"}": 3,
		">": 4,
	}
	scores := []float64{}
	for _, state := range states {
		if state.incomplete {
			score := 0
			for _, character := range state.completionCharacters {
				score = score * 5
				score += pointMap[character]
			}
			scores = append(scores, float64(score))
		}
	}
	median,_ := stats.Median(scores)
	return int(median)
}

func main() {
	lineStates := getLineStates("input.txt")
	points := getPointsForPart1(lineStates)
	fmt.Println("Part 1",points)

	pointsForPart2 := getPointsForPart2(lineStates)
	fmt.Println("Part 2",pointsForPart2)
}

