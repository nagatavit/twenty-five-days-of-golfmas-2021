package main

import (
	"fmt"
	"log"
	"strings"
)

type sevenSegmentDisplay struct {
	top                     string
	topLeft, topRight       string
	middle                  string
	bottomLeft, bottomRight string
	bottom                  string
}

func secondPart(entries []signalNote) {
	fmt.Println("====================")
	fmt.Println("Starting second part")
	fmt.Println("====================")

	totalSum := 0

	for _, entry := range entries {
		displayEdges := decodeSignals(entry.signals)
		totalSum += decodeDigits(entry.digits, displayEdges)
	}

	fmt.Println("total sum", totalSum)
}

///////////////////////////////////////////////////////////////////////////////
//                               Decode digits                               //
///////////////////////////////////////////////////////////////////////////////

func decodeDigits(digits [4]string, displayEdges sevenSegmentDisplay) (result int) {
	for i := 0; i < 4; i++ {
		result = result*10 + decodeSingleDigit(digits[i], displayEdges)
	}
	return result
}

func decodeSingleDigit(digit string, displayEdges sevenSegmentDisplay) int {
	// Check for the easier cases
	switch len(digit) {
	case 2:
		return 1
	case 3:
		return 7
	case 4:
		return 4
	case 7:
		return 8
	}

	// Bottom left edge divides the group into two:
	// 0, 2, 6 vs 3, 5, 9

	// 0, 2, 6
	if strings.ContainsAny(digit, displayEdges.bottomLeft) {
		// 2 and 6
		if strings.ContainsAny(digit, displayEdges.middle) {
			if strings.ContainsAny(digit, displayEdges.topLeft) {
				return 6
			} else {
				return 2
			}
		} else {
			return 0
		}
		// 3, 5, 9
	} else {
		// 3 and 9
		if strings.ContainsAny(digit, displayEdges.topRight) {
			if strings.ContainsAny(digit, displayEdges.topLeft) {
				return 9
			} else {
				return 3
			}
		} else {
			return 5
		}
	}
}

///////////////////////////////////////////////////////////////////////////////
//                               Decode signals                              //
///////////////////////////////////////////////////////////////////////////////

func decodeSignals(signals [10]string) (displayEdges sevenSegmentDisplay) {
	var numbersWeFound [10]string

	var candidatesWithFiveEdges []string
	var candidatesWithSixEdges []string

	// Found easy ones first
	for _, easyDigit := range signals {
		switch len(easyDigit) {
		case 2:
			numbersWeFound[1] = easyDigit
		case 3:
			numbersWeFound[7] = easyDigit
		case 4:
			numbersWeFound[4] = easyDigit
		case 7:
			numbersWeFound[8] = easyDigit
		case 5:
			candidatesWithFiveEdges = append(candidatesWithFiveEdges, easyDigit)
		case 6:
			candidatesWithSixEdges = append(candidatesWithSixEdges, easyDigit)
		default:
			// Shouldn't happen
			log.Fatal("How did you even ended up here?")
		}
	}

	displayEdges.top = getTopEdge(numbersWeFound[1], numbersWeFound[7])

	displayEdges.bottom, candidatesWithSixEdges, numbersWeFound[9] = getBottomEdge(displayEdges.top,
		numbersWeFound[4], candidatesWithSixEdges)

	displayEdges.middle, _, numbersWeFound[3] = getMiddleEdge(displayEdges.bottom, numbersWeFound[7],
		candidatesWithFiveEdges)

	displayEdges.bottomRight, numbersWeFound[6], numbersWeFound[0] = getBottomRightEdge(displayEdges.middle,
		numbersWeFound[1], candidatesWithSixEdges)

	displayEdges.topRight = subEdges(numbersWeFound[1], displayEdges.bottomRight)

	displayEdges.topLeft = getTopLeftEdge(displayEdges.middle, numbersWeFound[1], numbersWeFound[4])

	displayEdges.bottomLeft = getBottomLeftEdge(numbersWeFound[8], numbersWeFound[9])

	return displayEdges
}

