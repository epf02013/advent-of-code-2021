package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)


type TravelVector struct {
	aim int
	depth int
	horizontal int
}

type Command struct {
	aim int
	distance int
}


func parseCommand(command string) Command {
	parts := strings.Split(command, " ")
	commandType := parts[0]
	distance, _ := strconv.Atoi(parts[1])
	if commandType == "forward"{
		return Command{aim: 0, distance: distance}
	}
	if commandType == "down"{
		return Command{aim: distance, distance: 0}
	}
	return Command{aim: -distance, distance: 0}
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		fmt.Printf("error opening file: %v\n", err)
		os.Exit(1)
	}
	scanner := bufio.NewScanner(f)
	currentPosition := TravelVector{aim: 0, horizontal: 0, depth: 0}
	for scanner.Scan() {
		command := parseCommand(scanner.Text())
		currentPosition = TravelVector{
			depth: currentPosition.depth + command.distance*currentPosition.aim,
			horizontal: currentPosition.horizontal+command.distance,
			aim: currentPosition.aim + command.aim,
		}
	}
	fmt.Println(currentPosition.horizontal*currentPosition.depth)
}
