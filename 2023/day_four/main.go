package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

type Card struct {
	Winning []string
	Given   []string
	Points  int
}

const inputFileName = "day_4.txt"

var lines []string

func main() {
	parseInput()
	cards := makeCards()
	cards = calcPoints(cards)
	total := sumPoints(cards)
	fmt.Printf("\nANSWER: %d\n", total)
}

func parseInput() {
	body, err := os.ReadFile(inputFileName)
	if err != nil {
		log.Fatalf("Unable to read file: %v", err)
	}
	lines = strings.Split(string(body), "\n")
}

func makeCards() []Card {
	var cards []Card
	for _, line := range lines {
		values := strings.Split(strings.Split(line, ":")[1], "|")

		var card Card
		for i, v := range values {
			re := regexp.MustCompile("[0-9]+")
			numbers := re.FindAllString(v, -1)
			if i < 1 {
				card.Winning = numbers
			} else {
				card.Given = numbers
			}
		}
		cards = append(cards, card)
		// fmt.Println(card)
	}
	return cards
}

func calcPoints(cards []Card) []Card {
	for i := 0; i < len(cards); i++ {
		for _, g := range cards[i].Given {
			for _, w := range cards[i].Winning {
				if g == w {
					if cards[i].Points == 0 {
						cards[i].Points = 1
					} else {
						cards[i].Points *= 2
					}
					break
				}
			}
		}
		// fmt.Println(cards[i])
	}
	return cards
}

func sumPoints(cards []Card) int {
	sum := 0
	for _, card := range cards {
		sum += card.Points
		fmt.Println(card)
	}
	return sum
}
