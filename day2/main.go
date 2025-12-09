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

func partOne(filePath string) {
	file, err := os.ReadFile(filePath)
	check(err)

	idRanges := strings.Split(string(file), ",")

	invalidIds := 0

	for _, row := range idRanges {
		idRange := strings.Split(row, "-")
		if len(idRange) != 2 {
			log.Println("Skipping invalid range: ", row)
			continue
		}

		start, err := strconv.Atoi(idRange[0])
		check(err)
		end, err := strconv.Atoi(idRange[1])
		check(err)

		for i := start; i <= end; i++ {
			id := strconv.Itoa(i)
			if len(id)%2 == 0 {
				middleIndex := len(id) / 2
				leftSide := id[:middleIndex]
				rightSide := id[middleIndex:]

				// Check if the two halves of the ID are identical
				if leftSide == rightSide {
					invalidIds += i
				}
			}
		}

	}

	fmt.Println(invalidIds)
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Usage: main ./input.txt")
	}

	partOne(os.Args[1])
}
