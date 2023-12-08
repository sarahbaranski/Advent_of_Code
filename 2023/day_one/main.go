package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const search = string("0123456789")

func extractInput() []byte {
	body, err := os.ReadFile("text.txt")
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}
	return body
}

func searchLeftToRight(line string) string {
	for i := 0; i < len(line); i++ {
		currentChar := string(line[i])

		if strings.Contains(search, currentChar) {
			return currentChar
		}
	}
	return ""
}

func searchRightToLeft(line string) string {
	for i := len(line) - 1; i >= 0; i-- {
		currentChar := string(line[i])

		if strings.Contains(search, currentChar) {
			return currentChar
		}
	}
	return ""
}

func sumInts(combined []string) (int, error) {
	sum := 0
	for _, str := range combined {
		n, err := strconv.Atoi(string(str))
		if err != nil {
			log.Fatalf("unable to convert %v", err)
		}

		sum += n
	}
	return sum, nil
}

func main() {
	// extracts input text into byte[]
	input := extractInput()

	var combined []string
	for _, line := range strings.Split(string(input), "\n") {
		numString := searchLeftToRight(line) + searchRightToLeft(line)
		combined = append(combined, numString)
	}

	total, _ := sumInts(combined)
	fmt.Println(total)
}
