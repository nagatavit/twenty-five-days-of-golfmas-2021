package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func secondPart() {
	fmt.Println("====================")
	fmt.Println("Starting second part")
	fmt.Println("====================")

	f, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	var lastWindowSum int
	var slidingWindow []int

	// read the first three separately
	for i := 0; i < 3; i++ {
		scanner.Scan()
		curr, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}

		slidingWindow = append(slidingWindow, curr)
		lastWindowSum += curr
	}

	inc := 0

	for scanner.Scan() {
		curr, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}

		newWindowSum := lastWindowSum - slidingWindow[0] + curr

		if newWindowSum > lastWindowSum {
			inc++
		}

		slidingWindow = append(slidingWindow, curr)
		slidingWindow = slidingWindow[1:]
	}

	fmt.Println(inc)
}
