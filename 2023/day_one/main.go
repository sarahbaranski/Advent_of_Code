package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	body, err := os.ReadFile("text.txt")
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}

	search := string("0123456789")
	var left, right string
	var combined []string
	for _, lines := range strings.Split(string(body), "\n") {

		for i := 0; i < len(lines); i++ {
			if strings.Contains(search, string(lines[i])) {
				right = string(lines[i])
			}
		}

		for i := len(lines) - 1; i >= 0; i-- {
			if strings.Contains(search, string(lines[i])) {
				left = string(lines[i])
			}
		}

		combined = append(combined, left+right)
	}

	total := 0

	for _, str := range combined {
		n, err := strconv.Atoi(string(str))
		if err != nil {
			log.Fatalf("unable to convert %v", err)
		} else {
			total += n
		}
	}

	fmt.Println(total)
}
