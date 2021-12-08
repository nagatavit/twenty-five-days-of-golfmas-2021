package main

import (
	"fmt"
)

func firstPart(numbersDrawn []int, boards []Board) {
	fmt.Println("===================")
	fmt.Println("Starting first part")
	fmt.Println("===================")

	fmt.Println("numbersDrawn", numbersDrawn)

L:
	for _, currNumber := range numbersDrawn {
		for j := range boards {
			if fillBoardNumber(currNumber, &boards[j]) {
				sumUnusedNumbers := sumUnusedNumbers(boards[j])
				fmt.Println("currNumber", currNumber)
				fmt.Println("sumUnusedNumbers", sumUnusedNumbers)
				fmt.Println("sumUnusedNumbers * currNumber", sumUnusedNumbers*currNumber)
				break L
			}
		}
	}
}
