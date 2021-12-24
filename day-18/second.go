package main

import (
	"fmt"
)

func secondPart(homework []snailfishNumber) {
	fmt.Println("====================")
	fmt.Println("Starting second part")
	fmt.Println("====================")

	var homeworkButInAList [][]depthSlice

	for _, number := range homework {
		homeworkButInAList = append(homeworkButInAList, snailFishTreeToDepthSlice(&number, 0))
	}

	biggestMagnetude := 0

	for i, leftSide := range homeworkButInAList {
		for j, rightSide := range homeworkButInAList {
			if i == j {
				continue
			}

			// Okay, so I forgot about a little big thing in go. I
			// always remembered that in go, everything is passed by
			// value, but forgot the important part that slices are
			// just a header indicator to an array D:. So
			// modifications to the slice will affect the underlying
			// array. I'm too overdue with this challenge, so what I'm
			// just going to do is copy the slice to a new one on
			// every iteration.
			copyOfLeftSide := make([]depthSlice, len(leftSide))
			copyOfRightSide := make([]depthSlice, len(rightSide))

			copy(copyOfLeftSide, leftSide)
			copy(copyOfRightSide, rightSide)

			var partialResult []depthSlice
			partialResult = sumSnailfishNumber(copyOfLeftSide, copyOfRightSide)
			partialResult = reduceSum(partialResult)
			newMagnetude := magnetude(partialResult)

			if newMagnetude > biggestMagnetude {
				biggestMagnetude = newMagnetude
			}
		}
	}

	fmt.Println(biggestMagnetude)
}
