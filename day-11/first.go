package main

import (
	"fmt"
)

func flash(x, y int, octoMap *[][]dumboOctopus) {
	if (*octoMap)[x][y].hasFlashed {
		return
	}

	(*octoMap)[x][y].hasFlashed = true

L:
	for i := x - 1; i <= x+1; i++ {
		for j := y - 1; j <= y+1; j++ {
			if i < 0 || i > len(*octoMap)-1 {
				// jmp to outer i loop
				continue L
			} else if j < 0 || j > len((*octoMap)[i])-1 {
				// jmp to inner j loop
				continue
			} else {
				(*octoMap)[i][j].energy++
				if (*octoMap)[i][j].energy > 9 {
					flash(i, j, octoMap)
				}
			}
		}
	}
}

func resetStepFlashes(octoMap *[][]dumboOctopus) (numberOfFlashes int) {
	for i, line := range *octoMap {
		for j := range line {
			if (*octoMap)[i][j].energy > 9 {
				numberOfFlashes++
				(*octoMap)[i][j].energy = 0
				(*octoMap)[i][j].hasFlashed = false
			}
		}
	}
	return numberOfFlashes
}

func newStep(octoMap *[][]dumboOctopus) (numberOfFlashes int) {
	for i, line := range *octoMap {
		for j := range line {
			(*octoMap)[i][j].energy++

			if (*octoMap)[i][j].energy > 9 {
				flash(i, j, octoMap)
			}
		}
	}
	return resetStepFlashes(octoMap)
}

func firstPart(octoMap [][]dumboOctopus) {
	fmt.Println("===================")
	fmt.Println("Starting first part")
	fmt.Println("===================")

	numberOfSteps := 100
	totalNumberOfFlashes := 0

	for i := 0; i < numberOfSteps; i++ {
		totalNumberOfFlashes += newStep(&octoMap)
	}

	fmt.Println("number of steps", numberOfSteps)
	fmt.Println("number of flashes", totalNumberOfFlashes)
}
