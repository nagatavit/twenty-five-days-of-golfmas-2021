package main

import (
	"fmt"
)

type depthSlice struct {
	depth int
	value int
}

func snailFishTreeToDepthSlice(node *snailfishNumber, depth int) (result []depthSlice) {
	if node == nil {
		return
	}

	result = append(result, snailFishTreeToDepthSlice(node.leftPair, depth+1)...)

	if node.isLeaf() {
		result = append(result,
			depthSlice{
				depth: depth,
				value: node.value,
			},
		)
	}

	result = append(result, snailFishTreeToDepthSlice(node.rightPair, depth+1)...)

	return result
}

func sumSnailfishNumber(left, right []depthSlice) (result []depthSlice) {
	result = append(left, right...)

	for i := 0; i < len(result); i++ {
		result[i].depth++
	}

	return result
}

func explode(number []depthSlice) (result []depthSlice, exploded bool) {
	var leftExploded int
	var rightExploded int

	var explodePosition int

	for i := range number {
		if number[i].depth > 4 {
			explodePosition = i
			leftExploded = number[i].value
			rightExploded = number[i+1].value
			exploded = true
			break
		}
	}

	if exploded {
		// Add the left explosion
		if explodePosition-1 >= 0 {
			number[explodePosition-1].value += leftExploded
		}

		// Add the right explosion
		if explodePosition+2 < len(number) {
			number[explodePosition+2].value += rightExploded
		}

		number[explodePosition+1].depth--
		number[explodePosition+1].value = 0

		number = append(number[:explodePosition], number[explodePosition+1:]...)
	}

	return number, exploded
}

func split(number []depthSlice) (result []depthSlice, splitted bool) {
	var splitPosition int

	for i := range number {
		if number[i].value > 9 {
			splitPosition = i
			splitted = true
			break
		}
	}

	if splitted {
		value := number[splitPosition].value

		number[splitPosition].value = value / 2
		number[splitPosition].depth++

		newNumber := depthSlice{
			depth: number[splitPosition].depth,
			value: value/2 + value&0x1,
		}

		number = append(number[:splitPosition+1], append([]depthSlice{newNumber}, number[splitPosition+1:]...)...)
	}

	return number, splitted
}

func reduceSum(number []depthSlice) []depthSlice {
	isReduced := false

	var hasExploded bool
	var hasSplitted bool

	for !isReduced {
		isReduced = true

		number, hasExploded = explode(number)
		if hasExploded {
			isReduced = false
		} else {
			number, hasSplitted = split(number)
			if hasSplitted {
				isReduced = false
			}
		}
	}

	return number
}

func printDepthList(number []depthSlice) {
	depth := 0
	printed := 0

	for printed < len(number) {
		fmt.Print(depth, ": ")
		for _, depthSlice := range number {
			if depthSlice.depth == depth {
				fmt.Print(depthSlice.value)
				printed++
			} else {
				fmt.Print(" ")
			}

			fmt.Print(" ")
		}
		depth++
		fmt.Println()
	}
}

func magnetude(number []depthSlice) int {
	var magnetudeStack stack

	for _, element := range number {
		if magnetudeStack.Size() < 1 {
			magnetudeStack.Push(element)
		} else {
			magnetudeStack.Push(element)

			var partialResult depthSlice

			right := (magnetudeStack.Pop()).(depthSlice)
			left := (magnetudeStack.Pop()).(depthSlice)

			// Check if they are in the same level, if not, push them back to the stack
			if left.depth != right.depth {
				magnetudeStack.Push(left)
				magnetudeStack.Push(right)

			} else {
				partialResult.depth = left.depth - 1
				partialResult.value = 3*left.value + 2*right.value

				magnetudeStack.Push(partialResult)
			}
		}

		// simplify stack
		isStackSimplified := false
		var partialResult depthSlice
		for !isStackSimplified {
			if magnetudeStack.Size() >= 2 {
				right := (magnetudeStack.Pop()).(depthSlice)
				left := (magnetudeStack.Pop()).(depthSlice)

				if left.depth != right.depth {
					magnetudeStack.Push(left)
					magnetudeStack.Push(right)
					isStackSimplified = true
				} else {
					partialResult.depth = left.depth - 1
					partialResult.value = 3*left.value + 2*right.value
					magnetudeStack.Push(partialResult)
				}

			} else {
				isStackSimplified = true
			}
		}
	}

	return (magnetudeStack.Pop()).(depthSlice).value
}

func firstPart(homework []snailfishNumber) {
	fmt.Println("===================")
	fmt.Println("Starting first part")
	fmt.Println("===================")

	var homeworkButInAList [][]depthSlice

	for _, number := range homework {
		homeworkButInAList = append(homeworkButInAList, snailFishTreeToDepthSlice(&number, 0))
	}

	partialResult := homeworkButInAList[0]

	for i := 1; i < len(homeworkButInAList); i++ {
		partialResult = sumSnailfishNumber(partialResult, homeworkButInAList[i])

		partialResult = reduceSum(partialResult)
	}

	fmt.Println(magnetude(partialResult))
}
