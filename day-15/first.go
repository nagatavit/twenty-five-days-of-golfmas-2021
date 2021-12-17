package main

import (
	"fmt"
	"math"
)

type coordinate struct {
	x int
	y int
}

type dijkstraNode struct {
	currentTotalCost int
	position         coordinate
	previousPosition coordinate
}

// neighbors are a slice containing:
//
// top, right, bottom, left values of cavernMap (reference j, i)
//
// in that order
func getNeighborsCost(j, i int, dijkstraVisited *map[coordinate]dijkstraNode, cavernMap *[][]int) (neighbors []dijkstraNode) {
	var top, right, bottom, left dijkstraNode

	top.previousPosition = coordinate{j, i}
	right.previousPosition = coordinate{j, i}
	bottom.previousPosition = coordinate{j, i}
	left.previousPosition = coordinate{j, i}

	if i-1 < 0 {
		top.currentTotalCost = -1
	} else {
		top.currentTotalCost = (*cavernMap)[i-1][j] + (*dijkstraVisited)[coordinate{j, i}].currentTotalCost
		top.position = coordinate{j, i - 1}
	}

	if j+1 > len((*cavernMap)[0])-1 {
		right.currentTotalCost = -1
	} else {
		right.currentTotalCost = (*cavernMap)[i][j+1] + (*dijkstraVisited)[coordinate{j, i}].currentTotalCost
		right.position = coordinate{j + 1, i}
	}

	if i+1 > len(*cavernMap)-1 {
		bottom.currentTotalCost = -1
	} else {
		bottom.currentTotalCost = (*cavernMap)[i+1][j] + (*dijkstraVisited)[coordinate{j, i}].currentTotalCost
		bottom.position = coordinate{j, i + 1}
	}

	if j-1 < 0 {
		left.currentTotalCost = -1
	} else {
		left.currentTotalCost = (*cavernMap)[i][j-1] + (*dijkstraVisited)[coordinate{j, i}].currentTotalCost
		left.position = coordinate{j - 1, i}
	}

	return []dijkstraNode{top, right, bottom, left}
}

func updateNeighborsPaths(j, i int, neighbor dijkstraNode, dijkstraVisited *map[coordinate]dijkstraNode, dijkstraKnownNeighbors *map[coordinate]dijkstraNode) {
	if neighbor.currentTotalCost == -1 {
		return
	}

	// See if we already visited that neighbor
	if _, ok := (*dijkstraVisited)[neighbor.position]; ok {
		return
	}

	previousBest, ok := (*dijkstraKnownNeighbors)[neighbor.position]
	if !ok {
		(*dijkstraKnownNeighbors)[neighbor.position] = neighbor
	} else {

		if previousBest.currentTotalCost > neighbor.currentTotalCost {
			(*dijkstraKnownNeighbors)[neighbor.position] = neighbor
		}
	}
}

func getNewestMinimum(dijkstraKnownNeighbors *map[coordinate]dijkstraNode, dijkstraVisited *map[coordinate]dijkstraNode) (newMinimum coordinate) {
	minimumVal := math.MaxInt

	for coord, node := range *dijkstraKnownNeighbors {
		if node.currentTotalCost < minimumVal {
			minimumVal = node.currentTotalCost
			newMinimum = coord
		}
	}

	(*dijkstraVisited)[newMinimum] = (*dijkstraKnownNeighbors)[newMinimum]

	delete(*dijkstraKnownNeighbors, newMinimum)

	return newMinimum
}

func printMap(dijMap map[coordinate]dijkstraNode) {
	for key, val := range dijMap {
		fmt.Println("    ", key, val)
	}
}

func firstPart(cavernMap [][]int) {
	dijkstraVisited := make(map[coordinate]dijkstraNode)
	dijkstraKnownNeighbors := make(map[coordinate]dijkstraNode)

	// First position always has cost 0 (starting position)
	dijkstraVisited[coordinate{0, 0}] = dijkstraNode{
		currentTotalCost: 0,
		position:         coordinate{0, 0},
		previousPosition: coordinate{0, 0},
	}

	// Dijkstra needs to scan each node to find the minimum spanning tree
	numberOfSteps := len(cavernMap) * len(cavernMap[0])

	currXPosition := 0
	currYPosition := 0

	for len(dijkstraVisited) < numberOfSteps {
		neighbors := getNeighborsCost(currXPosition, currYPosition, &dijkstraVisited, &cavernMap)

		for _, neighbor := range neighbors {
			updateNeighborsPaths(currXPosition, currYPosition, neighbor, &dijkstraVisited, &dijkstraKnownNeighbors)
		}

		newCoord := getNewestMinimum(&dijkstraKnownNeighbors, &dijkstraVisited)
		currXPosition = newCoord.x
		currYPosition = newCoord.y
	}

	fmt.Println(dijkstraVisited[coordinate{
		x: len(cavernMap[0]) - 1,
		y: len(cavernMap) - 1,
	}])
}
