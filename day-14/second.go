package main

import (
	"fmt"
	"math"
)

func findCommonAndUncommonElementsOnPairs(polyPairs map[string]int, firstElement, lastElement string) (mostCommon, leastCommon int) {
	elementCount := make(map[string]int)

	for pair, count := range polyPairs {
		for _, element := range pair {
			elementCount[string(element)] += count
		}
	}

	// Add the first and last element to the total sum
	elementCount[firstElement]++
	elementCount[lastElement]++

	mostCommon = 0
	leastCommon = math.MaxInt

	for _, count := range elementCount {
		if count > mostCommon {
			mostCommon = count
		} else if count < leastCommon {
			leastCommon = count
		}
	}

	return mostCommon, leastCommon
}

func secondPart(polymerTemplate string, polyMap map[string]string) {
	fmt.Println("====================")
	fmt.Println("Starting second part")
	fmt.Println("====================")

	fmt.Println("polymerTemplate", polymerTemplate)

	// We are going to use a map of pairs, but because of that, each
	// element will be counted twice with the exception of the first
	// and last elements. The good thing is, because we are always
	// "inserting" into pairs, the first and last elements will always be the same
	firstElement := string(polymerTemplate[0])
	lastElement := string(polymerTemplate[len(polymerTemplate)-1])

	polymerPairs := make(map[string]int)

	// Initialize the map
	for i := range polymerTemplate {
		if i+1 > len(polymerTemplate)-1 {
			break
		}

		polyPair := string(polymerTemplate[i]) + string(polymerTemplate[i+1])

		if _, ok := polymerPairs[polyPair]; !ok {
			polymerPairs[polyPair] = 1
		} else {
			polymerPairs[polyPair]++
		}
	}

	fmt.Println(polymerPairs)

	iteractions := 40

	for i := 0; i < iteractions; i++ {
		newPolymerPairs := make(map[string]int)

		for polyPair, count := range polymerPairs {
			newInsertion, ok := polyMap[polyPair]
			if ok {
				firstPair := string(polyPair[0]) + newInsertion
				secondPair := newInsertion + string(polyPair[1])

				newPolymerPairs[firstPair] += count
				newPolymerPairs[secondPair] += count
			} else {
				newPolymerPairs[polyPair] += count
			}
		}

		polymerPairs = newPolymerPairs
	}

	mostCommon, leastCommon := findCommonAndUncommonElementsOnPairs(polymerPairs, firstElement, lastElement)
	fmt.Println("Most Common:", mostCommon/2)
	fmt.Println("Least Common:", leastCommon/2)
	fmt.Println("Diff:", mostCommon/2-leastCommon/2)
}
