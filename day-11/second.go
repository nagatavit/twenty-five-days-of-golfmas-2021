package main

import (
	"fmt"
)

func secondPart(octoMap [][]dumboOctopus) {
	fmt.Println("====================")
	fmt.Println("Starting second part")
	fmt.Println("====================")

	allFlashesSync := false

	i := 0
	for !allFlashesSync {
		if newStep(&octoMap) == len(octoMap)*len(octoMap[0]) {
			allFlashesSync = true
		}
		i++
	}

	fmt.Println("First synchronized flash", i)
}
