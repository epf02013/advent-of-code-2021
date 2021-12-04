package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
	    fmt.Printf("error opening file: %v\n", err)
	    os.Exit(1)
	}
	scanner := bufio.NewScanner(f)
	bitSums, rowCount := getBitSumsAndRowCount(scanner)
	commonBits := getCommonBits(bitSums, rowCount)
	gammaRate := parseIntFromArray(commonBits)
	mask := getMask(commonBits)
	epsilonRate := mask ^ gammaRate
	fmt.Println(gammaRate * epsilonRate)
}

func getCommonBits(bitSums []int, rowCount int) []int {
	var commonBits []int
	for _, bitSum := range bitSums {
		var bit int
		if (float64(bitSum) / float64(rowCount)) >= .5 {
			bit = 1
		} else {
			bit = 0
		}
		commonBits = append(commonBits, bit)
	}
	return commonBits
}

func parseIntFromArray(bitArray []int) int64 {
	var gammaStringArr []string
	for _, bit := range bitArray {
		gammaStringArr = append(gammaStringArr, strconv.FormatInt(int64(bit), 2))
	}
	gammaString := strings.Join(gammaStringArr, "")
	gammaRate, _ := strconv.ParseInt(gammaString, 2, 64)
	return gammaRate
}

func getMask(commonBits []int) int64 {
	maskBits := []string{}
	for range commonBits {
		maskBits = append(maskBits, "1")
	}
	mask, _ := strconv.ParseInt(strings.Join(maskBits, ""), 2, 64)
	return mask
}

func getBitSumsAndRowCount(scanner *bufio.Scanner) ([]int, int) {
	sumBits := []int{0, 0, 0, 0, 0}
	rowCount := 0
	for scanner.Scan() {
		text := scanner.Text()
		if len(text) < 1 {
			continue
		}
		currentBits := strings.Split(text, "")
		for i, bitStr := range currentBits {
			bitInt, _ := strconv.Atoi(bitStr)
			if len(sumBits) == i {
				sumBits = append(sumBits, bitInt)
			} else {
				sumBits[i] = sumBits[i] + bitInt
			}
		}
		rowCount += 1
	}
	return sumBits, rowCount
}
