package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetLineStateWhenLineIsCorrupted(t *testing.T) {
	line := "{()}(]"
	state  := getLineState(line)
	assert.True(t, state.isCorrupted, "It should return isCorrupted true")
	assert.Equal(t, "]",state.illegalCharacter, "It should return the mismatched character")
}
func TestGetLineStateWhenLineIsComplete(t *testing.T) {
	line := "{()}()"
	state  := getLineState(line)
	assert.False(t, state.isCorrupted, "It should return isCorrupted false")
}

func TestGetLineStateWhenLineIsInComplete(t *testing.T) {
	line := "{()}("
	state  := getLineState(line)
	assert.False(t, state.isCorrupted, "It should return isCorrupted false")
}

