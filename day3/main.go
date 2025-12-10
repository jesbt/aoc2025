package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		log.Fatalf("Error: %v", e)
	}
}

func partOne(input string) {
	banks := strings.Split(input, "\n")
	output := 0
	for _, bank := range banks {
		if strings.TrimSpace(bank) == "" {
			continue
		}
		largestNumber := 0
		for i := 0; i < len(bank); i++ {
			for j := i + 1; j < len(bank); j++ {
				n, err := strconv.Atoi((string(bank[i]) + string(bank[j])))
				check(err)
				if n > largestNumber {
					largestNumber = n
				}
			}
		}
		output += largestNumber
	}
	fmt.Println(output)
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Usage: main ./input.txt")
	}

	file, err := os.ReadFile(os.Args[1])
	check(err)

	partOne(string(file))
}