///////////////////////////////////////////////////////////////////////////////
//                            Get edges functions                            //
///////////////////////////////////////////////////////////////////////////////

func getTopEdge(one, seven string) string {
	return subEdges(seven, one)
}

func getBottomEdge(topEdge, four string, candidatesWithSixEdges []string) (bottomEdge string,
	newCandidatesWithSixEdges []string, nine string) {

	nineWithoutBottom := unionEdges(four, topEdge)

	nineIndex := 0

	for i, candidate := range candidatesWithSixEdges {
		bottomEdge = subEdges(candidate, nineWithoutBottom)
		if len(bottomEdge) == 1 {
			nineIndex = i
			break
		}
	}

	newCandidatesWithSixEdges = removeCandidateFromList(candidatesWithSixEdges, nineIndex)

	return bottomEdge, newCandidatesWithSixEdges, candidatesWithSixEdges[nineIndex]
}

func getMiddleEdge(bottomEdge, seven string, candidatesWithFiveEdges []string) (middleEdge string,
	newCandidatesWithFiveEdges []string, three string) {

	invertedC := unionEdges(seven, bottomEdge)

	threeIndex := 0

	for i, candidate := range candidatesWithFiveEdges {
		bottomEdge = subEdges(candidate, invertedC)
		if len(bottomEdge) == 1 {
			threeIndex = i
			break
		}
	}

	newCandidatesWithFiveEdges = removeCandidateFromList(candidatesWithFiveEdges, threeIndex)

	return bottomEdge, newCandidatesWithFiveEdges, candidatesWithFiveEdges[threeIndex]
}

// Only call this after 9 is already out, or else it will not work
func getBottomRightEdge(middleEdge, one string, candidatesWithSixEdges []string) (bottomRightEdge string,
	six, zero string) {

	sixIndex := 0

	// Assuming 9 is already out of the candidatesWithSixEdges, the only
	// one with a middleEdge is the 6
	for i, candidate := range candidatesWithSixEdges {
		if strings.ContainsAny(candidate, middleEdge) {
			sixIndex = i
			break
		}
	}

	bottomRightEdge = intersectionEdges(candidatesWithSixEdges[sixIndex], one)

	// sixIndex can only be 0 or 1 and the zero index will be the
	// remaining one. XORing the value with 0x1 will invert the value
	// from 0 to 1 or 1 to 0, giving us the remaining index for the
	// zero candidate.
	return bottomRightEdge, candidatesWithSixEdges[sixIndex], candidatesWithSixEdges[sixIndex^0x1]
}

func getTopLeftEdge(middleEdge, one, four string) (topLeftEdge string) {
	sidewaysT := unionEdges(one, middleEdge)
	return subEdges(four, sidewaysT)
}

func getBottomLeftEdge(eight, nine string) (bottomLeftEdge string) {
	return subEdges(eight, nine)
}

///////////////////////////////////////////////////////////////////////////////
//                            Auxiliary functions                            //
///////////////////////////////////////////////////////////////////////////////

func removeCandidateFromList(old []string, idx int) (new []string) {
	for i, candidate := range old {
		if i != idx {
			new = append(new, candidate)
		}
	}
	return new
}

func subEdges(original, sub string) (result string) {
L:
	for _, originalEdge := range original {
		for _, subEdge := range sub {
			if originalEdge == subEdge {
				continue L
			}
		}
		result += string(originalEdge)
	}
	return result
}

func unionEdges(original, appendable string) (result string) {
	result = original
L:
	for _, appendEdge := range appendable {
		for _, originalEdge := range original {
			if originalEdge == appendEdge {
				continue L
			}
		}
		result += string(appendEdge)
	}
	return result
}

func intersectionEdges(original, filter string) (result string) {
	for _, originalEdge := range original {
		for _, filterEdge := range filter {
			if originalEdge == filterEdge {
				result += string(filterEdge)
			}
		}
	}
	return result
}
