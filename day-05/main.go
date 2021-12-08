package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type coordinate struct {
	x, y int
}

type ventLine struct {
	coord1, coord2 coordinate
}

func readInput() (lines []ventLine) {
	f, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	var x1, y1, x2, y2 int

	for scanner.Scan() {
		_, err := fmt.Sscanf(scanner.Text(), "%d,%d -> %d,%d", &x1, &y1, &x2, &y2)
		if err != nil {
			log.Fatal(err)
		}

		lines = append(lines,
			ventLine{
				coord1: coordinate{
					x: x1,
					y: y1,
				},
				coord2: coordinate{
					x: x2,
					y: y2,
				},
			},
		)
	}

	return lines
}

func main() {
	ventLines := readInput()
	firstPart(ventLines)
	secondPart(ventLines)
}
