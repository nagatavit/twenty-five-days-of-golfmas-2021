package main

import (
	"fmt"
)

func checkIfLineIsCorrupted(line []string) (isCorrupted bool, illegalChar string) {
	var navSubStack stack

	for _, char := range line {
		if isOpenChar[char] {
			navSubStack.Push(char)
		} else {
			openChar := navSubStack.Pop()
			if closingPairs[openChar] != char {
				return true, char
			}
		}
	}

	return false, ""
}

func firstPart(navigationSubsystem [][]string) {
	fmt.Println("===================")
	fmt.Println("Starting first part")
	fmt.Println("===================")

	illegalCharSum := 0

	for _, line := range navigationSubsystem {
		isCorrupted, illegalChar := checkIfLineIsCorrupted(line)
		if isCorrupted {
			illegalCharSum += illegalCharPontuation[illegalChar]
		}
	}

	fmt.Println(illegalCharSum)
}
