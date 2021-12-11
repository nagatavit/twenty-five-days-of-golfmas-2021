package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

type dumboOctopus struct {
	energy     int
	hasFlashed bool
}

func readInput() (octoMap [][]dumboOctopus) {
	// f, err := os.Open("example-input")
	f, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		lineStr := scanner.Text()

		newLine := make([]dumboOctopus, len(lineStr))

		var err error

		for i, octopusRune := range lineStr {
			newLine[i].energy, err = strconv.Atoi(string(octopusRune))
			if err != nil {
				log.Fatal(err)
			}
		}

		octoMap = append(octoMap, newLine)
	}

	return octoMap
}

func printOctoMap(octoMap [][]dumboOctopus) {
	for _, line := range octoMap {
		for _, octopus := range line {
			fmt.Print(octopus.energy)
		}
		fmt.Println()
	}
}

func printOctoMapFlash(octoMap [][]dumboOctopus) {
	for _, line := range octoMap {
		for _, octopus := range line {
			if octopus.hasFlashed {
				fmt.Print(1)
			} else {
				fmt.Print(0)
			}
		}
		fmt.Println()
	}
}

func main() {
	octoMap := readInput()
	firstPart(octoMap)

	// Reread input as we modified it in the first part
	octoMap = readInput()
	secondPart(octoMap)
}
