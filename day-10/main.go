package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func readInput() (navigationSubsystem [][]string) {
	// f, err := os.Open("example-input")
	f, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		lineStr := scanner.Text()

		newLine := make([]string, len(lineStr))

		for i, navSubRune := range lineStr {
			newLine[i] = string(navSubRune)
		}

		navigationSubsystem = append(navigationSubsystem, newLine)
	}

	return navigationSubsystem
}

func printNavSubsystem(navigationSubSystem [][]string) {
	for _, line := range navigationSubSystem {
		fmt.Println(line)
	}
}

var isOpenChar = map[string]bool{
	"(": true,
	"[": true,
	"{": true,
	"<": true,
	")": false,
	"]": false,
	"}": false,
	">": false,
}

var closingPairs = map[string]string{
	"(": ")",
	"[": "]",
	"{": "}",
	"<": ">",
}

var illegalCharPontuation = map[string]int{
	")": 3,
	"]": 57,
	"}": 1197,
	">": 25137,
}

var incompleteCharPontuation = map[string]int{
	")": 1,
	"]": 2,
	"}": 3,
	">": 4,
}

func main() {
	navigationSubsystem := readInput()
	firstPart(navigationSubsystem)
	secondPart(navigationSubsystem)
}
