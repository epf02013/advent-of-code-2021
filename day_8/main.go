package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
	    fmt.Printf("error opening file: %v\n", err)
	    os.Exit(1)
	}
	scanner := bufio.NewScanner(f)

	count := 0
	for scanner.Scan() {
	    text := scanner.Text()
	    rightSide := strings.TrimPrefix(strings.Split(text, "|")[1], " ")
		numberStrings := strings.Split(rightSide, " ")
		for _, numberString := range numberStrings {
			length := len(numberString)
			if length == 2 || length == 4 || length == 3 || length == 7 {
				count += 1
			}
		}
	}
	fmt.Println(count)
}
