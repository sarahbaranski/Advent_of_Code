package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// max color dice
var blue = 14
var red = 12
var green = 13

func getSumIds(rolls string) int {
	gameInput := strings.Split(rolls, ": ")
	gameId := strings.Split(gameInput[0], " ")
	idNum, _ := strconv.Atoi(gameId[1])

	for _, pull := range strings.Split(strings.Trim(gameInput[1], " "), ";") {
		for _, dice := range strings.Split(pull, ",") {
			dice = strings.Trim(dice, " ")
			di := strings.Split(dice, " ")
			amount, _ := strconv.Atoi(di[0])

			switch di[1] {
			case "blue":
				if amount > blue {
					return 0
				}
			case "green":
				if amount > green {
					return 0
				}
			case "red":
				if amount > red {
					return 0
				}
			}
		}
	}

	return idNum
}

func main() {
	var sumId int
	body, err := os.ReadFile("day_2.txt")
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}

	for _, rolls := range strings.Split(string(body), "\n") {
		sumId += getSumIds(rolls)
	}

	fmt.Println(sumId)
}
