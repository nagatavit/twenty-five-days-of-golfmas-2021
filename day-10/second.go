package main

import (
	"fmt"
	"sort"
)

// Assuming here that the line is always just incomplete (not
// corrupted or strange in some other way)
func checkAutocompletePattern(line []string) (completionChars []string) {
	var navSubStack stack

	for _, char := range line {
		if isOpenChar[char] {
			navSubStack.Push(char)
		} else {
			navSubStack.Pop()
		}
	}

	remainingStackSize := navSubStack.Size()

	for i := 0; i < remainingStackSize; i++ {
		incompleteOpenChar := navSubStack.Pop()
		completionChars = append(completionChars, closingPairs[incompleteOpenChar])
	}

	return completionChars
}

func secondPart(navigationSubsystem [][]string) {
	fmt.Println("====================")
	fmt.Println("Starting second part")
	fmt.Println("====================")

	var autocompleteScores []int

	for _, line := range navigationSubsystem {
		autocompleteCharSum := 0

		isCorrupted, _ := checkIfLineIsCorrupted(line)
		if !isCorrupted {
			completionChars := checkAutocompletePattern(line)

			for _, completionChar := range completionChars {
				autocompleteCharSum *= 5
				autocompleteCharSum += incompleteCharPontuation[completionChar]
			}

			autocompleteScores = append(autocompleteScores, autocompleteCharSum)
		}
	}

	sort.Ints(autocompleteScores)

	fmt.Println(autocompleteScores)

	fmt.Println(autocompleteScores[len(autocompleteScores)/2])
}
