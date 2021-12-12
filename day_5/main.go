package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fileName := "input.txt"
	lines := readLineDefinitionsFromFile(fileName)
	part1(lines)
	part2(lines)
}

func part1(lines []Line) {
	validLines := filterOutNonHorizontalOrVerticalLines(lines)
	pointCounts := getIntersectionCountsAtEachPoint(validLines)
	pointsWithMoreThan2Intersections := numberOfPointsWithMoreThan2IntersectingLines(pointCounts)

	fmt.Println("Part 1 count:", pointsWithMoreThan2Intersections)
}

func part2(lines []Line) {
	pointCounts := getIntersectionCountsAtEachPoint(lines)
	pointsWithMoreThan2Intersections := numberOfPointsWithMoreThan2IntersectingLines(pointCounts)

	fmt.Println("Part 2 count:", pointsWithMoreThan2Intersections)
}

func numberOfPointsWithMoreThan2IntersectingLines(pointCounts map[Point]int) int {
	pointsWithMoreThan2Intersections := 0
	for _, v := range pointCounts {
		if v >= 2 {
			pointsWithMoreThan2Intersections += 1
		}
	}
	return pointsWithMoreThan2Intersections
}

func getIntersectionCountsAtEachPoint(validLines []Line) map[Point]int {
	pointCounts := make(map[Point]int)
	for _, line := range validLines {
		for _, point := range line.getPoints() {
			if val, ok := pointCounts[point]; ok {
				pointCounts[point] = val + 1
			} else {
				pointCounts[point] = 1
			}
		}
	}
	return pointCounts
}

func filterOutNonHorizontalOrVerticalLines(lines []Line) []Line {
	var validLines []Line
	for _, line := range lines {
		slope := line.simplifiedSlope()
		if slope.x == 0 || slope.y == 0 {
			validLines = append(validLines, line)
		}
	}
	return validLines
}

func readLineDefinitionsFromFile(fileName string) []Line {
	f, err := os.Open(fileName)
	if err != nil {
		fmt.Printf("error opening file: %v\n", err)
		os.Exit(1)
	}
	scanner := bufio.NewScanner(f)

	var lines []Line
	for scanner.Scan() {
		text := scanner.Text()
		line := ParseLine(text)
		lines = append(lines, line)
	}
	return lines
}
