package main

import "fmt"

func secondPart(numbersDrawn []int, boards []Board) {
	fmt.Println("====================")
	fmt.Println("Starting second part")
	fmt.Println("====================")

	fmt.Println("numbersDrawn", numbersDrawn)

	totalBoards := len(boards)
	bingoBoards := 0

L:
	for _, currNumber := range numbersDrawn {
		for j := range boards {
			if boards[j].Bingo {
				continue
			}

			if fillBoardNumber(currNumber, &boards[j]) {
				boards[j].Bingo = true
				bingoBoards++

				if bingoBoards == totalBoards {
					sumUnusedNumbers := sumUnusedNumbers(boards[j])
					fmt.Println("currNumber", currNumber)
					fmt.Println("sumUnusedNumbers", sumUnusedNumbers)
					fmt.Println("sumUnusedNumbers * currNumber", sumUnusedNumbers*currNumber)

					break L
				}
			}
		}
	}

}
