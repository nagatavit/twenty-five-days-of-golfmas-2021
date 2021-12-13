package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type cave struct {
	name           string
	connections    []*cave
	isBigSize      bool
	hasBeenVisited bool
}

func readInput() (cavesListing map[string]*cave) {
	// f, err := os.Open("example-input-3")
	f, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	// Connections between pairwise caves.
	var cavesAndConnections [][2]string

	// Listings of all the caves by the name
	cavesListing = make(map[string]*cave)

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		line := scanner.Text()

		caves := strings.Split(line, "-")

		cavesAndConnections = append(cavesAndConnections, [2]string{caves[0], caves[1]})

		for i := 0; i < 2; i++ {
			if _, ok := cavesListing[caves[i]]; !ok {
				newCave := new(cave)
				newCave.name = caves[i]

				if strings.ToUpper(newCave.name) == newCave.name {
					newCave.isBigSize = true
				}

				cavesListing[caves[i]] = newCave
			}
		}
	}

	for _, caves := range cavesAndConnections {
		for i := range caves {
			// i can assume only two values: 0 or 1 (position of the
			// origin cave connection and position of the destine cave
			// connection)
			//
			// The i ^ 0x1 index gimmick is just to address the
			// opposite side of the cave connection
			cavesListing[caves[i]].connections = append(
				cavesListing[caves[i]].connections,
				cavesListing[caves[i^0x1]],
			)
		}
	}

	return cavesListing
}

func printCavesFromListings(cavesListing map[string]*cave) {
	for i, cave := range cavesListing {
		fmt.Printf("%p: \n", cavesListing[i])
		fmt.Printf("\tname:\t %s\n", cave.name)
		fmt.Printf("\tconnections:\t %v\n", cave.connections)
		fmt.Printf("\tisBigSize:\t %v\n", cave.isBigSize)
	}
}

func main() {
	// Add location (line and file) of log calls
	log.SetFlags(log.Lshortfile)

	cavesListing := readInput()
	firstPart(cavesListing)
}
