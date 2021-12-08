package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

type fish struct {
	counter int
}

func (fish *fish) AnotherDayPassesBy() (newborn bool) {
	if fish.counter == 0 {
		fish.counter = 6
		return true
	}

	fish.counter--
	return false
}

func readInput() (initialFish []fish) {
	f, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	scanner.Scan()
	initialFishStr := strings.Split(scanner.Text(), ",")

	initialFish = make([]fish, len(initialFishStr))

	for i, fishStr := range initialFishStr {
		initialFish[i].counter, err = strconv.Atoi(fishStr)
		if err != nil {
			log.Fatal(err)
		}
	}

	return initialFish
}

func main() {
	initialFish := readInput()
	firstPart(initialFish)

	// Reset the initial condition as we changed the values in the
	// first part by using the AnotherDayPassesBy method
	initialFish = readInput()
	secondPart(initialFish)
}
