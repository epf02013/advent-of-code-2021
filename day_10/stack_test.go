package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPop(t *testing.T) {
	stack := Stack{items: []string{"<", "{"}}
	res, item, _ := stack.pop()
	assert.Equal(t, []string{"<"}, res.items, "It should slice the items")
	assert.Equal(t, "{", item, "It should return the final element")
}

func TestPeek(t *testing.T) {
	stack := Stack{items: []string{"<", "{"}}
	item := stack.peek()
	assert.Equal(t, "{", item, "It should return the final element")
}

func TestPush(t *testing.T) {
	stack := Stack{items: []string{"<"}}
	res := stack.push("{")
	assert.Equal(t, []string{"<", "{"}, res.items, "It should append the string to items")
}

func TestSize(t *testing.T) {
	stack := Stack{items: []string{"<", "{"}}
	size := stack.size()
	assert.Equal(t, 2, size, "It should return the number of items")
}
