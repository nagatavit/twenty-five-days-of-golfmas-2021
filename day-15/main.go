package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func readInput() (cavernMap [][]int) {
	// f, err := os.Open("example-input")
	f, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		lineStr := scanner.Text()

		newLine := make([]int, len(lineStr))

		for i, heightRune := range lineStr {
			newLine[i], err = strconv.Atoi(string(heightRune))
			if err != nil {
				log.Fatal(err)
			}
		}

		cavernMap = append(cavernMap, newLine)
	}

	return cavernMap
}

func main() {
	// Add location (line and file) of log calls
	log.SetFlags(log.Lshortfile)

	cavernMap := readInput()

	// Putting the prints over here because I will call the first part
	// in the second one (yes, I'm too lazy to encapsulate all of that)
	fmt.Println("===================")
	fmt.Println("Starting first part")
	fmt.Println("===================")
	firstPart(cavernMap)

	fmt.Println("===================")
	fmt.Println("Starting second part")
	fmt.Println("===================")
	secondPart(cavernMap)
}
