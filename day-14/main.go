package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type insertionRule struct {
	basePair  string
	insertion string
}

func readInput() (polymerTemplate string, polyMap map[string]string) {
	// f, err := os.Open("example-input")
	f, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	// Scan the first line
	scanner.Scan()
	polymerTemplate = scanner.Text()

	// Just to skip the second line
	scanner.Scan()

	polyMap = make(map[string]string)

	for scanner.Scan() {
		var newRule insertionRule

		fmt.Sscanf(scanner.Text(), "%s -> %s", &newRule.basePair, &newRule.insertion)

		polyMap[newRule.basePair] = newRule.insertion
	}

	return polymerTemplate, polyMap
}

func printPolymerRules(rules []insertionRule) {
	for _, rule := range rules {
		fmt.Printf("%v\n", rule)
	}
}

func main() {
	// Add location (line and file) of log calls
	log.SetFlags(log.Lshortfile)

	polymerTemplate, polyMap := readInput()
	firstPart(polymerTemplate, polyMap)
	secondPart(polymerTemplate, polyMap)
}
