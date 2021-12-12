package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	repeatingLifeSpan := 6
	maxLifeSpan := 8

	input := getInitialFishTimers("input.txt")
	sumFor18Days := getNumberOfFishAfterXDays(maxLifeSpan, input, 18, repeatingLifeSpan)
	fmt.Println("Sum for 18 days", sumFor18Days)
	sumFor80Days := getNumberOfFishAfterXDays(maxLifeSpan, input, 80, repeatingLifeSpan)
	fmt.Println("Sum for 80 days", sumFor80Days)
	sumFor256Days := getNumberOfFishAfterXDays(maxLifeSpan, input, 256, repeatingLifeSpan)
	fmt.Println("Sum for 80 days", sumFor256Days)
}

func getInitialFishTimers(fileName string) []int {
	f, err := os.Open(fileName)
	if err != nil {
		fmt.Printf("error opening file: %v\n", err)
		os.Exit(1)
	}
	scanner := bufio.NewScanner(f)

	scanner.Scan()
	text := scanner.Text()
	numberStrings := strings.Split(text, ",")
	var numbers []int
	for _, numString := range numberStrings {
		number, _ := strconv.Atoi(numString)
		numbers = append(numbers, number)
	}
	return numbers
}

func getNumberOfFishAfterXDays(maxFishTimer int, initialFishTimers []int, totalNumberOfDays int, repeatingFishTimer int) int {
	buckets := initializeBuckets(maxFishTimer, initialFishTimers)
	for day := 1; day < totalNumberOfDays+1; day++ {
		numberTurningOver := buckets[0]
		for i := 0; i < len(buckets)-1; i++ {
			buckets[i] = buckets[i+1]
		}
		buckets[len(buckets)-1] = numberTurningOver
		buckets[repeatingFishTimer] += numberTurningOver
	}
	sum := 0
	for _, value := range buckets {
		sum += value
	}
	return sum
}

func initializeBuckets(maxFishTimer int, initialFishTimers []int) []int {
	var buckets []int
	for i := 0; i < maxFishTimer+1; i++ {
		buckets = append(buckets, 0)
	}
	for _, lifeSpan := range initialFishTimers {
		buckets[lifeSpan] += 1
	}
	return buckets
}
