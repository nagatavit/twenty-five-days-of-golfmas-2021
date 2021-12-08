package main

import (
	"bufio"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func readInput() (crabsPosition []int) {
	f, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	scanner.Scan()
	crabsInitialPositionStr := strings.Split(scanner.Text(), ",")

	crabsPosition = make([]int, len(crabsInitialPositionStr))

	for i, crabsPosStr := range crabsInitialPositionStr {
		crabsPosition[i], err = strconv.Atoi(crabsPosStr)
		if err != nil {
			log.Fatal(err)
		}
	}

	sort.Ints(crabsPosition)

	return crabsPosition
}

func main() {
	crabsInitialPosition := readInput()
	firstPart(crabsInitialPosition)
}
