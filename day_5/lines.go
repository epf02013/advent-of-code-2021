package main

import (
	"strconv"
	"strings"
)

type Slope struct {
	x int
	y int
}
type Point struct {
	x int
	y int
}
type Line struct {
	p0 Point
	p1 Point
}

func parsePoint(p0 string) Point {
	parts := strings.Split(p0, ",")
	x, _ := strconv.Atoi(parts[0])
	y, _ := strconv.Atoi(parts[1])
	return Point{x, y}
}
func ParseLine(lineString string) Line {
	parts := strings.Split(lineString, " -> ")
	p0 := parsePoint(parts[0])
	p1 := parsePoint(parts[1])
	return Line{
		p0: p0,
		p1: p1,
	}
}

func (line Line) getPoints() []Point {
	var points []Point
	simplifiedSlope := line.simplifiedSlope()
	var currentPoint = line.p0
	for currentPoint != line.p1  {
		points = append(points, currentPoint)
		currentPoint = Point{
			x: currentPoint.x + simplifiedSlope.x,
			y: currentPoint.y + simplifiedSlope.y,
		}
	}
	points = append(points, line.p1)
	return points
}

func (line Line) simplifiedSlope() Slope {
	deltaX := line.p1.x - line.p0.x
	deltaY := line.p1.y - line.p0.y
	gcd := GetGCD(deltaX, deltaY)
	return Slope{deltaX / gcd, deltaY / gcd}
}

func GetGCD(x int, y int) int {
	if x < 0 {
		x = x * -1
	}
	if y < 0 {
		x = y * -1
	}
	var a int
	var b int
	if x > y {
		a = x
		b = y
	} else {
		a = y
		b = x
	}
	for a > 0 && b > 0 {
		remainder := a % b
		a = b
		b = remainder
	}
	if a == 0 {
		return b
	}
	return a
}
