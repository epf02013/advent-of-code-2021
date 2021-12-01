package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		fmt.Printf("error opening file: %v\n",err)
		os.Exit(1)
	}
	scanner := bufio.NewScanner(f)
	var previousDepth = ""
	increases := 0
	for scanner.Scan() {
		currentDepth := scanner.Text()
		if previousDepth != "" && previousDepth < currentDepth {
			increases += 1
		}
		previousDepth = currentDepth
	}
	fmt.Println(increases)
}
