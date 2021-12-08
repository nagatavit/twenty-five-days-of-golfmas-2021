package main

import "fmt"

// getLineDirection gets two y values and get the line direction (signal)
func getLineDirection(y1, y2 int) int {
	if y2-y1 > 0 {
		return 1
	} else {
		return -1
	}
	// The zero case shouldn't happen (at least, this shouldn't be
	// called when y1 == y2)
}

func sortTwoCoordsByX(coord1, coord2 coordinate) (begin, finish coordinate, direction int) {
	if coord1.x > coord2.x {
		return coord2, coord1, getLineDirection(coord2.y, coord1.y)
	} else {
		return coord1, coord2, getLineDirection(coord1.y, coord2.y)
	}
}

func markCoordinatesFromLines(begin, finish coordinate) (marked []coordinate) {
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

	// 45 degree angle
	startCoord, finishCoord, dir := sortTwoCoordsByX(begin, finish)
	j := startCoord.y
	for i := startCoord.x; i <= finishCoord.x; i++ {
		marked = append(marked,
			coordinate{x: i, y: j},
		)
		j += dir
	}

	return marked
}

func secondPart(lines []ventLine) {
	fmt.Println("====================")
	fmt.Println("Starting second part")
	fmt.Println("====================")

	filledLines := make(map[coordinate]int)

	for _, line := range lines {
		coords := markCoordinatesFromLines(line.coord1, line.coord2)
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
