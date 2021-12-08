package main

import "fmt"

func sortTwoValues(a, b int) (begin, finish int) {
	if a > b {
		return b, a
	} else {
		return a, b
	}
}

func markCoordinatesFromPerpendicularLines(begin, finish coordinate) (marked []coordinate) {
	if begin.x != finish.x && begin.y != finish.y {
		return nil
	}

	if begin.x == finish.x {
		startY, finishY := sortTwoValues(begin.y, finish.y)
		for i := startY; i <= finishY; i++ {
			marked = append(marked,
				coordinate{x: begin.x, y: i},
			)
		}
		return marked
	}

	if begin.y == finish.y {
		startX, finishX := sortTwoValues(begin.x, finish.x)
		for i := startX; i <= finishX; i++ {
			marked = append(marked,
				coordinate{x: i, y: begin.y},
			)
		}
		return marked
	}

	return marked
}

func firstPart(lines []ventLine) {
	fmt.Println("===================")
	fmt.Println("Starting first part")
	fmt.Println("===================")

	filledLines := make(map[coordinate]int)

	for _, line := range lines {
		coords := markCoordinatesFromPerpendicularLines(line.coord1, line.coord2)
		for _, coord := range coords {
			filledLines[coord]++
		}
	}

	overlapCount := 0

	for _, coordCount := range filledLines {
		if coordCount >= 2 {
			overlapCount++
		}
	}

	fmt.Println(overlapCount)
}
