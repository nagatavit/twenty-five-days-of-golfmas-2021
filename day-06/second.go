package main

import "fmt"

type fishGroupByDay struct {
	counter  int
	newborns int
}

func (fishGroup *fishGroupByDay) increaseNewborns(newNewborns int) {
	fishGroup.newborns = newNewborns
}

func (fishGroup *fishGroupByDay) generateNewborns() (newNewborns int) {
	newNewborns = fishGroup.counter

	// reset the newborns and add them to the adult pop for the next
	// iteration. We can do this because the next iteration with this
	// group would be in 7 days, which is already past the time for a
	// newborn to become an adult.
	fishGroup.counter += fishGroup.newborns
	fishGroup.newborns = 0

	return newNewborns
}

func printFishCountByDay(fishCountByDay []fishGroupByDay) {
	for i, fishByDay := range fishCountByDay {
		fmt.Println(i, fishByDay)
	}
}

func secondPart(initialFish []fish) {
	fmt.Println("====================")
	fmt.Println("Starting second part")
	fmt.Println("====================")

	fishCountByDay := make([]fishGroupByDay, 7)

	// Move the initial fish to this new slice
	for _, fish := range initialFish {
		fishCountByDay[fish.counter].counter++
	}

	// Two days to account for the newborns (I guess?)
	days := 256 + 2

	for i := 0; i < days; i++ {
		newNewborns := fishCountByDay[i%7].generateNewborns()
		fishCountByDay[(i+2)%7].increaseNewborns(newNewborns)
	}

	totalFish := 0

	for i := 0; i < 7; i++ {
		totalFish += fishCountByDay[i].counter
	}

	fmt.Println("shoal", totalFish)
}
