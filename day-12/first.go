package main

import (
	"fmt"
)

// From what I saw (only checked for direct connections on big caves),
// seems like in all the inputs there isn't a way to be stuck in a
// loop, but it doesn't cost to check.
var maxDepthSearch = 10000

func findPossiblePaths(depth int, currentCave *cave) (pathsFound [][]string) {
	if depth > maxDepthSearch {
		return nil
	}

	if currentCave.name == "end" {
		pathsFound = append(pathsFound, []string{"end"})
		return pathsFound
	} else if !currentCave.isBigSize && currentCave.hasBeenVisited {
		return nil
	}

	var newPaths [][]string

	currentCave.hasBeenVisited = true

	depth++
	for i := range currentCave.connections {
		newPaths = findPossiblePaths(depth, currentCave.connections[i])
		if newPaths == nil {
			continue
		}

		for _, newPath := range newPaths {
			newPath = append([]string{currentCave.name}, newPath...)
			pathsFound = append(pathsFound, newPath)
		}
	}

	currentCave.hasBeenVisited = false

	return pathsFound
}

func firstPart(cavesListing map[string]*cave) {
	fmt.Println("===================")
	fmt.Println("Starting first part")
	fmt.Println("===================")

	startCave := cavesListing["start"]

	paths := findPossiblePaths(0, startCave)

	fmt.Println("Number of paths found:", len(paths))
}
