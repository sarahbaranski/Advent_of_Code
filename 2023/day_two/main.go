package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func getSumIds(rolls string) int {
	gameInput := strings.Split(rolls, ": ")
	var maxBlue, maxGreen, maxRed int

	for _, pull := range strings.Split(strings.Trim(gameInput[1], " "), ";") {
		for _, dice := range strings.Split(pull, ",") {
			dice = strings.Trim(dice, " ")
			di := strings.Split(dice, " ")
			amount, _ := strconv.Atoi(di[0])

			switch di[1] {
			case "blue":
				if amount > maxBlue {
					maxBlue = amount
				}
			case "green":
				if amount > maxGreen {
					maxGreen = amount
				}
			case "red":
				if amount > maxRed {
					maxRed = amount
				}
			}
		}
	}

	return maxBlue * maxRed * maxGreen
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
