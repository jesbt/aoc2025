package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("You must specify an input file")
		return
	}

	input := os.Args[1]
	file, err := os.Open(input)
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	dial := 50
	dialSize := 100
	password := 0

	for scanner.Scan() {
		line := scanner.Text()
		if len(line) < 2 {
			fmt.Println("Skipping invalid line:", line)
			continue
		}
		side := line[0]
		steps, err := strconv.Atoi(line[1:])

		check(err)
		switch side {
		case 'L':
			dial = (dial - steps) % dialSize

			// To fix negative modulo behaviour
			if dial < 0 {
				dial += dialSize
			}
		case 'R':
			dial = (dial + steps) % dialSize
		default:
			fmt.Println("Skipping invalid line:", line)
			continue
		}

		if dial == 0 {
			password++
		}
		fmt.Println(strconv.Itoa(dial))
	}
	check(scanner.Err())

	fmt.Printf("This is the password: %d\n", password)
}
