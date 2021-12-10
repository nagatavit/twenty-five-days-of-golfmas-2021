package main

import (
	"fmt"
)

// sides -> top, right, bottom, left
func checkIfMinimum(sides []int, middle int) bool {
	for _, side := range sides {
		if side < 0 {
			continue
		}

		if middle >= side {
			return false
		}
	}
	return true
}

func getSides(i, j int, heightmap *[][]int) (top, right, bottom, left int) {
	if i-1 < 0 {
		top = -1
	} else {
		top = (*heightmap)[i-1][j]
	}

	if j+1 > len((*heightmap)[0])-1 {
		right = -1
	} else {
		right = (*heightmap)[i][j+1]
	}

	if i+1 > len((*heightmap)[0])-1 {
		bottom = -1
	} else {
		bottom = (*heightmap)[i+1][j]
	}

	if j-1 < 0 {
		left = -1
	} else {
		left = (*heightmap)[i][j-1]
	}
	return top, right, bottom, left
}

func firstPart(heightmap [][]int) {
	riskLevel := 0

	for i, line := range heightmap {
		for j, middle := range line {
			top, right, bottom, left := getSides(i, j, &heightmap)

			if checkIfMinimum([]int{top, right, bottom, left}, middle) {
				riskLevel += middle + 1
			}

		}
	}

	fmt.Println("risk level:", riskLevel)
}
