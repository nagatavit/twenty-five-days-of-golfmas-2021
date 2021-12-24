package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

type snailfishNumber struct {
	value     int
	leftPair  *snailfishNumber
	rightPair *snailfishNumber
}

type regularNumberType int

func readInput() (homework []snailfishNumber) {
	// f, err := os.Open("example-input-3")
	f, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		lineStr := scanner.Text()

		var inputStack stack

		for _, lineElem := range lineStr {
			lineElem := string(lineElem)

			if lineElem != "]" {
				inputStack.Push(lineElem)

			} else {
				var snailNumberList []interface{}

				for {
					poppedElement := inputStack.Pop()

					poppedElementStr, ok := poppedElement.(string)
					if ok && poppedElementStr == "[" {
						break
					}

					snailNumberList = append([]interface{}{poppedElement}, snailNumberList...)
				}

				var newSnailfishNumber snailfishNumber

				var leftPair snailfishNumber
				var rightPair snailfishNumber

				switch snailNumber := snailNumberList[0].(type) {
				case string:
					leftPair.value, err = strconv.Atoi(snailNumber)
					if err != nil {
						log.Fatal(err)
					}
				case snailfishNumber:
					leftPair = snailNumber
				}

				switch snailNumber := snailNumberList[2].(type) {
				case string:
					rightPair.value, err = strconv.Atoi(snailNumber)
					if err != nil {
						log.Fatal(err)
					}
				case snailfishNumber:
					rightPair = snailNumber
				}

				newSnailfishNumber.leftPair = &leftPair
				newSnailfishNumber.rightPair = &rightPair

				inputStack.Push(newSnailfishNumber)
			}
		}

		fullLineParsed := inputStack.Pop()

		homework = append(homework, fullLineParsed.(snailfishNumber))
	}

	return
}

func (number snailfishNumber) isLeaf() bool {
	if number.leftPair == nil && number.rightPair == nil {
		return true
	}
	return false
}

func printInorderTree(root *snailfishNumber) {
	if root == nil {
		return
	}

	printInorderTree(root.leftPair)

	if root.isLeaf() {
		fmt.Println(root.value)
	}

	printInorderTree(root.rightPair)
}

func main() {
	// Add location (line and file) of log calls
	log.SetFlags(log.Lshortfile)

	homework := readInput()
	firstPart(homework)
	secondPart(homework)
}
