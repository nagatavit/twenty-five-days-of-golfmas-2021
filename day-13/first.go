package main

import (
	"fmt"
)

func foldYAxis(manual [][]bool, x int) (newManual [][]bool) {
	newManual = make([][]bool, len(manual))

	for i := range manual {
		newManualLine := make([]bool, len(manual[0])/2)

		for j := 0; j < len(manual[i])/2; j++ {
			newManualLine[j] = manual[i][j] || manual[i][len(manual[i])-1-j]
		}

		newManual[i] = newManualLine
	}
	return
}

func foldXAxis(manual [][]bool, y int) (newManual [][]bool) {
	newManual = manual[0:y]
	for i := range newManual {
		for j := range newManual[i] {
			newManual[i][j] = newManual[i][j] || manual[len(manual)-1-i][j]
		}
	}
	return
}

func countMarkedDots(manual [][]bool) (numDots int) {
	for _, line := range manual {
		for _, val := range line {
			if val {
				numDots++
			}
		}
	}
	return numDots
}

func firstPart(manual [][]bool, foldingInstructions []instruction) {
	fmt.Println("===================")
	fmt.Println("Starting first part")
	fmt.Println("===================")

	newManual := manual

	for _, inst := range foldingInstructions {
		switch inst.axis {
		case "x":
			newManual = foldYAxis(newManual, inst.position)
		case "y":
			newManual = foldXAxis(newManual, inst.position)
		}

		// Stop at the first instruction
		break
	}

	numDots := countMarkedDots(newManual)

	fmt.Println("Number of Dots after the first fold", numDots)
}
