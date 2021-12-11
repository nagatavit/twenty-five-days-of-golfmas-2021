package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func readInput() (heightmap [][]int) {
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

		heightmap = append(heightmap, newLine)
	}

	return heightmap
}

func main() {
	heightmap := readInput()
	firstPart(heightmap)
	secondPart(heightmap)
}
