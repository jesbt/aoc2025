package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func check(e error) {
	if e != nil {
		log.Fatalf("Error: %v", e)
	}
}

func partOne(filePath string) {
	file, err := os.Open(filePath)
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	dial := 50
	dialSize := 100
	zeros := 0

	for scanner.Scan() {
		line := scanner.Text()
		if len(line) < 2 {
			log.Println("Skipping invalid line: ", line)
			continue
		}

		direction := line[0]
		steps, err := strconv.Atoi(line[1:])
		check(err)

		if steps >= dialSize {
			steps = steps % dialSize
		}

		switch direction {
		case 'L':
			nextDial := dial - steps
			if nextDial < 0 {
				nextDial = dialSize + nextDial
			}
			dial = nextDial
		case 'R':
			nextDial := dial + steps
			if nextDial >= dialSize {
				nextDial = nextDial % dialSize
			}
			dial = nextDial
		default:
			log.Printf("Skipping invalid direction: %c", direction)
			continue
		}
		if dial == 0 {
			zeros++
		}
	}
	check(scanner.Err())
	fmt.Println(zeros)
}

func partTwo(filePath string) {
	file, err := os.Open(filePath)
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	dial := 50
	dialSize := 100
	zeros := 0

	for scanner.Scan() {
		line := scanner.Text()
		if len(line) < 2 {
			log.Println("Skipping invalid line: ", line)
			continue
		}

		direction := line[0]
		steps, err := strconv.Atoi(line[1:])
		check(err)

		if steps >= dialSize {
			zeros += steps / dialSize
			steps = steps % dialSize
		}

		switch direction {
		case 'L':
			nextDial := dial - steps
			if nextDial < 0 {
				if dial > 0 {
					zeros++
				}
				nextDial = dialSize + nextDial
			}
			dial = nextDial
		case 'R':
			nextDial := dial + steps
			if nextDial >= dialSize {
				if nextDial != dialSize {
					zeros++
				}
				nextDial = nextDial % dialSize
			}
			dial = nextDial
		default:
			log.Println("Skipping invalid direction")
			continue
		}
		if dial == 0 {
			zeros++
		}
	}
	check(scanner.Err())
	fmt.Println(zeros)
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("You must specify an input file")
		return
	}

	input := os.Args[1]
	partOne(input)
	partTwo(input)
}
