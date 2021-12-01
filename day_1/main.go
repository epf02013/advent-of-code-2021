package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	part1()
	part2()
}

type RollingSum struct {
	startIndex int
	endIndex int
	array *[]int}
func (rs RollingSum) addElement(index int) RollingSum{
	if rs.endIndex-rs.startIndex >= 2 {
		rs.startIndex += 1
	}
	rs.endIndex = index
	return rs
}

func (rs *RollingSum) isFull() bool {
	return	rs.endIndex-rs.startIndex >= 2
}
func (rs *RollingSum) sum() int{
	sum := 0
	for i := rs.startIndex; i <= rs.endIndex; i++ {
		sum += (*rs.array)[i]
	}
	return sum
}

func part2() {
	fmt.Println("Part 2")
	f, err := os.Open("input.txt")
	if err != nil {
		fmt.Printf("error opening file: %v\n", err)
		os.Exit(1)
	}
	scanner := bufio.NewScanner(f)
	array  := []int{}
	for scanner.Scan() {
		parseInt, err := strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Printf("error parsing int %v\n", err)
			os.Exit(1)
		}
		array = append(array, parseInt)
	}
	trailingSum := RollingSum{array: &array, startIndex: 0, endIndex: 0}
	leadingSum := RollingSum{array: &array, startIndex: 1, endIndex: 1}
	increases := 0

	for i := 2; i < len(array); i++ {
		trailingSum = trailingSum.addElement(i-1)
		leadingSum =leadingSum.addElement(i)
		if trailingSum.isFull() && leadingSum.isFull() {
			if trailingSum.sum() < leadingSum.sum() {
				increases += 1
			}
		}
	}
	fmt.Println(increases)
}

func part1() {
	fmt.Println("Part 1")
	f, err := os.Open("input.txt")
	if err != nil {
		fmt.Printf("error opening file: %v\n", err)
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
