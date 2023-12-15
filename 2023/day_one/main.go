package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func extractInput() []byte {
	body, err := os.ReadFile("text.txt")
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}

	return body
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

func searchForNumbers(input []byte) []string {
	var combined []string
	right := -1
	left := 0

	numbers := map[string]int{
		"zero":  0,
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
		"1":     1,
		"2":     2,
		"3":     3,
		"4":     4,
		"5":     5,
		"6":     6,
		"7":     7,
		"8":     8,
		"9":     9,
	}

	for _, line := range strings.Split(string(input), "\n") {
		for i := 0; i < len(line); i++ {
			word := line[i:]
			for k, v := range numbers {
				if strings.HasPrefix(word, k) {
					right = v
					break
				}
			}
		}

		for i := len(line); i >= 0; i-- {
			word := line[i:]
			for k, v := range numbers {
				if strings.HasPrefix(word, k) {
					left = v
					break
				}
			}
		}

		combinedString := strconv.Itoa(left) + strconv.Itoa(right)
		combined = append(combined, combinedString)
	}

	return combined
}

func main() {
	// extracts input text into byte[]
	input := extractInput()

	total, _ := sumInts(searchForNumbers(input))
	fmt.Println(total)
}
