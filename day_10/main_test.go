package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetLineStateWhenLineIsCorrupted(t *testing.T) {
	line := "{()}(]"
	state := getLineState(line)
	assert.True(t, state.isCorrupted, "It should return isCorrupted true")
	assert.Equal(t, "]", state.illegalCharacter, "It should return the mismatched character")
}
func TestGetLineStateWhenLineIsComplete(t *testing.T) {
	line := "{()}()"
	state := getLineState(line)
	assert.False(t, state.isCorrupted, "It should return isCorrupted false")
	assert.False(t, state.incomplete, "It should return incomplete false")
}

func TestGetLineStateWhenLineIsIncomplete(t *testing.T) {
	line := "{()({}["
	state := getLineState(line)
	assert.False(t, state.isCorrupted, "It should return isCorrupted false")
	assert.True(t, state.incomplete, "It should return incomplete true")
	assert.Equal(t, []string{"]", ")", "}"}, state.completionCharacters, "It should return the missing closing characters")
}

func TestGetPoints(t *testing.T) {
	state1 := LineState{isCorrupted: true, illegalCharacter: ")"}
	state2 := LineState{isCorrupted: true, illegalCharacter: ">"}
	state3 := LineState{isCorrupted: false}
	lineStates := []LineState{state1, state2, state3}
	points := getPointsForPart1(lineStates)
	assert.Equal(t, 25140, points)
}
