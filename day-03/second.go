package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
)

func readInput() (bitLen int, sortedArray []int) {
	f, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	// read the first one to get the len
	scanner.Scan()
	firstDiag := scanner.Text()
	sortedArray = append(sortedArray, bitStringToDec(firstDiag))
	bitLen = len(firstDiag) - 1

	// This time, let's work with real bits, not just representations
	// of ones :D
	for scanner.Scan() {
		sortedArray = append(sortedArray, bitStringToDec(scanner.Text()))
	}

	sort.Ints(sortedArray)

	return bitLen, sortedArray
}

func secondPart() {
	fmt.Println("===================")
	fmt.Println("Running second part")

	bitLen, sortedInput := readInput()

	var oxygenGeneratorRating int
	var co2ScrubberRating int

	oxygenCandidates := sortedInput
	co2Candidates := sortedInput

	// Because the array is sorted, reading from left to right, the
	// bits are going to be separated in two ranges, one of zeros and
	// one with ones. We just need to find where this threshold lies
	// and calculate the most and least common bit.
	for i := bitLen; i >= 0; i-- {
		rangeDivisor := 1 << i

		oxygenCandidates, oxygenGeneratorRating = getLowerAndUpperBounds(oxygenCandidates, oxygenGeneratorRating, rangeDivisor, true)

		co2Candidates, co2ScrubberRating = getLowerAndUpperBounds(co2Candidates, co2ScrubberRating, rangeDivisor, false)
	}

	fmt.Println("oxygenGeneratorRating", oxygenGeneratorRating)
	fmt.Println("co2ScrubberRating", co2ScrubberRating)
	fmt.Println("oxygenGeneratorRating * co2ScrubberRating", oxygenGeneratorRating*co2ScrubberRating)
}

func getLowerAndUpperBounds(input []int, currentRating int, rangeDivisor int, getMoreFrequent bool) ([]int, int) {
	// Check if there's already only one solution
	if len(input) == 1 {
		return input, input[0]
	}

	var thresholdIndex int

	for i, currReading := range input {
		if currReading >= (currentRating | rangeDivisor) {
			thresholdIndex = i
			break
		}
	}

	// check if the zeros is greater or ones
	zeroRange := thresholdIndex

	// +1 to include the current value (included on len())
	oneRange := len(input) - thresholdIndex

	if getMoreFrequent {
		if oneRange >= zeroRange {
			input = input[thresholdIndex:]
			currentRating |= rangeDivisor
		} else {
			input = input[:thresholdIndex]
		}
	} else {
		if oneRange >= zeroRange {
			input = input[:thresholdIndex]
		} else {
			input = input[thresholdIndex:]
			currentRating |= rangeDivisor
		}
	}

	return input, currentRating
}
