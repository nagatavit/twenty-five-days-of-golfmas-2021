package main

import (
	"fmt"
	"sort"
)

func exploreBasil(i, j int, heightmap *[][]int, visitedMap *[][]bool) (numVisited int) {
	// Not even sure if we need this second verification, but putting
	// it here just to be sure
	if (*visitedMap)[i][j] || (*heightmap)[i][j] < 0 {
		return 0
	}

	// Mark as visited first
	(*visitedMap)[i][j] = true
	numVisited += 1

	top, right, bottom, left := getSides(i, j, heightmap)
	if top > (*heightmap)[i][j] {
		numVisited += exploreBasil(i-1, j, heightmap, visitedMap)
	}

	if right > (*heightmap)[i][j] {
		numVisited += exploreBasil(i, j+1, heightmap, visitedMap)
	}

	if bottom > (*heightmap)[i][j] {
		numVisited += exploreBasil(i+1, j, heightmap, visitedMap)
	}

	if left > (*heightmap)[i][j] {
		numVisited += exploreBasil(i, j-1, heightmap, visitedMap)
	}

	return numVisited
}

func printVisitedMap(visitedMap [][]bool) {
	for _, line := range visitedMap {
		for _, val := range line {
			if val {
				fmt.Print(1)
			} else {
				fmt.Print(0)
			}
		}
		fmt.Println()
	}
}

func printHeightMap(heightMap [][]int) {
	for _, line := range heightMap {
		fmt.Println(line)
	}
}

func secondPart(heightmap [][]int) {
	fmt.Println("====================")
	fmt.Println("Starting second part")
	fmt.Println("====================")

	var visitedMap [][]bool

	// Pre fill 9s as visited already (basil borders)
	for _, line := range heightmap {
		newLine := make([]bool, len(line))

		for j, val := range line {
			if val == 9 {
				newLine[j] = true
			}
		}

		visitedMap = append(visitedMap, newLine)
	}

	var topBasins []int

	for i, line := range heightmap {
		for j, middle := range line {
			top, right, bottom, left := getSides(i, j, &heightmap)

			if checkIfMinimum([]int{top, right, bottom, left}, middle) {
				topBasins = append(topBasins, exploreBasil(i, j, &heightmap, &visitedMap))
			}
		}
	}

	sort.Ints(topBasins)

	basinsProd := topBasins[len(topBasins)-1] *
		topBasins[len(topBasins)-2] *
		topBasins[len(topBasins)-3]

	fmt.Println("basins product", basinsProd)
}
