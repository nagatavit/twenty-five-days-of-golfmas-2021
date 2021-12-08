package main

import "fmt"

func absVal(x int) int {
	if x >= 0 {
		return x
	} else {
		return -x
	}
}

func firstPart(crabsPosition []int) {
	fmt.Println("===================")
	fmt.Println("Starting first part")
	fmt.Println("===================")

	// The value that minimizes the total sum of errors (distances in
	// our case) is the median
	middleIdx := (len(crabsPosition) / 2) - 1

	middlePos := crabsPosition[middleIdx]

	fmt.Println("Tactical crab position:", middlePos)

	totalFuelSpent := 0

	for _, pos := range crabsPosition {
		totalFuelSpent += absVal(pos - middlePos)
	}

	fmt.Println("Crab total fuel:", totalFuelSpent)
}
