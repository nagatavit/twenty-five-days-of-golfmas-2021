package main

import "fmt"

func firstPart(entries []signalNote) {
	fmt.Println("===================")
	fmt.Println("Starting first part")
	fmt.Println("===================")

	// easiest digits have the following number of edges:
	// 1: 2 edges
	// 4: 4 edges
	// 7: 3 edges
	// 8: 7 edges
	var easyDigitEdgesToLen = map[int]bool{
		2: true,
		4: true,
		3: true,
		7: true,
	}

	easyDigitCount := 0

	for _, entry := range entries {
		for _, digit := range entry.digits {
			if _, ok := easyDigitEdgesToLen[len(digit)]; ok {
				easyDigitCount++
			}
		}
	}

	fmt.Println("easy digits: ", easyDigitCount)
}
