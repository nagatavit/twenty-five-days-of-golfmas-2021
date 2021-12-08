package main

import "fmt"

func secondPart(crabsPosition []int) {
	fmt.Println("===================")
	fmt.Println("Starting second part")
	fmt.Println("===================")

	maxCrabPosition := crabsPosition[len(crabsPosition)-1] + 1

	// Slice to map the distance by the fuel wasted
	distanceWeight := make([]int, maxCrabPosition)
	for i := 1; i < maxCrabPosition; i++ {
		distanceWeight[i] = distanceWeight[i-1] + i
	}

	// Count how many crabs are in each position
	crabCountByPosition := make([]int, maxCrabPosition)
	for _, crabPos := range crabsPosition {
		crabCountByPosition[crabPos]++
	}

	currentBestCandidate := int(^uint(0) >> 1)
	bestPosition := 0

	for i := 0; i < len(crabCountByPosition); i++ {
		fuelUsed := 0

		for j := 0; j < len(crabCountByPosition); j++ {
			if crabCountByPosition[j] == 0 || i == j {
				continue
			}

			fuelUsed += crabCountByPosition[j] * distanceWeight[absVal(j-i)]
		}

		if fuelUsed <= currentBestCandidate {
			bestPosition = i
			currentBestCandidate = fuelUsed
		}
	}

	fmt.Println("Tactical crab position:", bestPosition)
	fmt.Println("Crab total fuel:", currentBestCandidate)
}
