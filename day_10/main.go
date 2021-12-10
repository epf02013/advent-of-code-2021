package main

import "strings"

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
