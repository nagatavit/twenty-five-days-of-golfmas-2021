package main

import (
	"fmt"
)

func findPossiblePathsVisitingASmallOne(depth int, currentCave *cave, smallHasBeenVisited bool, smallVisitedTwice string) (pathsFound [][]string) {
	if depth > maxDepthSearch {
		return nil
	}

	if currentCave.name == "end" {
		pathsFound = append(pathsFound, []string{"end"})
		return pathsFound
	} else if !currentCave.isBigSize && currentCave.hasBeenVisited {
		if smallHasBeenVisited || currentCave.name == "start" {
			return nil
		} else {
			smallVisitedTwice = currentCave.name
			smallHasBeenVisited = true
		}
	}

	var newPaths [][]string

	currentCave.hasBeenVisited = true

	depth++
	for i := range currentCave.connections {
		newPaths = findPossiblePathsVisitingASmallOne(depth, currentCave.connections[i], smallHasBeenVisited, smallVisitedTwice)
		if newPaths == nil {
			continue
		}

		for _, newPath := range newPaths {
			newPath = append([]string{currentCave.name}, newPath...)
			pathsFound = append(pathsFound, newPath)
		}
	}

	if !smallHasBeenVisited || currentCave.name != smallVisitedTwice {
		currentCave.hasBeenVisited = false
	}

	return pathsFound
}

func secondPart(cavesListing map[string]*cave) {
	fmt.Println("====================")
	fmt.Println("Starting second part")
	fmt.Println("====================")

	startCave := cavesListing["start"]

	paths := findPossiblePathsVisitingASmallOne(0, startCave, false, "")

	fmt.Println("Number of paths found:", len(paths))
}
