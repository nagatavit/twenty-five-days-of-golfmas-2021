package main

import (
	"fmt"
)

func secondPart(manual [][]bool, foldingInstructions []instruction) {
	fmt.Println("====================")
	fmt.Println("Starting second part")
	fmt.Println("====================")

	newManual := manual

	for _, inst := range foldingInstructions {
		switch inst.axis {
		case "x":
			newManual = foldYAxis(newManual, inst.position)
		case "y":
			newManual = foldXAxis(newManual, inst.position)
		}
	}

	printManual(newManual)
}
