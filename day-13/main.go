package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type instruction struct {
	axis     string
	position int
}

func readInput() (manual [][]bool, foldingInstructions []instruction) {
	// f, err := os.Open("example-input")
	f, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	var dotsCoord [][2]int
	var biggestXValue, biggestYValue int

	// Manual dots ////////////////////////////////////////////////////////////

	// Scan dots coordinates
	for scanner.Scan() {
		lineStr := scanner.Text()

		// Newline that separates the coords from the folding
		// instructions
		if lineStr == "" {
			break
		}

		var newCoord [2]int

		_, err := fmt.Sscanf(lineStr, "%d,%d", &newCoord[0], &newCoord[1])
		if err != nil {
			log.Fatal(err)
		}

		if newCoord[0] > biggestXValue {
			biggestXValue = newCoord[0]
		}

		if newCoord[1] > biggestYValue {
			biggestYValue = newCoord[1]
		}

		dotsCoord = append(dotsCoord, newCoord)
	}

	// Allocate the "manual"
	manual = make([][]bool, biggestYValue+1)
	for i := range manual {
		manual[i] = make([]bool, biggestXValue+1)
	}

	// Fill the dots on the manual
	for _, coord := range dotsCoord {
		manual[coord[1]][coord[0]] = true
	}

	// Folding instructions ///////////////////////////////////////////////////

	// Scan the folding instructions
	for scanner.Scan() {
		lineStr := scanner.Text()

		var newInstruction instruction

		var instStr string

		fmt.Sscanf(lineStr, "fold along %s", &instStr)

		axisAndVal := strings.Split(instStr, "=")

		newInstruction.axis = axisAndVal[0]
		newInstruction.position, err = strconv.Atoi(axisAndVal[1])
		if err != nil {
			log.Fatal(err)
		}

		foldingInstructions = append(foldingInstructions, newInstruction)
	}

	return manual, foldingInstructions
}

func printManual(manual [][]bool) {
	for _, line := range manual {
		for _, val := range line {
			if val {
				fmt.Print("█")
			} else {
				fmt.Print("□")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

func main() {
	// Add location (line and file) of log calls
	log.SetFlags(log.Lshortfile)

	manual, foldingInstructions := readInput()
	firstPart(manual, foldingInstructions)
	secondPart(manual, foldingInstructions)
}
