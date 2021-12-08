package main

import "fmt"

func firstPart(initialFish []fish) {
	fmt.Println("===================")
	fmt.Println("Starting first part")
	fmt.Println("===================")

	days := 80

	for i := 0; i < days; i++ {
		for j := range initialFish {
			if initialFish[j].AnotherDayPassesBy() {
				initialFish = append(initialFish,
					fish{counter: 8},
				)
			}
		}
	}

	fmt.Println("shoal", len(initialFish))
}
