package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Vector struct {
	depth      int
	horizontal int
}

func (v1 Vector) add(v2 Vector) Vector {
	return Vector{
		depth:      v1.depth + v2.depth,
		horizontal: v1.horizontal + v2.horizontal,
	}
}

func parseVector(command string) Vector {
	parts := strings.Split(command, " ")
	commandType := parts[0]
	distance, _ := strconv.Atoi(parts[1])
	if commandType == "forward" {
		return Vector{depth: 0, horizontal: distance}
	}
	if commandType == "down" {
		return Vector{depth: distance, horizontal: 0}
	}
	return Vector{depth: -distance, horizontal: 0}
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		fmt.Printf("error opening file: %v\n", err)
		os.Exit(1)
	}
	scanner := bufio.NewScanner(f)
	currentPosition := Vector{depth: 0, horizontal: 0}
	for scanner.Scan() {
		vector := parseVector(scanner.Text())
		currentPosition = currentPosition.add(vector)
	}
	fmt.Println(currentPosition.horizontal * currentPosition.depth)
}
