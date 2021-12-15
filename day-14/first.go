package main

import (
	"fmt"
)

func findCommonAndUncommonElements(polymerTemplate string) (mostCommon, leastCommon int) {
	elementCount := make(map[string]int)

	for _, element := range polymerTemplate {
		if count, ok := elementCount[string(element)]; !ok {
			elementCount[string(element)] = 1
		} else {
			elementCount[string(element)] = count + 1
		}
	}

	mostCommon = 0
	leastCommon = len(polymerTemplate)

	for _, count := range elementCount {
		if count > mostCommon {
			mostCommon = count
		}

		if count < leastCommon {
			leastCommon = count
		}
	}

	return mostCommon, leastCommon
}

func firstPart(polymerTemplate string, polyMap map[string]string) {
	fmt.Println("===================")
	fmt.Println("Starting first part")
	fmt.Println("===================")

	newPolymer := ""

	iteractions := 10

	fmt.Println("polymerTemplate", polymerTemplate)

	for i := 0; i < iteractions; i++ {
		for i := range polymerTemplate {
			newPolymer += string(polymerTemplate[i])

			if i+1 > len(polymerTemplate)-1 {
				break
			}

			oldSeq := string(polymerTemplate[i]) + string(polymerTemplate[i+1])

			newInsertion, ok := polyMap[oldSeq]
			if ok {
				newPolymer += newInsertion
			}
		}

		polymerTemplate = newPolymer
		newPolymer = ""
	}

	mostCommon, leastCommon := findCommonAndUncommonElements(polymerTemplate)
	fmt.Println("Most Common:", mostCommon)
	fmt.Println("Least Common:", leastCommon)
	fmt.Println("Diff:", mostCommon-leastCommon)
}
