package main

import (
	"bufio"
	"fmt"
	"github.com/montanaflynn/stats"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	crabPositions := parseCrabPositionsFromFile("input.txt")

	part1(crabPositions)

	part2(crabPositions)
}

func part2(crabPositions []float64) {
	mean, _ := stats.Mean(crabPositions)
	floor := math.Floor(mean)
	ceiling := math.Floor(mean)
	floorSum := 0.0
	ceilingSum := 0.0
	for _, number := range crabPositions {
		n := math.Abs(floor - number)
		floorSum += (n * (n + 1)) / 2

		n = math.Abs(ceiling - number)
		ceilingSum += (n * (n + 1)) / 2
	}
	fmt.Println("Part 2 sum", int(math.Min(ceilingSum, floorSum)))
}

func part1(crabPositions []float64) {
	median, _ := stats.Median(crabPositions)

	roundedMedian, _ := stats.Round(median, 0)
	sum := 0.0
	for _, number := range crabPositions {
		sum += math.Abs(roundedMedian - number)
	}
	fmt.Println("Part 1 sum:", sum)
}

func parseCrabPositionsFromFile(fileName string) []float64 {
	f, err := os.Open(fileName)
	if err != nil {
		fmt.Printf("error opening file: %v\n", err)
		os.Exit(1)
	}
	scanner := bufio.NewScanner(f)
	scanner.Scan()
	text := scanner.Text()
	numStrings := strings.Split(text, ",")
	var numbers []float64
	for _, numString := range numStrings {
		number, _ := strconv.ParseFloat(numString, 64)
		numbers = append(numbers, number)
	}
	return numbers
}
