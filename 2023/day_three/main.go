package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type EnginePart struct {
	Value  int
	Length int
	Row    int
	Column int
}

type Coordinate struct {
	Row    int
	Column int
}

const inputFileName = "day_3.txt"
const validSymbols string = "/*&%+-$@#="

var lines []string
var linesCount int

func main() {
	parseInput()

	// find all numbers and create EnginePart objects
	potentialParts := makeEngineParts()

	// find neighbors of each engine part to decide if its a valid part
	var validParts []EnginePart
	for _, part := range potentialParts {
		if checkNeighbors(lines, part) {
			validParts = append(validParts, part)
		}
	}

	// add together valid parts
	total := 0
	for _, part := range validParts {
		total += part.Value
	}
	fmt.Printf("\nANSWER: %d\n", total)
}

func parseInput() {
	body, err := os.ReadFile(inputFileName)
	if err != nil {
		log.Fatalf("Unable to read file: %v", err)
	}
	lines = strings.Split(string(body), "\n")
	linesCount = len(lines)
}

func makeEngineParts() []EnginePart {
	var potentialParts []EnginePart
	for i, line := range lines {
		re := regexp.MustCompile("[0-9]+")
		numbers := re.FindAllString(string(line), -1)

		for j := 0; j < len(numbers); j++ {
			value, err := strconv.Atoi(numbers[j])
			if err != nil {
				log.Fatalf("Unable to convert string to int: %v", err)
			}

			// create new EnginePart objects
			part := EnginePart{
				Value:  value,
				Length: len(strconv.Itoa(value)),
				Row:    i,
				Column: strings.Index(line, numbers[j]),
			}
			potentialParts = append(potentialParts, part)
		}
	}
	return potentialParts
}

func checkNeighbors(lines []string, part EnginePart) bool {
	directions := makeDirections(part.Length)
	// fmt.Println(directions)

	indices := getNeighborCoordinate(lines[part.Row], part, directions)
	// fmt.Println(indices)

	for _, neighborCoordinate := range indices {
		neighbor := lines[neighborCoordinate.Row][neighborCoordinate.Column]
		if strings.Contains(validSymbols, string(neighbor)) {
			fmt.Printf("VALID\tRow: %d\tColumn: %d\tValue: %d\tSymbol: %c\n", part.Row, part.Column, part.Value, neighbor)
			return true
		}
	}
	fmt.Printf("INVALID\tRow: %d\tColumn: %d\tValue: %d\tSymbol: N/A\n", part.Row, part.Column, part.Value)
	return false
}

func getNeighborCoordinate(line string, part EnginePart, directions []Coordinate) []Coordinate {
	var coordinates []Coordinate
	for _, direction := range directions {
		index := Coordinate{part.Row + direction.Row, part.Column + direction.Column}
		if isValidIndex(index, line) {
			coordinates = append(coordinates, index)
		}
	}
	return coordinates
}

func makeDirections(partLength int) []Coordinate {
	directions := []Coordinate{
		{1, partLength},  // up_right
		{0, partLength},  // right
		{-1, partLength}, // down_right
		{1, -1},          // up_left
		{0, -1},          // left
		{-1, -1},         // down_left
	}

	for i := 0; i < partLength; i++ {
		directions = append(directions, Coordinate{1, 0 + i})  // up
		directions = append(directions, Coordinate{-1, 0 + i}) // down
	}

	// fmt.Println(directions)
	return directions
}

func isValidIndex(index Coordinate, line string) bool {
	return index.Column >= 0 && index.Column < len(line) && index.Row >= 0 && index.Row < linesCount
}
